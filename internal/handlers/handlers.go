package handlers

import "github.com/itsoeh/academy-advising-administration-api/internal/handlers/schedule"

// Handlers structure that manages the handlers
type handlers struct {
	Schedule schedule.ScheduleHandler
}

// NewHandlers returns a handler struct that contains all the handlers from schedules and teacher schedules
func NewHandlers() *handlers {
	return  &handlers{
		Schedule: schedule.NewScheduleHandler(),
	}
}
