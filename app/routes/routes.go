package route

import (
	"capstone/controller/products"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	ProductController products.ProductController
}

func (cl *ControllerList) Init(e *echo.Echo) {
	// Heartbeat
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Alive!")
	})

	// apiV1 := e.Group("/api/v1")

	// products
	product := e.Group("/product")
	product.GET("", cl.ProductController.GetAll)
	product.POST("", cl.ProductController.Create)
	product.GET("/:id", cl.ProductController.GetProductByID)
	product.PUT("/:id", cl.ProductController.UpdateProduct)
	product.DELETE("/:id", cl.ProductController.DeleteProduct)
}