package requests

import "capstone/businesses/products"

type Product struct {
	Code         string `bson:"code" json:"code"`
	Description  string `bson:"description" json:"description"`
	Nominal      string `bson:"nominal" json:"nominal"`
	Details      string `bson:"details" json:"details"`
	Price        int	`bson:"price" json:"price"`
	Type         string `bson:"type" json:"type"`
	ActivePeriod string `bson:"active_period" json:"active_period"`
	Status       string `bson:"status" json:"status"`
	IconUrl      string `bson:"icon_url" json:"icon_url"`
}

func (p *Product) ToDomain() *products.Domain {
	return &products.Domain{
		Code: p.Code,
		Description: p.Description,
		Nominal: p.Nominal,
		Details: p.Details,
		Price: p.Price,
		Type: p.Type,
		ActivePeriod: p.ActivePeriod,
		Status: p.Status,
		IconUrl: p.IconUrl,
	}
}
