package http

import (
	"github.com/jhowilbur/event-backend/pkg/entity"
	"github.com/jhowilbur/event-backend/pkg/operation"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// EventRouting is the interface that wraps the event routing methods.
type EventRouting interface {
	RoutingRecoverById(ctx echo.Context) error
	RoutingRecoverAll(ctx echo.Context) error
	RoutingRegister(ctx echo.Context) error
}

// defaultEventRouting is the default implementation of the event routing.
type defaultEventRouting struct {
	event operation.EventProcess
}

// NewEventRouting returns the default implementation of the event routing.
func NewEventRouting(event operation.EventProcess) EventRouting {
	return defaultEventRouting{
		event: event,
	}
}

func (d defaultEventRouting) RoutingRecoverById(ctx echo.Context) error {
	eventId, _ := strconv.Atoi(ctx.Param("id"))
	processRecover, err := d.event.ProcessRecoverById(uint(eventId))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error to recover event")
	} else if processRecover == nil {
		return ctx.JSON(http.StatusNotFound, "Event not found")
	}

	return ctx.JSON(http.StatusOK, processRecover)
}

func (d defaultEventRouting) RoutingRecoverAll(ctx echo.Context) error {
	processRecover, err := d.event.ProcessRecoverAll()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error to recover events")
	} else if processRecover == nil {
		return ctx.JSON(http.StatusNotFound, "Events not found")
	}

	return ctx.JSON(http.StatusOK, processRecover)
}

func (d defaultEventRouting) RoutingRegister(ctx echo.Context) error {
	event := new(entity.NewEvent)

	err := ctx.Bind(event)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error to bind event")
	}

	err = d.event.ProcessRegister(event)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Error to register event")
	}

	return ctx.NoContent(http.StatusCreated)
}
