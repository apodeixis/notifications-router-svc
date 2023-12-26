package handlers

import (
	"net/http"
	"net/url"

	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/registration/ctx"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/registration/responses"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/apodeixis/notifications-router-svc/internal/service/api/registration/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RegisterService(w http.ResponseWriter, r *http.Request) {
	log := ctx.Log(r)
	request, err := requests.NewRegisterService(r)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	endpoint, err := url.Parse(request.Data.Attributes.Endpoint)
	if err != nil {
		log.WithError(err).Error("invalid request")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{"/data/attributes/endpoint": err})...)
		return
	}
	svc := &data.NotificatorService{
		Endpoint: endpoint,
		Channels: request.Data.Attributes.Channels,
	}
	if err := ctx.NotificatorsStorage(r).Add(svc); err != nil {
		log.WithError(err).Error("failed to add notificator service")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	response := responses.ComposeNotificatorService(svc)
	ape.Render(w, response)
}
