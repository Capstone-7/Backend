package main

import (
	route "capstone/app/routes"
	drivers "capstone/drivers"
	mongo_driver "capstone/drivers/mongo"
	utils "capstone/utils"
	"fmt"

	_productsUseCase "capstone/businesses/products"
	_productController "capstone/controller/products"
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

	//product
	productRepo := drivers.NewProductRepository(mongo_driver.GetDB())
	productUseCase := _productsUseCase.NewProductUseCase(productRepo)
	productController := _productController.NewProductController(productUseCase)
	
	// Init routes
	appRoute := route.ControllerList{
		ProductController: *productController,
	}
	appRoute.Init(e)



	// Start in HTTPS mode
	fmt.Println("Starting server...")
	appPort := ":"+utils.ReadENV("APP_PORT")
	e.StartTLS(appPort, "cert.pem", "key.pem")
}