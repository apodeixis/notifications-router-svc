package responses

import (
	"strconv"

	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/resources"
)

func ComposeNotification(
	notification *data.Notification,
	deliveries []data.Delivery,
) *resources.CreateNotification200Response {
	deliveriesResources := make([]resources.DeliveryKey, len(deliveries))
	for i, delivery := range deliveries {
		deliveriesResources[i] = resources.DeliveryKey{
			Id:   strconv.FormatInt(delivery.ID, 10),
			Type: "delivery",
		}
	}
	return &resources.CreateNotification200Response{
		Data: resources.Notification{
			Id:   strconv.FormatInt(notification.ID, 10),
			Type: "notification",
			Attributes: resources.NotificationAllOfAttributes{
				CreatedAt: notification.CreatedAt.Unix(),
				Topic:     notification.Topic,
				Channel:   notification.Channel,
				Message:   resources.Message(notification.Message),
			},
			Relationships: resources.NotificationAllOfRelationships{
				Deliveries: resources.NotificationAllOfRelationshipsDeliveries{
					Data: deliveriesResources,
				},
			},
		},
	}
}
