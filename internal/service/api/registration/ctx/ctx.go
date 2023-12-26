package ctx

import (
	"context"
	"net/http"

	"github.com/apodeixis/notifications-router-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota + 1

	notificatorsStorageCtxKey
)

func SetLog(v *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, v)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func SetNotificatorsStorage(v data.NotificatorsStorage) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, notificatorsStorageCtxKey, v)
	}
}

func NotificatorsStorage(r *http.Request) data.NotificatorsStorage {
	return r.Context().Value(notificatorsStorageCtxKey).(data.NotificatorsStorage)
}
