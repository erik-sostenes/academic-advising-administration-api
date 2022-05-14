package teacher

import (
	"net/http"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"github.com/itsoeh/academy-advising-administration-api/internal/services/teacher"
	"github.com/labstack/echo/v4"
)
// TeacherHandler contains all the http handlers to receive requests and responses from the teacher lookup
type TeacherHandler interface {
	// HandlerFindTeachersByCareerAndSubject http controller which is responsible for looking up teachers through an http request 
	HandlerFindTeachersByCareerAndSubject(teacher.TeacherService) echo.HandlerFunc
}
// teacherHandler implements the TeacherHandler interface
type teacherHandler struct{}

// NewTeacherHandler returns the default TeacherHandler interface implementation
func NewTeacherHandler() TeacherHandler {
	return &teacherHandler{}
}

func (*teacherHandler) HandlerFindTeachersByCareerAndSubject(services teacher.TeacherService) echo.HandlerFunc {
	return func(c echo.Context) error {
		teacherCards, err := services.FindTeachers(c.Request().Context(), c.Param("subject_id"), c.Param("university_course_id"))	

		if _, ok := err.(model.StatusBadRequest); ok {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": err.Error()})
		}

		if _, ok := err.(model.InternalServerError); ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.Response{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, model.Response{"data": teacherCards})
	}
}

