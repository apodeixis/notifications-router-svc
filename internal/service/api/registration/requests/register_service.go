package requests

import (
	"encoding/json"
	"net/http"

	"github.com/apodeixis/notifications-router-svc/resources"

	"github.com/go-ozzo/ozzo-validation/v4/is"

	"gitlab.com/distributed_lab/logan/v3/errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RegisterServiceRequest struct {
	Data resources.NotificatorService
}

func NewRegisterService(r *http.Request) (*RegisterServiceRequest, error) {
	request := new(RegisterServiceRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return nil, errors.Wrap(err, "failed to decode request body")
	}
	return request, request.validate()
}

func (r *RegisterServiceRequest) validate() error {
	return validation.Errors{
		"data/attributes/endpoint": validation.Validate(&r.Data.Attributes.Endpoint, validation.Required, is.RequestURI),
		"data/attributes/channels": validation.Validate(&r.Data.Attributes.Channels, validation.Required),
	}.Filter()
}
