package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/apodeixis/notifications-router-svc/internal/data"
)

const notificationsTableName = "notifications"

func NewNotificationsQ(db *pgdb.DB) data.NotificationsQ {
	return &notificationsQ{
		db:  db.Clone(),
		sql: sq.Select("n.*").From(fmt.Sprintf("%s as n", notificationsTableName)),
	}
}

type notificationsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func (q *notificationsQ) New() data.NotificationsQ {
	return NewNotificationsQ(q.db)
}

func (q *notificationsQ) Get() (*data.Notification, error) {
	var result data.Notification
	err := q.db.Get(&result, q.sql)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &result, err
}

func (q *notificationsQ) Select() ([]data.Notification, error) {
	var result []data.Notification
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *notificationsQ) Transaction(fn func(q data.NotificationsQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *notificationsQ) Insert(value data.Notification) (*data.Notification, error) {
	clauses := map[string]interface{}{
		"topic":   value.Topic,
		"channel": value.Channel,
		"message": value.Message,
	}
	result := new(data.Notification)
	stmt := sq.Insert(notificationsTableName).SetMap(clauses).Suffix("RETURNING *")
	err := q.db.Get(result, stmt)
	return result, err
}

func (q *notificationsQ) InsertDeliveries(deliveries []data.Delivery) ([]data.Delivery, error) {
	if len(deliveries) == 0 {
		return nil, errors.New("empty array is not allowed")
	}
	columns := []string{
		"notification_id",
		"destination",
		"destination_type",
		"status",
		"sent_at",
	}
	stmt := sq.Insert(deliveriesTableName).Columns(columns...)
	for _, item := range deliveries {
		stmt = stmt.Values([]interface{}{
			item.NotificationID,
			item.Destination,
			item.DestinationType,
			item.Status,
			item.SentAt,
		}...)
	}
	stmt = stmt.Suffix("RETURNING *")
	var result []data.Delivery
	err := q.db.Select(&result, stmt)
	return result, err
}

func (q *notificationsQ) FilterByID(ids ...int64) data.NotificationsQ {
	q.sql = q.sql.Where(sq.Eq{"n.id": ids})
	return q
}
