package handlers

import (
	"github.com/itsoeh/academic-advising-administration-api/internal/handlers/schedule"
	"github.com/itsoeh/academic-advising-administration-api/internal/handlers/teacher"
	"github.com/itsoeh/academic-advising-administration-api/internal/handlers/user"

)

// Handlers structure that manages the handlers
type handlers struct {
	Schedule schedule.ScheduleHandler
	Teacher teacher.TeacherHandler
	User user.UserHandler
}

// NewHandlers returns a handler struct that contains all the handlers from schedules and teacher schedules
func NewHandlers() *handlers {
	return  &handlers{
		Schedule: schedule.NewScheduleHandler(),
		Teacher: teacher.NewTeacherHandler(),
		User: user.NewUserHandler(),
	}
}
