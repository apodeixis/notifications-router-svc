package processor

import (
	"github.com/pkg/errors"

	"github.com/apodeixis/notifications-router-svc/internal/connectors/notifications"
	"github.com/apodeixis/notifications-router-svc/internal/connectors/notifications/rest"
	"github.com/apodeixis/notifications-router-svc/internal/data"
)

type notificatorsConnectorProvider struct {
	storage data.NotificatorsStorage
}

func (p *notificatorsConnectorProvider) getByChannel(chanel string) (notifications.Connector, error) {
	service, err := p.storage.GetByChannel(chanel)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get notificators service")
	}

	return rest.NewRestNotificationsConnector(service.Endpoint), nil
}
