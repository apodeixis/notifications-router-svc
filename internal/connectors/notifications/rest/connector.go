package rest

import (
	"net/http"
	"net/url"

	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/tokend/connectors/signed"

	"github.com/apodeixis/notifications-router-svc/internal/connectors/notifications"
)

type restNotificationsConnector struct {
	endpoint  *url.URL
	connector *jsonapi.Connector
}

func NewRestNotificationsConnector(endpoint *url.URL) notifications.Connector {
	host, _ := url.Parse(endpoint.Host)
	client := signed.NewClient(http.DefaultClient, host)
	return &restNotificationsConnector{
		endpoint:  endpoint,
		connector: jsonapi.NewConnector(client),
	}
}
