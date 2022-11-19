package route

import (
	"capstone/app/middlewares"
	"capstone/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) Init(e *echo.Echo) {
	// Setup midllewares
	admin := middlewares.Roles{
		Role: []string{"admin"},
	}

	// Heartbeat
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Alive!")
	})

	apiV1 := e.Group("/api/v1")

	// User
	users := apiV1.Group("/user")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)
	users.GET("/profile", cl.UserController.GetProfile, middlewares.AuthMiddleware)
	users.PUT("/profile", cl.UserController.UpdateProfile, middlewares.AuthMiddleware)
	users.PUT("/update-password", cl.UserController.UpdatePassword, middlewares.AuthMiddleware)
	
	users.POST("/request-otp", cl.UserController.RequestOTP)
	users.POST("/verify-email", cl.UserController.VerifyEmail)
	users.POST("/reset-password", cl.UserController.ResetPassword)

	users.GET("/all", cl.UserController.GetAllUsers, admin.Middleware)
	users.GET("/:id", cl.UserController.GetUserByID, admin.Middleware)
	users.PUT("/:id", cl.UserController.UpdateUserByID, admin.Middleware)
	users.DELETE("/:id", cl.UserController.DeleteUserByID, admin.Middleware)
}