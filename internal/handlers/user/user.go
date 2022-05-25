package user

import (
	"net/http"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/services/user"
	"github.com/labstack/echo/v4"
)

// UserHandler contains all the http handlers to receive user login requests and responses
type UserHandler interface {
// StudentLoginHandler http handler which is responsible for student login via http request
	StudentLoginHandler(user.UserService) echo.HandlerFunc
// TeacherLoginHandler http handler which is responsible for teacher login via http request
	TeacherLoginHandler(user.UserService) echo.HandlerFunc
}
// userHandler implements the UserHandler interface
type userHandler struct {}

// NewUserHandler returns the default UserHandler interface implementation
func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (*userHandler) StudentLoginHandler(services user.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCredentials, err :=  services.GetStudentPasswordByTuition(c.Request().Context(), c.Param("tuition"), c.Param("email"), c.Param("password"))	

		if _, ok := err.(model.InternalServerError); ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.StatusBadRequest); ok {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.NotFound); ok {
			return echo.NewHTTPError(http.StatusNotFound, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.Forbidden); ok {
			return echo.NewHTTPError(http.StatusNotFound, model.Response{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, model.Response{"data": userCredentials})	
	}
}

func (*userHandler) TeacherLoginHandler(services user.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCredentials, err :=  services.GetTeacherPasswordByTuition(c.Request().Context(), c.Param("tuition"), c.Param("email"), c.Param("password"))	

		if _, ok := err.(model.InternalServerError); ok {
			return echo.NewHTTPError(http.StatusInternalServerError, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.StatusBadRequest); ok {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.NotFound); ok {
			return echo.NewHTTPError(http.StatusNotFound, model.Response{"error": err.Error()})
		}
		if _, ok := err.(model.Forbidden); ok {
			return echo.NewHTTPError(http.StatusNotFound, model.Response{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, model.Response{"data": userCredentials})	
	}
}
