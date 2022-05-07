package dependency

import (
	"os"
	"strings"

	"github.com/itsoeh/academy-advising-administration-api/internal/repository"
	"github.com/itsoeh/academy-advising-administration-api/internal/repository/schedule"
	"github.com/itsoeh/academy-advising-administration-api/internal/server"
	"github.com/itsoeh/academy-advising-administration-api/internal/services"
)

const defaultPort = ":9090"

// Run method that is responsible for injecting dependencies
func Run() error {
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		port = defaultPort 
	}

	if _, err := repository.LoadSqlConnection(); err != nil {
		return err
	}

	// inject dependencies 	
	DB := repository.NewDB()
	r := schedule.NewScheduleStorer(DB)
	s := services.NewScheduleService(r)

	start := server.NewServer(port, s)

	return start.Run()
}
