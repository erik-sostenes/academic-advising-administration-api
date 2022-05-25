package server

import (
	"github.com/itsoeh/academic-advising-administration-api/internal/handlers"
	a "github.com/itsoeh/academic-advising-administration-api/internal/handlers/middleware"
	"github.com/itsoeh/academic-advising-administration-api/internal/services/schedule"
	"github.com/itsoeh/academic-advising-administration-api/internal/services/teacher"
	"github.com/itsoeh/academic-advising-administration-api/internal/services/user"
	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

type server struct {
	port string
	engine *echo.Echo
	scheduleService schedule.ScheduleService
	userService user.UserService
	teacherService teacher.TeacherService
}

// NewServer to start the server
func NewServer(port string, schedule schedule.ScheduleService, user user.UserService, teacher teacher.TeacherService) server {
	s :=  server{
		port: port,
		engine: echo.New(),
		scheduleService: schedule,
		userService: user,
		teacherService: teacher,	
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

	// user login
	route.GET("/student-authorization/:tuition/:email/:password",
		h.User.StudentLoginHandler(s.userService),
	)
	route.GET("/teacher-authorization/:tuition/:email/:password",
		h.User.TeacherLoginHandler(s.userService),
	)
	// schedule teacher
	route.GET("/find-teachers/:subject_id/:university_course_id", a.Authentication(
		h.Teacher.HandlerFindTeachersByCareerAndSubject(s.teacherService),
	))
	// teachers
	route.GET("/student-requests/:teacher_tuition", a.Authentication(
		h.Teacher.HandlerFindStudentRequests(s.teacherService),
	))

	route.GET("/student-requests-accepted/:teacher_tuition", a.Authentication(
		h.Teacher.HandlerFindStudentRequestsAccepted(s.teacherService),
	))
	route.POST("/create", a.Authentication(
		h.Schedule.HandlerCreateTeacherSchedule(s.scheduleService)),
	)
	route.GET("/get/:teacher_id/:is_active", a.Authentication(
		h.Schedule.HandlerGetTeacherSchedule(s.scheduleService)),
	)
}
