package entity

import "time"

// Event is the model for the event
type Event struct {
	ID          uint      `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	InitialDate time.Time `json:"initial_date" db:"initialdate"`
	FinalDate   time.Time `json:"final_date" db:"finaldate"`
}

// NewEvent is for registering a new event
type NewEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	InitialDate time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
}
