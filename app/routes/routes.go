package route

import "github.com/labstack/echo/v4"

type ControllerList struct {
}

func (cl *ControllerList) Init(e *echo.Echo) {
	// Heartbeat
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Alive!")
	})

	// apiV1 := e.Group("/api/v1")
}