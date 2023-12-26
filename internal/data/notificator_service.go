package data

import "net/url"

type NotificatorsStorage interface {
	Add(svc *NotificatorService) error
	GetByChannel(channel string) (*NotificatorService, error)
	Channels() ([]string, error)
}

type NotificatorService struct {
	Endpoint *url.URL
	Channels []string
}
