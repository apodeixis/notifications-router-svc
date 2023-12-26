package config

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const notificatorConfigKey = "notificator"

type NotificatorConfig struct {
	DefaultChannelsPriority []string `fig:"default_channels_priority,required"`
	DefaultLocale           string   `fig:"default_locale"`
}

type Notificator interface {
	NotificatorConfig() *NotificatorConfig
}

func NewNotificator(getter kv.Getter) Notificator {
	return &notificator{
		getter: getter,
	}
}

type notificator struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *notificator) NotificatorConfig() *NotificatorConfig {
	return c.once.Do(func() interface{} {
		cfg := &NotificatorConfig{
			DefaultLocale: "en",
		}
		err := figure.
			Out(cfg).
			From(kv.MustGetStringMap(c.getter, notificatorConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out notificator config"))
		}
		return cfg
	}).(*NotificatorConfig)
}
