package server

import (
	"github.com/itsoeh/academy-advising-administration-api/internal/handlers"
	a "github.com/itsoeh/academy-advising-administration-api/internal/handlers/middleware"
	"github.com/itsoeh/academy-advising-administration-api/internal/services/schedule"
	"github.com/itsoeh/academy-advising-administration-api/internal/services/user"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

type server struct {
	port string
	engine *echo.Echo
	scheduleService schedule.ScheduleService
	userService user.UserService
}

// NewServer to start the server
func NewServer(port string, schedule schedule.ScheduleService, user user.UserService) server {
	s :=  server{
		port: port,
		engine: echo.New(),
		scheduleService: schedule,
		userService: user,
	}

	s.SetAllEndpoints()

	return s
}

// Run will start running the program on the defined port
func (s *server) Run() error {
	return  s.engine.Start(s.port)
}

// SetAllEndpoints contains all endpoints
func (s *server) SetAllEndpoints() {
	h := handlers.NewHandlers()

	// Add middlewares 
	s.engine.Use(m.Logger(), m.Recover(), m.CORS())

	route := s.engine.Group("/v1/itsoeh/academic-advising-administration-api")

	route.POST("/create", a.Authentication(h.Schedule.HandlerCreateTeacherSchedule(s.scheduleService)))
	route.GET("/get/:teacher_id/:is_active", a.Authentication(h.Schedule.HandlerGetTeacherSchedule(s.scheduleService)))

	route.GET("/student-authorization", h.User.StudentLoginHandler(s.userService))
	route.GET("/teacher-authorization", h.User.TeacherLoginHandler(s.userService))
}
