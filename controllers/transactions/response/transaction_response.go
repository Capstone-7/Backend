package response

type ReviewTransaction struct {
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	ProductPrice       int64  `json:"product_price"`
	AdminFee           int64  `json:"admin_fee"`
	TotalPrice         int64  `json:"total_price"`
}