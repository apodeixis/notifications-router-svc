package processor

import (
	"errors"

	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/internal/data/postgres"
	"github.com/apodeixis/notifications-router-svc/internal/types"
)

func newQuerier(db *pgdb.DB) *querier {
	return &querier{
		deliveriesQ:    postgres.NewDeliveriesQ(db),
		notificationsQ: postgres.NewNotificationsQ(db),
	}
}

type querier struct {
	deliveriesQ    data.DeliveriesQ
	notificationsQ data.NotificationsQ
}

func (q *querier) getPendingDeliveries() ([]data.Delivery, error) {
	return q.deliveriesQ.New().
		JoinNotification().
		FilterByStatus(types.DeliveryStatusNotSent).
		Select()
}

func (q *querier) getNotification(id int64) (data.Notification, error) {
	result, err := q.notificationsQ.New().
		FilterByID(id).
		Get()
	if result == nil {
		return data.Notification{}, errors.New("notification does not exist")
	}
	return *result, err
}

func (q *querier) setDeliveryStatus(id int64, status types.DeliveryStatus) error {
	_, err := q.deliveriesQ.New().
		FilterByID(id).
		SetStatus(status).
		Update()
	return err
}
