package dependency

import (
	"os"
	"strings"

	"github.com/itsoeh/academy-advising-administration-api/internal/repository"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/schedule"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/teacher"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/user"
	"github.com/itsoeh/academy-advising-administration-api/internal/server"
	"github.com/itsoeh/academy-advising-administration-api/internal/services"
	s "github.com/itsoeh/academy-advising-administration-api/internal/services/schedule"
	u "github.com/itsoeh/academy-advising-administration-api/internal/services/user"
	t "github.com/itsoeh/academy-advising-administration-api/internal/services/teacher"
)

const defaultPort = ":9090"

// Run method that is responsible for injecting dependencies
func Run() error {
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		port = defaultPort 
	}
	
	certifier := services.Certifier{}
	if err := certifier.Certificates("../internal/utils/app.rsa.pub", "../internal/utils/app.rsa"); err != nil {
		return err
	}
	if _, err := repository.LoadSqlConnection(); err != nil {
		return err
	}

	// inject dependencies 	
	DB := repository.NewDB()
	scheduleStorer := schedule.NewScheduleStorer(DB)
	userStorer := user.NewUserStorer(DB)
	teacherStorer := teacher.NewTeacherStorer(DB)

	scheduleService := s.NewScheduleService(scheduleStorer)
	userService := u.NewUserService(userStorer)
	teacherService := t.NewTeacherService(teacherStorer)
	start := server.NewServer(port, scheduleService, userService, teacherService)

	return start.Run()
}
