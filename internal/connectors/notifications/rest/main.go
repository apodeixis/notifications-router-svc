package rest

import (
	"context"

	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/internal/providers/identifier"
)

func (c *restNotificationsConnector) SendNotification(identifier identifier.Identifier, message data.Message, channel string) error {
	request := composeCreateNotificationRequest(identifier, message, channel)
	return c.connector.PostJSON(c.endpoint, request, context.Background(), nil)
}

func composeCreateNotificationRequest(identifier identifier.Identifier, message data.Message, channel string) *createNotificationRequest {
	const createNotificationType = "create_notification"
	return &createNotificationRequest{
		Data: createNotificationData{
			Type: createNotificationType,
			Attributes: createNotificationAttributes{
				Message: message,
				Channel: channel,
			},
			Relationships: createNotificationRelationships{
				Destination: createNotificationDestination{
					Data: DestinationKey{
						ID:   identifier.ID,
						Type: identifier.Type,
					},
				},
			},
		},
	}
}
