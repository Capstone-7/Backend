package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReviewTransaction struct {
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	ProductPrice       int64  `json:"product_price"`
	AdminFee           int64  `json:"admin_fee"`
	TotalPrice         int64  `json:"total_price"`
}

type TransactionResponse struct {
	ID                   string  `json:"id" bson:"_id,omitempty"`
	UserEmail            string  `json:"user_email" bson:"user_email"`
	ProductCode          string  `json:"product_code" bson:"product_code"`
	ProductDescription   string  `json:"product_description" bson:"product_description"`
	ProductPrice         int64   `json:"product_price" bson:"product_price"`
	AdminFee             int64   `json:"admin_fee" bson:"admin_fee"`
	TotalPrice           int64   `json:"total_price" bson:"total_price"`
	XenditPaymentURL     string  `json:"xendit_payment_url" bson:"xendit_payment_url"`
	XenditStatus         string  `json:"xendit_status" bson:"xendit_status"`
	XenditPaymentMethod  string  `json:"xendit_payment_method" bson:"xendit_payment_method"`
	XenditPaymentChannel string  `json:"xendit_payment_channel" bson:"xendit_payment_channel"`
	Status               string  `json:"status" bson:"status"`
	Created              primitive.DateTime `json:"created" bson:"created"`
	Updated              primitive.DateTime   `json:"updated" bson:"updated"`
}