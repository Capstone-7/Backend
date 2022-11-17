package middlewares

import (
	appjwt "capstone/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type Roles struct {
	Role []string
}

func (r *Roles) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get token from header
		tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
		
		if tokenString == "" {
			return echo.ErrUnauthorized
		}

		// Validate token
		err := appjwt.ValidateToken(tokenString)
		if err != nil {
			return echo.ErrUnauthorized
		}

		// Get user roles
		roles := appjwt.GetRoles(tokenString)
		for _, role := range roles {
			for _, allowedRole := range r.Role {
				if role == allowedRole {
					return next(c)
				}
			}
		}

		return echo.ErrUnauthorized
	}
}