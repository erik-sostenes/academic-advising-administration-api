package user

import (
	"net/http"

	"github.com/itsoeh/academy-advising-administration-api/internal/model"
	"github.com/itsoeh/academy-advising-administration-api/internal/services/user"
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
		userLogin := &model.Login{}
		
		if err := c.Bind(userLogin); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": "The user login structure is wrong."})
		}

		stringToken, err :=  services.GetStudentPasswordByTuition(c.Request().Context(), userLogin.Tuition, userLogin.Email, userLogin.Password)	

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
		return c.JSON(http.StatusOK, model.Response{"token": stringToken})	
	}
}

func (*userHandler) TeacherLoginHandler(services user.UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		userLogin := &model.Login{}
		
		if err := c.Bind(userLogin); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, model.Response{"error": "The user login structure is wrong."})
		}

		stringToken, err :=  services.GetStudentPasswordByTuition(c.Request().Context(), userLogin.Tuition, userLogin.Email, userLogin.Password)	

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
		return c.JSON(http.StatusOK, model.Response{"token": stringToken})	
	}
}
