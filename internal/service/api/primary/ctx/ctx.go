package ctx

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/notifications-router-svc/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota + 1

	notificationsQCtxKey
	deliveriesQCtxKey
)

func SetLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func SetNotificationsQ(entry data.NotificationsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, notificationsQCtxKey, entry)
	}
}

func NotificationsQ(r *http.Request) data.NotificationsQ {
	return r.Context().Value(notificationsQCtxKey).(data.NotificationsQ).New()
}

func SetDeliveriesQ(entry data.DeliveriesQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, deliveriesQCtxKey, entry)
	}
}

func DeliveriesQ(r *http.Request) data.DeliveriesQ {
	return r.Context().Value(deliveriesQCtxKey).(data.DeliveriesQ).New()
}
