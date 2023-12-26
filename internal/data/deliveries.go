package data

import (
	"time"

	"github.com/apodeixis/notifications-router-svc/internal/types"
)

type DeliveriesQ interface {
	New() DeliveriesQ

	Get() (*Delivery, error)
	Select() ([]Delivery, error)
	Update() ([]Delivery, error)

	Transaction(fn func(q DeliveriesQ) error) error

	FilterByStatus(statuses ...types.DeliveryStatus) DeliveriesQ
	FilterByID(ids ...int64) DeliveriesQ

	JoinNotification() DeliveriesQ

	SetStatus(status types.DeliveryStatus) DeliveriesQ
}

type Delivery struct {
	ID              int64                 `db:"id"`
	NotificationID  int64                 `db:"notification_id"`
	Destination     string                `db:"destination"`
	DestinationType types.DestinationType `db:"destination_type"`
	Status          types.DeliveryStatus  `db:"status"`
	SentAt          time.Time             `db:"sent_at"`
}
