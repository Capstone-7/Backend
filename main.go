package main

import (
	route "capstone/app/routes"
	drivers "capstone/drivers"
	mongo_driver "capstone/drivers/mongo"
	utils "capstone/utils"
	"fmt"

	_productsUseCase "capstone/businesses/products"
	_userUseCase "capstone/businesses/users"
	_productController "capstone/controllers/products"
	_userController "capstone/controllers/users"

	_transactionUseCase "capstone/businesses/transactions"
	_transactionController "capstone/controllers/transactions"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Setup DB
	client, err := mongo_driver.Connect()
	if err != nil {
		panic(err)
	}
	mongo_driver.SetClient(client)

	e := echo.New()

	//product
	productRepo := drivers.NewProductRepository(mongo_driver.GetDB())
	productUseCase := _productsUseCase.NewProductUseCase(productRepo)
	productController := _productController.NewProductController(productUseCase)

	// User
	userRepo := drivers.NewUserRepository(mongo_driver.GetDB())
	otpRepo := drivers.NewOTPRepository(mongo_driver.GetDB())
	userUseCase := _userUseCase.NewUserUseCase(userRepo, otpRepo)
	userController := _userController.NewUserController(userUseCase)

	// Transaction
	transactionRepo := drivers.NewTransactionRepository(mongo_driver.GetDB())
	transactionUseCase := _transactionUseCase.NewTransactionUseCase(transactionRepo)
	transactionController := _transactionController.NewTransactionController(transactionUseCase, productUseCase, userUseCase)

	
	// Init routes
	appRoute := route.ControllerList{
		UserController: *userController,
		ProductController: *productController,
		TransactionController: *transactionController,
	}
	appRoute.Init(e)

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Start in HTTPS mode
	fmt.Println("Starting server...")
	appPort := ":"+utils.ReadENV("APP_PORT")
	e.StartTLS(appPort, "cert.pem", "key.pem")
}