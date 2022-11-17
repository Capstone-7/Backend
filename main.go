package main

import (
	route "capstone/app/routes"
	utils "capstone/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	appPort := ":"+utils.ReadENV("APP_PORT")
	
	// Init routes
	appRoute := route.ControllerList{}
	appRoute.Init(e)
	
	// Start server
	fmt.Println("Starting server...")
	
	// Start in HTTPS mode
	e.StartTLS(appPort, "cert.pem", "key.pem")
}