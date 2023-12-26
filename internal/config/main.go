package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer

	Notificator
	RegistrationAPIer
	S3
}

type config struct {
	getter kv.Getter

	comfig.Logger
	pgdb.Databaser
	comfig.Listenerer

	Notificator
	RegistrationAPIer
	S3
}

func New(getter kv.Getter) Config {
	return &config{
		getter: getter,

		Databaser:  pgdb.NewDatabaser(getter),
		Listenerer: comfig.NewListenerer(getter),
		Logger:     comfig.NewLogger(getter, comfig.LoggerOpts{}),

		Notificator:       NewNotificator(getter),
		RegistrationAPIer: NewRegistrationAPIer(getter),
		S3:                NewS3(getter),
	}
}
