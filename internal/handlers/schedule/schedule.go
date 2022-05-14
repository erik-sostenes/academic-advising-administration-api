package schedule

import (
	"net/http"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/services/schedule"
	"github.com/labstack/echo/v4"
)

// ScheduleHandler contains all http handlers to receive requests and responses from schedules
type ScheduleHandler interface {
	// HandlerCreateTeacherSchedule http handler which is responsible for creating a teacher schedule via request
	HandlerCreateTeacherSchedule(schedule.ScheduleService) echo.HandlerFunc
	// HandlerGetTeacherSchedule http that is responsible for responding to all the schedules that a teacher has
	HandlerGetTeacherSchedule(schedule.ScheduleService) echo.HandlerFunc
}

type scheduleHandler struct {}

// NewScheduleHandler returns the ScheduleHandler interface with all its http methods.
func NewScheduleHandler() ScheduleHandler {
	return &scheduleHandler{}
}

func (*scheduleHandler) HandlerCreateTeacherSchedule(services schedule.ScheduleService) echo.HandlerFunc {
	return func(c echo.Context) error {
		schedule := &model.Schedule{}

		if err := c.Bind(schedule); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": "The schedule structure is wrong."})
		}

		err := services.CreateSchedule(c.Request().Context(),
			schedule.ScheduleAt, schedule.FromDate, schedule.ToDate, schedule.TeacherTuition,
		)

		if _, ok := err.(model.StatusBadRequest); ok {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": err.Error()})
		}

		if _, ok := err.(model.InternalServerError); ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.Response{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, model.Response{"message": "The schedule has been created successfully."})
	}
} 

func (*scheduleHandler) HandlerGetTeacherSchedule(services schedule.ScheduleService) echo.HandlerFunc {
	return func(c echo.Context) error {
		teacherSchedules, err := services.GetSchedulesByTeacherTuition(c.Request().Context(), c.Param("teacher_id"), c.Param("is_active"))

		if _, ok := err.(model.NotFound); ok {
			return echo.NewHTTPError(http.StatusNotFound, model.Response{"error": err.Error()})
		}

		if _, ok := err.(model.InternalServerError); ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.StatusBadRequest); ok {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, model.Response{"data": teacherSchedules})
	}
}
