package handlers

import (
	"github.com/itsoeh/academy-advising-administration-api/internal/handlers/schedule"
	"github.com/itsoeh/academy-advising-administration-api/internal/handlers/user"
)

// Handlers structure that manages the handlers
type handlers struct {
	Schedule schedule.ScheduleHandler
	User user.UserHandler
}

// NewHandlers returns a handler struct that contains all the handlers from schedules and teacher schedules
func NewHandlers() *handlers {
	return  &handlers{
		Schedule: schedule.NewScheduleHandler(),
		User: user.NewUserHandler(),
	}
}
