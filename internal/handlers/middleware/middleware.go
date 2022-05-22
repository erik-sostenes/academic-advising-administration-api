package middleware

import (
	"net/http"

	"github.com/itsoeh/academic-advising-administration-api/internal/model"
	"github.com/itsoeh/academic-advising-administration-api/internal/services"
	"github.com/labstack/echo/v4"
)

// Authentication middleware that validates the user's token
func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("token")
		token := services.Token{}

		if err := token.ValidateToken(tokenString); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, model.Response{"error" : err.Error()})
		}

		return next(c)
	}
}
