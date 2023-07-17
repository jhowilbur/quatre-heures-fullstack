package operation

import "github.com/jhowilbur/event-backend/pkg/entity"

type EventProcess interface {
	ProcessRecoverById(eventId uint) (*entity.Event, error)
	ProcessRecoverAll() ([]entity.Event, error)
	ProcessRegister(event *entity.NewEvent) error
}
