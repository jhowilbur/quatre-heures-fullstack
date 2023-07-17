package event

import (
	"github.com/jhowilbur/event-backend/pkg/entity"
	"github.com/jhowilbur/event-backend/pkg/operation"
	"github.com/jhowilbur/event-backend/pkg/storage"
)

type defaultEventProcess struct {
	event storage.EventRepository
}

func NewDefaultEventProcess(event storage.EventRepository) operation.EventProcess {
	return defaultEventProcess{
		event: event,
	}
}

func (d defaultEventProcess) ProcessRecoverById(eventId uint) (*entity.Event, error) {
	return d.event.FindOneById(eventId)
}

func (d defaultEventProcess) ProcessRecoverAll() ([]entity.Event, error) {
	return d.event.FindAll()
}

func (d defaultEventProcess) ProcessRegister(event *entity.NewEvent) error {
	return d.event.Create(event)
}
