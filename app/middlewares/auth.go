package middlewares

import (
	"capstone/helpers"
	appjwt "capstone/utils/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get token from header
		tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helpers.Response{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
		}

		// Validate token
		err := appjwt.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.Response{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Data:    nil,
			})
		}

		return next(c)
	}
}