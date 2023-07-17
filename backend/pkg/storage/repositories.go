package storage

import (
	"github.com/jhowilbur/event-backend/pkg/entity"
)

// EventRepository interface for event repository
type EventRepository interface {
	FindAll() ([]entity.Event, error)
	FindOneById(id uint) (*entity.Event, error)
	Create(event *entity.NewEvent) error
}
