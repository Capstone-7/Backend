package transactions

import (
	"capstone/businesses/products"
	"capstone/businesses/users"
	"capstone/controllers/transactions/requests"
	"capstone/controllers/transactions/response"
	"capstone/helpers"
	"capstone/utils"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionUseCase struct {
	TransactionRepository Repository
}

func NewTransactionUseCase(transactionRepository Repository) UseCase {
	return &TransactionUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (t *TransactionUseCase) InitTransaction(product *products.Domain, user *users.Domain, customer_id string) Domain {
	// Generate random price for product type BILLS
	if product.Type == "BILL" {
		rand.Seed(time.Now().UnixNano())
		product.Price = helpers.PRNG(customer_id)
	}

	// Create Transaction
	transaction := Domain{
		UserID : user.ID,
		ProductID : product.ID,
		ProductPrice: product.Price,
		CustomerID: customer_id,
		AdminFee: 0,
		TotalPrice : product.Price,
		Status : "PENDING",
		Created : primitive.NewDateTimeFromTime(time.Now()),
		Updated : primitive.NewDateTimeFromTime(time.Now()),
		Deleted : primitive.NewDateTimeFromTime(time.Time{}),
	}

	return transaction
}

func (t *TransactionUseCase) CreateTransaction(product *products.Domain, user *users.Domain, req *requests.SubmitTransactionRequest) (Domain, error) {
	// Init Transaction
	transaction := t.InitTransaction(product, user, req.CustomerID)

	// Create XENDIT Invoice
	invoice, err := t.CreateXenditInvoice(product, user, &transaction, req)
	if err != nil {
		return transaction, err
	}

	// Update transaction
	transaction.XenditInvoiceID = invoice.ID
	transaction.XenditPaymentURL = invoice.InvoiceURL
	transaction.XenditStatus = invoice.Status
	transaction.XenditExternalID = invoice.ExternalID

	// Create Transaction
	transaction, err = t.TransactionRepository.Create(&transaction)
	
	return transaction, err
}

func (t *TransactionUseCase) CreateXenditInvoice(product *products.Domain, user *users.Domain, transaction *Domain, req *requests.SubmitTransactionRequest) (*xendit.Invoice, error) {
	// Create Xendit Invoice
	xendit.Opt.SecretKey = utils.ReadENV("XENDIT_SECRET")

	customer := xendit.InvoiceCustomer{
		GivenNames:   user.Name,
		Email:        user.Email,
	}

	item := xendit.InvoiceItem{
		Name:		product.Description,
		Quantity:	1,
		Price:		float64(product.Price),
		Category:	product.Category,
	}

	fee := xendit.InvoiceFee{
		Type:         "ADMIN",
		Value:        0,
	}
	
	// Current timestamp
	timestamp := time.Now().Unix()
	transaction.XenditExternalID = "INV-"+ user.ID.Hex() + "-" + fmt.Sprint(timestamp)
	
	data := invoice.CreateParams{
		ExternalID:         transaction.XenditExternalID,
		Amount:             float64(product.Price),
		Description:        "Invoice PayOll " + user.ID.Hex() + "-" + fmt.Sprint(timestamp),
		InvoiceDuration:    86400,
		Customer:           customer,
		SuccessRedirectURL: req.SuccessRedirectURL,
		FailureRedirectURL: req.FailureRedirectURL,
		Currency:           "IDR",
		Items:              []xendit.InvoiceItem{item},
		Fees:               []xendit.InvoiceFee{fee},
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		return resp, errors.New("Failed to create Xendit Invoice")
	}

	return resp, nil
}

func (t *TransactionUseCase) ChangeTransactionStatus(id, status string) (Domain, error) {
	// Get Transaction
	ObjID, _ := primitive.ObjectIDFromHex(id)
	transaction, err := t.TransactionRepository.GetByID(ObjID)
	if err != nil {
		return transaction, errors.New("Transaction not found")
	}

	// Change Transaction Status
	transaction.Status = status
	transaction.Updated = primitive.NewDateTimeFromTime(time.Now())

	// Update Transaction
	transaction, err = t.TransactionRepository.Update(&transaction)

	return transaction, err
}

// ChangePaymentStatus
func (t *TransactionUseCase) UpdatePaymentInfo(id, status, payment_method, payment_channel string) (Domain, error) {
	// Get Transaction
	ObjID, _ := primitive.ObjectIDFromHex(id)
	transaction, err := t.TransactionRepository.GetByID(ObjID)
	if err != nil {
		return transaction, errors.New("Transaction not found")
	}

	// Change Transaction Status
	transaction.XenditStatus = status
	transaction.XenditPaymentMethod = payment_method
	transaction.XenditPaymentChannel = payment_channel
	transaction.Updated = primitive.NewDateTimeFromTime(time.Now())

	// Update Transaction
	transaction, err = t.TransactionRepository.Update(&transaction)

	return transaction, err
}

func (t *TransactionUseCase) GetTransactionByID(id string) (Domain, error) {
	// Get Transaction
	ObjID, _ := primitive.ObjectIDFromHex(id)
	transaction, err := t.TransactionRepository.GetByID(ObjID)
	if err != nil {
		return transaction, errors.New("Transaction not found")
	}

	return transaction, err
}

func (t *TransactionUseCase) GetTransactionsByUserID(userID string) ([]Domain, error) {
	// Get Transactions
	ObjID, _ := primitive.ObjectIDFromHex(userID)
	transactions, err := t.TransactionRepository.GetByUserID(ObjID)
	if err != nil {
		return transactions, errors.New("Transactions not found")
	}

	return transactions, err
}

// Get By Xendit Invoice ID
func (t *TransactionUseCase) GetTransactionByXenditInvoiceID(id string) (Domain, error) {
	// Get Transaction
	transaction, err := t.TransactionRepository.GetByXenditInvoiceID(id)
	if err != nil {
		return transaction, errors.New("Transaction not found")
	}

	return transaction, err
}

// Get Transaction History By User ID
func (t *TransactionUseCase) GetTransactionHistoryByUserID(userID string) ([]response.TransactionResponse, error) {
	// Get Transactions
	ObjID, _ := primitive.ObjectIDFromHex(userID)

	transactions, err := t.TransactionRepository.GetAllTransactionHistoryByUserID(ObjID)
	if err != nil {
		return transactions, err
	}

	return transactions, err
}

// Get Transaction History By ID
func (t *TransactionUseCase) GetTransactionHistoryByID(id string, user *users.Domain) (response.TransactionResponse, error) {
	// Get Transaction
	ObjID, _ := primitive.ObjectIDFromHex(id)

	transaction, err := t.TransactionRepository.GetByID(ObjID)
	if err != nil {
		return response.TransactionResponse{}, err
	}

	// Check if user is the owner of the transaction, if admin then return the transaction
	if transaction.UserID.Hex() != user.ID.Hex() && user.Role != "admin" {
		return response.TransactionResponse{}, errors.New("You are not the owner of this transaction")
	}

	// Get the response data
	transactionResp, err := t.TransactionRepository.GetTransactionHistoryByID(ObjID)

	return transactionResp, err
}

func (t *TransactionUseCase) GetAllTransaction(page, limit int64, status string) ([]response.TransactionResponse, error) {
	// Get Transactions
	transactions, err := t.TransactionRepository.GetAllTransaction(page, limit, status)
	if err != nil {
		return transactions, err
	}

	return transactions, err
}

// Count
func (t *TransactionUseCase) GetTotalTransaction() (int64, error) {
	// Get Transactions
	total, err := t.TransactionRepository.Count()
	if err != nil {
		return total, err
	}

	return total, err
}