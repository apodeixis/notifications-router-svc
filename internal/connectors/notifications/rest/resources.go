package rest

import (
	"github.com/apodeixis/notifications-router-svc/internal/data"
)

type (
	createNotificationRequest struct {
		Data createNotificationData `json:"data"`
	}
	createNotificationData struct {
		Type          string                          `json:"type"`
		Attributes    createNotificationAttributes    `json:"attributes"`
		Relationships createNotificationRelationships `json:"relationships"`
	}
	createNotificationAttributes struct {
		Message data.Message `json:"message"`
		Channel string       `json:"channel"`
	}
	createNotificationRelationships struct {
		Destination createNotificationDestination `json:"destination"`
	}
	createNotificationDestination struct {
		Data DestinationKey `json:"data"`
	}
	DestinationKey struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
)
