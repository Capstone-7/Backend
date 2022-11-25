package route

import (
	"capstone/app/middlewares"
	"capstone/controllers/products"
	"capstone/controllers/transactions"
	"capstone/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	ProductController products.ProductController
	UserController users.UserController
	TransactionController transactions.TransactionController
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

 	// products
	product := apiV1.Group("/product")
	product.GET("", cl.ProductController.GetAll, admin.Middleware)
	product.GET("/all", cl.ProductController.GetAll, admin.Middleware)
	product.GET("/count", cl.ProductController.GetTotalProducts, admin.Middleware)
	product.POST("", cl.ProductController.Create, admin.Middleware)
	product.GET("/categories", cl.ProductController.GetCategoryList, middlewares.AuthMiddleware)
	product.GET("/categories/:product_type", cl.ProductController.GetCategoriesByProductType, middlewares.AuthMiddleware)

	product.GET("/by_category/:category", cl.ProductController.GetProductsByCategory, middlewares.AuthMiddleware)
	product.GET("/by_type/:product_type", cl.ProductController.GetProductsByProductType, middlewares.AuthMiddleware)

	product.GET("/:id", cl.ProductController.GetProductByID)
	product.PUT("/:id", cl.ProductController.UpdateProduct, admin.Middleware)
	product.DELETE("/:id", cl.ProductController.DeleteProduct, admin.Middleware)
  
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
	users.GET("/count", cl.UserController.GetTotalUsers, admin.Middleware)

	users.GET("", cl.UserController.GetAllUsers, admin.Middleware)
	users.GET("/all", cl.UserController.GetAllUsers, admin.Middleware)
	users.GET("/:id", cl.UserController.GetUserByID, admin.Middleware)
	users.PUT("/:id", cl.UserController.UpdateUserByID, admin.Middleware)
	users.DELETE("/:id", cl.UserController.DeleteUserByID, admin.Middleware)

	// Transaction
	transactions := apiV1.Group("/transaction")
	transactions.POST("/review", cl.TransactionController.ReviewTransaction, middlewares.AuthMiddleware)
	transactions.POST("/submit", cl.TransactionController.SubmitTransaction, middlewares.AuthMiddleware)
	transactions.GET("/history", cl.TransactionController.GetTransactionHistory, middlewares.AuthMiddleware)
	transactions.GET("/history/all", cl.TransactionController.GetAllTransaction, admin.Middleware)
	transactions.GET("/history/:id", cl.TransactionController.GetTransactionHistoryByID, middlewares.AuthMiddleware)
	transactions.GET("/count", cl.TransactionController.GetTotalTransaction, admin.Middleware)
	transactions.PUT("/:id", cl.TransactionController.ChangeTransactionStatus, admin.Middleware)
	// callback
	transactions.POST("/callback", cl.TransactionController.XenditCallback)
}