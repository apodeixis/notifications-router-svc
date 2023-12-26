package responses

import (
	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/resources"
)

func ComposeNotificatorService(svc *data.NotificatorService) *resources.RegisterServiceRequest {
	return &resources.RegisterServiceRequest{
		Data: resources.NotificatorService{
			Type: "notificator_service",
			Attributes: resources.NotificatorServiceAllOfAttributes{
				Channels: svc.Channels,
				Endpoint: svc.Endpoint.String(),
			},
		},
	}
}
