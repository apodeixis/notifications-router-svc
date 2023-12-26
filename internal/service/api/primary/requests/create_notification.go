package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/apodeixis/notifications-router-svc/internal/types"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/apodeixis/notifications-router-svc/resources"
)

func NewCreateNotification(r *http.Request) (*resources.CreateNotificationRequest, error) {
	request := new(resources.CreateNotificationRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal request body")
	}
	return request, validateCreateNotificationRequest(request)
}

func validateCreateNotificationRequest(r *resources.CreateNotificationRequest) error {
	return mergeErrors(validation.Errors{
		"/data/attributes/topic":              validation.Validate(&r.Data.Attributes.Topic, validation.Required),
		"/data/attributes/message/type":       validation.Validate(&r.Data.Attributes.Message.Type, validation.Required),
		"/data/attributes/message/attributes": validation.Validate(&r.Data.Attributes.Message.Attributes, validation.Required),
	},
		validateDestinationsList(r.Data.Relationships.Destinations.Data),
	).Filter()
}

func validateDestinationsList(destinations []resources.DestinationKey) validation.Errors {
	validationErrors := validation.Errors{
		"/data/relationships/destinations/data": validation.Validate(&destinations, validation.Required),
	}

	for i, destination := range destinations {
		validationErrors[fmt.Sprintf("/data/relationships/destinations/data/%d", i)] =
			validateDestination(destination)
	}

	return validationErrors
}

func validateDestination(destination resources.DestinationKey) error {
	switch destination.Type {
	case string(types.DestinationTypeEmail):
		if !govalidator.IsEmail(destination.Id) {
			return errors.New("must be valid email")
		}
	}
	return nil
}

func mergeErrors(validationErrors ...validation.Errors) validation.Errors {
	result := make(validation.Errors)
	for _, errs := range validationErrors {
		for key, err := range errs {
			result[key] = err
		}
	}
	return result
}
