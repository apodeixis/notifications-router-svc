package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/apodeixis/notifications-router-svc/internal/service/api/primary/ctx"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/primary/responses"
	"github.com/apodeixis/notifications-router-svc/internal/types"

	"github.com/apodeixis/notifications-router-svc/internal/data"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/apodeixis/notifications-router-svc/internal/service/api/primary/requests"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewCreateNotification(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	var (
		notification *data.Notification
		deliveries   []data.Delivery
	)
	err = ctx.NotificationsQ(r).Transaction(func(q data.NotificationsQ) error {
		notification, err = q.Insert(data.Notification{
			Topic:   request.Data.Attributes.Topic,
			Channel: request.Data.Attributes.Channel,
			Message: data.Message(request.Data.Attributes.Message),
		})
		if err != nil {
			return errors.Wrap(err, "failed to insert notification")
		}
		deliveries = make([]data.Delivery, len(request.Data.Relationships.Destinations.Data))
		for i, destination := range request.Data.Relationships.Destinations.Data {
			deliveries[i] = data.Delivery{
				NotificationID:  notification.ID,
				Destination:     destination.Id,
				DestinationType: types.DestinationType(destination.Type),
				Status:          types.DeliveryStatusNotSent,
			}
		}
		deliveries, err = q.InsertDeliveries(deliveries)
		if err != nil {
			return errors.Wrap(err, "failed to insert deliveries")
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := responses.ComposeNotification(notification, deliveries)
	ape.Render(w, response)
}
