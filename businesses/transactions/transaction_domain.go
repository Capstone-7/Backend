package transactions

import (
	"capstone/businesses/products"
	"capstone/businesses/users"
	"capstone/controllers/transactions/requests"
	"capstone/controllers/transactions/response"
	"time"

	"github.com/xendit/xendit-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	CustomerID string `bson:"customer_id" json:"customer_id"`
	ProductPrice int64 `bson:"product_price" json:"product_price"`
	AdminFee int64 `bson:"admin_fee" json:"admin_fee"`
	TotalPrice int64 `bson:"total_price" json:"total_price"`
	XenditInvoiceID string `bson:"xendit_invoice_id" json:"xendit_invoice_id"`
	XenditPaymentURL string `bson:"xendit_payment_url" json:"xendit_payment_url"`
	XenditStatus string `bson:"xendit_status" json:"xendit_status"`
	XenditExternalID string `bson:"xendit_external_id" json:"xendit_external_id"`
	XenditPaymentMethod string `bson:"xendit_payment_method" json:"xendit_payment_method"`
	XenditPaymentChannel string `bson:"xendit_payment_channel" json:"xendit_payment_channel"`
	Status string `bson:"status" json:"status"`
	Created primitive.DateTime `bson:"created" json:"created"`
	Updated primitive.DateTime `bson:"updated" json:"updated"`
	Deleted primitive.DateTime `bson:"deleted" json:"deleted"`
}

type UseCase interface {
	InitTransaction(product *products.Domain, user *users.Domain, customer_id string) Domain
	CreateXenditInvoice(product *products.Domain, user *users.Domain, transcation *Domain, req *requests.SubmitTransactionRequest) (*xendit.Invoice, error)
	CreateTransaction(product *products.Domain, user *users.Domain, req *requests.SubmitTransactionRequest) (Domain, error)
	ChangeTransactionStatus(id, status string) (Domain, error)
	UpdatePaymentInfo(id, status, payment_method, payment_channel string) (Domain, error)
	GetTransactionByID(id string) (Domain, error)
	GetTransactionsByUserID(userID string) ([]Domain, error)
	GetTransactionByXenditInvoiceID(xenditInvoiceID string) (Domain, error)

	GetTransactionHistoryByUserID(userID string) ([]response.TransactionResponse, error)
	GetTransactionHistoryByID(id string, user *users.Domain) (response.TransactionResponse, error)
	GetAllTransaction(page, limit int64, status string) ([]response.TransactionResponse, error)
	GetTotalTransaction() (int64, error)
	GetTopProductsByCategory() (map[string]int, error)
	GetIncomePerDay() (map[time.Time]map[string]int64, error)
}

type Repository interface {
	Create(domain *Domain) (Domain, error)
	Update(domain *Domain) (Domain, error)
	GetAll(filter map[string]interface{}) ([]Domain, error)
	GetByID(id primitive.ObjectID) (Domain, error)
	GetByUserID(userID primitive.ObjectID) ([]Domain, error)
	GetByXenditInvoiceID(xenditInvoiceID string) (Domain, error)
	
	GetTransactionHistoryByID(id primitive.ObjectID) (response.TransactionResponse, error)
	GetAllTransactionHistoryByUserID(userID primitive.ObjectID) ([]response.TransactionResponse, error)
	GetAllTransaction(page, limit int64, status string) ([]response.TransactionResponse, error)
	Count() (int64, error)
	GetTopProductsByCategory() (map[string]int, error)
	GetIncomeByTypeBetweenDate(startDate, endDate primitive.DateTime) (map[string]int64, error)
}