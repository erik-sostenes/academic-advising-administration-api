package server

import (
	"log"

	"github.com/itsoeh/academy-advising-administration-api/internal/handlers"
	"github.com/itsoeh/academy-advising-administration-api/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	port string
	engine *echo.Echo
	services services.ScheduleService
}

// NewServer to start the server
func NewServer(port string, services services.ScheduleService) server {
	s :=  server{
		port: port,
		engine: echo.New(),
		services: services,
	}

	s.SetAllEndpoints()

	return s
}

// Run will start running the program on the defined port
func (s *server) Run() error {
	log.Printf("Initialize server on the port %v", s.port)

	return  s.engine.Start(s.port)
}

// SetAllEndpoints contains all endpoints
func (s *server) SetAllEndpoints() {
	h := handlers.NewHandlers()

	// Add middlewares 
	s.engine.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	route := s.engine.Group("/v1/itsoeh/academic-advising-administration-api")

	route.POST("/create", h.Schedule.CreateHandler(s.services))
	route.GET("/get/:teacher_id/:is_active", h.Schedule.GetHandler(s.services))
}
