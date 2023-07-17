package context

import (
	"github.com/jhowilbur/event-backend/pkg/operation"
	"github.com/jhowilbur/event-backend/pkg/operation/event"
	"github.com/jhowilbur/event-backend/pkg/routes/http"
	"github.com/jhowilbur/event-backend/pkg/storage/db"
	"github.com/jmoiron/sqlx"
)

func NewEventRoute(database *sqlx.DB) http.EventRouting {
	return http.NewEventRouting(
		NewEventOperation(database),
	)
}

func NewEventOperation(database *sqlx.DB) operation.EventProcess {
	return event.NewDefaultEventProcess(
		db.NewEventRepository(database),
	)
}
