package registration

import (
	"net"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/apodeixis/notifications-router-svc/internal/config"
	"github.com/apodeixis/notifications-router-svc/internal/data"
	base "github.com/apodeixis/notifications-router-svc/internal/service/api"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/registration/ctx"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/registration/handlers"

	"github.com/go-chi/chi/v5"
	"gitlab.com/distributed_lab/ape"
)

type api struct {
	router   chi.Router
	listener net.Listener
	log      *logan.Entry
}

func (a *api) Start() error {
	a.log.Info("Registration api started on ", a.listener.Addr().String())
	return http.Serve(a.listener, a.router)
}

func NewAPI(cfg config.Config, storage data.NotificatorsStorage) base.API {
	return &api{
		router:   newRouter(cfg, storage),
		listener: cfg.RegistrationAPIListener(),
		log:      cfg.Log(),
	}
}

func newRouter(cfg config.Config, storage data.NotificatorsStorage) chi.Router {
	r := chi.NewRouter()
	r.Use(
		ape.RecoverMiddleware(cfg.Log()),
		ape.LoganMiddleware(cfg.Log()),
		ape.CtxMiddleware(
			ctx.SetLog(cfg.Log()),

			ctx.SetNotificatorsStorage(storage),
		),
	)
	r.Route("/internal", func(r chi.Router) {
		r.Post("/services", handlers.RegisterService)
	})
	return r
}
