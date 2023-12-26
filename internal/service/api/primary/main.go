package primary

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/notifications-router-svc/internal/config"
	"github.com/apodeixis/notifications-router-svc/internal/data/postgres"
	base "github.com/apodeixis/notifications-router-svc/internal/service/api"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/primary/ctx"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/primary/handlers"

	"gitlab.com/distributed_lab/ape"
)

type api struct {
	router   chi.Router
	listener net.Listener
	log      *logan.Entry
}

func (a *api) Start() error {
	a.log.Info("Primary api started on ", a.listener.Addr().String())
	return http.Serve(a.listener, a.router)
}

func NewAPI(cfg config.Config) base.API {
	return &api{
		router:   newRouter(cfg),
		listener: cfg.Listener(),
		log:      cfg.Log(),
	}
}

func newRouter(cfg config.Config) chi.Router {
	r := chi.NewRouter()
	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			ctx.SetLog(cfg.Log()),

			ctx.SetNotificationsQ(postgres.NewNotificationsQ(cfg.DB())),
		),
	)
	r.Post("/notifications", handlers.CreateNotification)

	return r
}
