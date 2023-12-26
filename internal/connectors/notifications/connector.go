package notifications

import (
	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/internal/providers/identifier"
)

type Connector interface {
	SendNotification(destination identifier.Identifier, message data.Message, channel string) error
}
