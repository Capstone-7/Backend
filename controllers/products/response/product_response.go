package response

import (
	"capstone/businesses/products"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID			primitive.ObjectID  `bson:"_id,omitempty" json:"_id"`
	Code         string `bson:"code" json:"code"`
	Description  string `bson:"description" json:"description"`
	Nominal      string `bson:"nominal" json:"nominal"`
	Details      string `bson:"details" json:"details"`
	Price        int64	`bson:"price" json:"price"`
	Type         string `bson:"type" json:"type"`
	Category	 string `bson:"category" json:"category"`
	ActivePeriod int	`bson:"active_period" json:"active_period"`
	Status       string `bson:"status" json:"status"`
	IconUrl      string `bson:"icon_url" json:"icon_url"`
}

func FromDomain(domain *products.Domain) *Product {
	return &Product{
		ID: domain.ID,
		Code: domain.Code,
		Description: domain.Description,
		Nominal: domain.Nominal,
		Details: domain.Details,
		Price: domain.Price,
		Type: domain.Type,
		Category: domain.Category,
		ActivePeriod: domain.ActivePeriod,
		Status: domain.Status,
		IconUrl: domain.IconUrl,
	}
}

func FromDomainArray(domain []products.Domain) []Product {
	var products []Product
	for _, v := range domain {
		products = append(products, *FromDomain(&v))
	}
	return products
}