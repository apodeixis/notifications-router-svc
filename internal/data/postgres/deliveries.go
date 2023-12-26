package postgres

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/internal/types"
)

const deliveriesTableName = "deliveries"

func NewDeliveriesQ(db *pgdb.DB) data.DeliveriesQ {
	return &deliveriesQ{
		db:        db.Clone(),
		sql:       sq.Select("deliveries.*").From(deliveriesTableName),
		sqlUpdate: sq.Update(deliveriesTableName).Suffix("returning *"),
	}
}

type deliveriesQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *deliveriesQ) New() data.DeliveriesQ {
	return NewDeliveriesQ(q.db)
}

func (q *deliveriesQ) Get() (*data.Delivery, error) {
	var result data.Delivery
	err := q.db.Get(&result, q.sql)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}

func (q *deliveriesQ) Select() ([]data.Delivery, error) {
	var result []data.Delivery
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *deliveriesQ) Update() ([]data.Delivery, error) {
	var result []data.Delivery
	err := q.db.Select(&result, q.sqlUpdate)

	return result, err
}

func (q *deliveriesQ) Transaction(fn func(q data.DeliveriesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *deliveriesQ) FilterByStatus(statuses ...types.DeliveryStatus) data.DeliveriesQ {
	stmt := sq.Eq{"deliveries.status": statuses}
	q.sql = q.sql.Where(stmt)
	q.sqlUpdate = q.sqlUpdate.Where(stmt)
	return q
}

func (q *deliveriesQ) FilterByID(ids ...int64) data.DeliveriesQ {
	stmt := sq.Eq{"deliveries.id": ids}
	q.sql = q.sql.Where(stmt)
	q.sqlUpdate = q.sqlUpdate.Where(stmt)
	return q
}

func (q *deliveriesQ) JoinNotification() data.DeliveriesQ {
	stmt := fmt.Sprintf("%s as notification on notification.id = deliveries.notification_id",
		notificationsTableName)
	q.sql = q.sql.Join(stmt)
	return q
}

func (q *deliveriesQ) SetStatus(status types.DeliveryStatus) data.DeliveriesQ {
	q.sqlUpdate = q.sqlUpdate.Set("status", status)
	return q
}
