package main

import (
	route "capstone/app/routes"
	mongo_driver "capstone/drivers/mongo"
	utils "capstone/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// Setup DB
	// Setup DB
	client, err := mongo_driver.Connect()
	if err != nil {
		panic(err)
	}
	mongo_driver.SetClient(client)

	e := echo.New()
	
	// Init routes
	appRoute := route.ControllerList{}
	appRoute.Init(e)
	
	// Start in HTTPS mode
	fmt.Println("Starting server...")
	appPort := ":"+utils.ReadENV("APP_PORT")
	e.StartTLS(appPort, "cert.pem", "key.pem")
}