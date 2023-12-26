package processor

import (
	"context"
	"time"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/notifications-router-svc/internal/service/runners"
	"github.com/apodeixis/notifications-router-svc/internal/types"

	"github.com/apodeixis/notifications-router-svc/internal/providers/templates"

	"github.com/apodeixis/notifications-router-svc/internal/providers/identifier"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/distributed_lab/running"

	"github.com/apodeixis/notifications-router-svc/internal/data"

	"github.com/apodeixis/notifications-router-svc/internal/config"
)

const (
	serviceName = "Notifications-processor"
)

func New(cfg config.Config, storage data.NotificatorsStorage) runners.Runner {
	return &processor{
		log:            cfg.Log().WithField("runner", serviceName),
		notificatorCfg: cfg.NotificatorConfig(),
		querier:        newQuerier(cfg.DB()),
		notificationsConnectorProvider: &notificatorsConnectorProvider{
			storage: storage,
		},
		templatesHelper: &templatesHelper{
			notificatorCfg:    cfg.NotificatorConfig(),
			templatesProvider: templates.NewS3TemplatesProvider(cfg.AwsConfig(), cfg.Bucket()),
		},
	}
}

type processor struct {
	log                            *logan.Entry
	querier                        *querier
	notificatorCfg                 *config.NotificatorConfig
	notificationsConnectorProvider *notificatorsConnectorProvider
	identifierProvider             identifier.IdentifierProvider
	templatesHelper                *templatesHelper
}

func (p *processor) Run(ctx context.Context) error {
	p.log.Info(serviceName, " started")
	running.WithBackOff(ctx, p.log,
		serviceName,
		p.processNotifications,
		3*time.Second,
		3*time.Second,
		3*time.Second,
	)
	return nil
}

func (p *processor) processNotifications(_ context.Context) error {
	deliveries, err := p.querier.getPendingDeliveries()
	if err != nil {
		return errors.Wrap(err, "failed to get pending deliveries")
	}
	for _, delivery := range deliveries {
		p.log.WithFields(getLoggerFields(delivery)).
			Info("processing notification")
		if err := p.processDelivery(delivery); err != nil {
			p.log.WithFields(getLoggerFields(delivery)).WithError(err).
				Error("failed to send to notification, marking it as failed")
			if err := p.querier.setDeliveryStatus(delivery.ID, types.DeliveryStatusFailed); err != nil {
				return errors.Wrap(err, "failed to set delivery status")
			}
		}
	}
	return nil
}

func (p *processor) processDelivery(delivery data.Delivery) error {
	notification, err := p.querier.getNotification(delivery.NotificationID)
	if err != nil {
		return errors.Wrap(err, "failed to get notification")
	}

	channelsList, err := p.getChannels(notification)
	if err != nil {
		return errors.Wrap(err, "failed to get channels")
	}

	for _, channel := range channelsList {
		err = p.sendNotification(channel, delivery, notification)
		if err != nil {
			p.log.WithFields(getLoggerFields(delivery)).
				WithError(err).
				Warnf("failed to send notification with channel - %s, try next channel", channel)
			continue
		}

		if err := p.querier.setDeliveryStatus(delivery.ID, types.DeliveryStatusSent); err != nil {
			return errors.Wrap(err, "failed to set delivery status")
		}
		return nil
	}

	return errors.New("failed to send notification via all available channels")
}

func (p *processor) sendNotification(channel string, delivery data.Delivery, notification data.Notification) error {
	message, err := p.templatesHelper.buildMessage(channel, notification)
	if err != nil {
		return errors.Wrap(err, "failed to create message from template")
	}

	id := p.getIdentifier(delivery)

	connector, err := p.notificationsConnectorProvider.getByChannel(channel)
	if err != nil {
		return errors.Wrap(err, "failed to get notifications connector")
	}

	p.log.WithFields(map[string]interface{}{
		"channel": channel,
		"message": message,
	}).Debug("Sending message")

	err = connector.SendNotification(id, message, channel)
	if err != nil {
		return errors.Wrap(err, "failed to send notification")
	}

	return nil
}

func (p *processor) getChannels(notification data.Notification) ([]string, error) {
	if notification.Channel != nil {
		return []string{*notification.Channel}, nil
	}
	return p.notificatorCfg.DefaultChannelsPriority, nil
}

func (p *processor) getIdentifier(delivery data.Delivery) identifier.Identifier {
	return identifier.Identifier{
		ID:   delivery.Destination,
		Type: string(delivery.DestinationType),
	}
}

func getLoggerFields(delivery data.Delivery) map[string]interface{} {
	return map[string]interface{}{
		"delivery_id":     delivery.ID,
		"notification_id": delivery.NotificationID,
		"destination":     delivery.Destination,
	}
}
