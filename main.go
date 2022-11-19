package main

import (
	route "capstone/app/routes"
	drivers "capstone/drivers"
	mongo_driver "capstone/drivers/mongo"
	utils "capstone/utils"
	"fmt"

	_userUseCase "capstone/businesses/users"
	_userController "capstone/controllers/users"

	"github.com/labstack/echo/v4"
)

func main() {
	// Setup DB
	client, err := mongo_driver.Connect()
	if err != nil {
		panic(err)
	}
	mongo_driver.SetClient(client)

	e := echo.New()

	// User
	userRepo := drivers.NewUserRepository(mongo_driver.GetDB())
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userController := _userController.NewUserController(userUseCase)
	
	// Init routes
	appRoute := route.ControllerList{
		UserController: *userController,
	}
	appRoute.Init(e)
	
	// Start in HTTPS mode
	fmt.Println("Starting server...")
	appPort := ":"+utils.ReadENV("APP_PORT")
	e.StartTLS(appPort, "cert.pem", "key.pem")
}