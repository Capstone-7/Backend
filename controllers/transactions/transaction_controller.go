package transactions

import (
	"capstone/businesses/products"
	"capstone/businesses/transactions"
	"capstone/businesses/users"
	"capstone/controllers/transactions/requests"
	"capstone/controllers/transactions/response"
	"capstone/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	appjwt "capstone/utils/jwt"

	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
)

type TransactionController struct {
	TransactionUseCase transactions.UseCase
	ProductUseCase    products.UseCase
	UserUseCase       users.UseCase
}

func NewTransactionController(transactionUseCase transactions.UseCase, productUseCase products.UseCase, userUseCase users.UseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
		ProductUseCase:    productUseCase,
		UserUseCase:       userUseCase,
	}
}

func (t *TransactionController) ReviewTransaction(c echo.Context) error {
	request := requests.ReviewTransactionRequest{}
	c.Bind(&request)

	// Get user id from token
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id from token
	user_id := appjwt.GetID(tokenString)

	// Get user data
	user, err := t.UserUseCase.GetByID(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: "Please login again",
			Data:    nil,
		})
	}
	
	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Get Product
	product, err := t.ProductUseCase.GetProductByCode(request.ProductCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response{
			Status:  http.StatusNotFound,
			Message: "Product not found",
			Data:    nil,
		})
	}

	// Init transaction
	transaction := t.TransactionUseCase.InitTransaction(&product, &user, request.CustomerID)
	response := response.ReviewTransaction{
		ProductCode: product.Code,
		CustomerID: request.CustomerID,
		ProductDescription: product.Description,
		ProductPrice: product.Price,
		AdminFee: transaction.AdminFee,
		TotalPrice: transaction.TotalPrice,
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

func (t *TransactionController) SubmitTransaction(c echo.Context) error {
	request := requests.SubmitTransactionRequest{}
	c.Bind(&request)

	// Get user id from token
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id from token
	user_id := appjwt.GetID(tokenString)

	// Get user data
	user, err := t.UserUseCase.GetByID(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: "Please login again",
			Data:    nil,
		})
	}

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Get Product
	product, err := t.ProductUseCase.GetProductByCode(request.ProductCode)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response{
			Status:  http.StatusNotFound,
			Message: "Product not found",
			Data:    nil,
		})
	}

	// Submit Transaction
	transaction, err := t.TransactionUseCase.CreateTransaction(&product, &user, &request)
	fmt.Println("error", err)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    transaction,
	})
}

func (t *TransactionController) XenditCallback(c echo.Context) error {
	fmt.Println("CALLBACK CALLED")
	// Get request body
	request := xendit.Invoice{}
	c.Bind(&request)

	// print json request
	json_req, _ := json.Marshal(request)
	fmt.Println("JSON REQUEST", string(json_req))

	// Get transaction
	transaction, err := t.TransactionUseCase.GetTransactionByXenditInvoiceID(request.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response{
			Status:  http.StatusNotFound,
			Message: "Transaction not found",
			Data:    nil,
		})
	}

	// Change payment status
	transaction, err = t.TransactionUseCase.UpdatePaymentInfo(transaction.ID.Hex(), request.Status, request.PaymentMethod, request.PaymentChannel)

	// Get product
	product, err := t.ProductUseCase.GetProductByCode(transaction.ProductID.Hex())
	
	// handle special case for manual update status
	if product.Code != "SPECIALCASE" {
		transaction, err = t.TransactionUseCase.ChangeTransactionStatus(transaction.ID.Hex(), "SUCCESS")
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    transaction,
	})
}

// Get Transaction History by user id
func (t *TransactionController) GetTransactionHistory(c echo.Context) error {
	// Get user id from token
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id from token
	user_id := appjwt.GetID(tokenString)

	// Get user data
	user, err := t.UserUseCase.GetByID(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: "Please login again",
			Data:    nil,
		})
	}

	// Get transaction history
	transactions, err := t.TransactionUseCase.GetTransactionHistoryByUserID(user.ID.Hex())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    transactions,
	})
}

// GetTransactionHistoryByID
func (t *TransactionController) GetTransactionHistoryByID(c echo.Context) error {
	transactionID := c.Param("id")

	// Get user id from token
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id from token
	user_id := appjwt.GetID(tokenString)

	// Get user data
	user, err := t.UserUseCase.GetByID(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: "Please login again",
			Data:    nil,
		})
	}

	// Get transaction history
	transaction, err := t.TransactionUseCase.GetTransactionHistoryByID(transactionID, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    transaction,
	})
}

// Change Transaction Status
func (t *TransactionController) ChangeTransactionStatus(c echo.Context) error {
	transactionID := c.Param("id")

	// Get request body
	request := requests.ChangeTransactionStatus{}
	c.Bind(&request)

	// Update transaction status
	transaction, err := t.TransactionUseCase.ChangeTransactionStatus(transactionID, request.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    transaction,
	})
}

// Get All Transaction By Admin
func (t *TransactionController) GetAllTransaction(c echo.Context) error {
	// Get query params
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	status := c.QueryParam("status")
	
	// Get transaction history
	transactions, err := t.TransactionUseCase.GetAllTransaction(int64(page), int64(limit), status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    transactions,
	})
}

// Get Total Transaction
func (t *TransactionController) GetTotalTransaction(c echo.Context) error {
	// Get total transaction
	total, err := t.TransactionUseCase.GetTotalTransaction()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    total,
	})
}

// Get Top Products By Category
func (t *TransactionController) GetTopProductsByCategory(c echo.Context) error {
	// Get top products by category
	topProducts, err := t.TransactionUseCase.GetTopProductsByCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    topProducts,
	})
}

// Get Income Per Day
func (t *TransactionController) GetIncomePerDay(c echo.Context) error {
	// Get income per day
	income, err := t.TransactionUseCase.GetIncomePerDay()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    income,
	})
}