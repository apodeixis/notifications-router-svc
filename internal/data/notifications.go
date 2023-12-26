package data

import (
	"time"
)

type NotificationsQ interface {
	New() NotificationsQ

	Get() (*Notification, error)
	Select() ([]Notification, error)

	Transaction(fn func(q NotificationsQ) error) error

	Insert(data Notification) (*Notification, error)
	InsertDeliveries(data []Delivery) ([]Delivery, error)

	FilterByID(id ...int64) NotificationsQ
}

type Notification struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Topic     string    `db:"topic"`
	Channel   *string   `db:"channel"`
	Message   Message   `db:"message"`
}
