package config

import (
	"net"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const registrationApiConfigKey = "registration_api"

type registrationAPIConfig struct {
	Addr string `fig:"addr,required"`
}

type RegistrationAPIer interface {
	RegistrationAPIListener() net.Listener
}

func NewRegistrationAPIer(getter kv.Getter) RegistrationAPIer {
	return &registrationAPIer{
		getter: getter,
	}
}

type registrationAPIer struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *registrationAPIer) RegistrationAPIListener() net.Listener {
	return c.once.Do(func() interface{} {
		cfg := new(registrationAPIConfig)
		err := figure.
			Out(cfg).
			From(kv.MustGetStringMap(c.getter, registrationApiConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out registration api config"))
		}
		listener, err := net.Listen("tcp", cfg.Addr)
		if err != nil {
			panic(errors.Wrap(err, "failed to listen on registration api addr"))
		}
		return listener
	}).(net.Listener)
}
