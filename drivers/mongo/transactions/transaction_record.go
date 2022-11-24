package transactions

import (
	"capstone/businesses/transactions"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
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

func FromDomain(domain *transactions.Domain) *Transaction {
	return &Transaction{
		ID: domain.ID,
		UserID: domain.UserID,
		ProductID: domain.ProductID,
		CustomerID: domain.CustomerID,
		ProductPrice: domain.ProductPrice,
		AdminFee: domain.AdminFee,
		TotalPrice: domain.TotalPrice,
		XenditInvoiceID: domain.XenditInvoiceID,
		XenditPaymentURL: domain.XenditPaymentURL,
		XenditStatus: domain.XenditStatus,
		XenditExternalID: domain.XenditExternalID,
		XenditPaymentMethod: domain.XenditPaymentMethod,
		XenditPaymentChannel: domain.XenditPaymentChannel,
		Status: domain.Status,
		Created: domain.Created,
		Updated: domain.Updated,
		Deleted: domain.Deleted,
	}
}

func FromDomainArray(domain []transactions.Domain) []Transaction {
	var res []Transaction
	for _, val := range domain {
		res = append(res, *FromDomain(&val))
	}
	return res
}

func (rec *Transaction) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID: rec.ID,
		UserID: rec.UserID,
		ProductID: rec.ProductID,
		CustomerID: rec.CustomerID,
		ProductPrice: rec.ProductPrice,
		AdminFee: rec.AdminFee,
		TotalPrice: rec.TotalPrice,
		XenditInvoiceID: rec.XenditInvoiceID,
		XenditPaymentURL: rec.XenditPaymentURL,
		XenditStatus: rec.XenditStatus,
		XenditExternalID: rec.XenditExternalID,
		XenditPaymentMethod: rec.XenditPaymentMethod,
		XenditPaymentChannel: rec.XenditPaymentChannel,
		Status: rec.Status,
		Created: rec.Created,
		Updated: rec.Updated,
		Deleted: rec.Deleted,
	}
}