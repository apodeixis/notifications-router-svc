package service

import (
	"context"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/apodeixis/notifications-router-svc/internal/config"
	"github.com/apodeixis/notifications-router-svc/internal/data/memory"
	"github.com/apodeixis/notifications-router-svc/internal/service/api"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/primary"
	"github.com/apodeixis/notifications-router-svc/internal/service/api/registration"
	"github.com/apodeixis/notifications-router-svc/internal/service/runners"
	"github.com/apodeixis/notifications-router-svc/internal/service/runners/processor"
)

type Service struct {
	primaryAPI      api.API
	registrationAPI api.API
	processor       runners.Runner
}

func New(cfg config.Config) *Service {
	storage := memory.NewMemoryNotificatorsStorage()
	return &Service{
		primaryAPI:      primary.NewAPI(cfg),
		registrationAPI: registration.NewAPI(cfg, storage),
		processor:       processor.New(cfg, storage),
	}
}

func (s *Service) Run(ctx context.Context) (executionError error) {
	go func() {
		if err := s.primaryAPI.Start(); err != nil {
			executionError = errors.Wrap(err, "primary API exited with error")
		}
	}()

	go func() {
		if err := s.registrationAPI.Start(); err != nil {
			executionError = errors.Wrap(err, "registration API exited with error")
		}
	}()

	if err := s.processor.Run(ctx); err != nil {
		executionError = errors.Wrap(err, "processor runner exited with error")
	}

	return executionError
}
