package memory

import (
	"errors"
	"sync"

	"github.com/apodeixis/notifications-router-svc/internal/data"
)

func NewMemoryNotificatorsStorage() data.NotificatorsStorage {
	return &memoryNotificatorsStorage{}
}

type memoryNotificatorsStorage struct {
	storage sync.Map
}

func (s *memoryNotificatorsStorage) Add(svc *data.NotificatorService) error {
	for _, channel := range svc.Channels {
		s.storage.Store(channel, svc)
	}
	return nil
}

func (s *memoryNotificatorsStorage) GetByChannel(channel string) (*data.NotificatorService, error) {
	service, ok := s.storage.Load(channel)
	if !ok {
		return nil, errors.New("service for this channel is not registered")
	}
	return service.(*data.NotificatorService), nil
}

func (s *memoryNotificatorsStorage) Channels() ([]string, error) {
	keys := make([]string, 0)
	s.storage.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys, nil
}
