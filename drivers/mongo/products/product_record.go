package products

import (
	"capstone/businesses/products"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Code         string             `bson:"code" json:"code"`
	Description  string             `bson:"description" json:"description"`
	Nominal      string             `bson:"nominal" json:"nominal"`
	Details      string             `bson:"details" json:"details"`
	Price        int64             	`bson:"price" json:"price"`
	Type         string             `bson:"type" json:"type"`
	Category     string             `bson:"category" json:"category"`
	ActivePeriod int	            `bson:"active_period" json:"active_period"`
	Status       string             `bson:"status" json:"status"`
	IconUrl      string             `bson:"icon_url" json:"icon_url"`
	Created      primitive.DateTime `bson:"created" json:"created"`
	Updated      primitive.DateTime `bson:"updated" json:"updated"`
	Deleted      primitive.DateTime `bson:"deleted" json:"deleted"`
}

func FromDomain(domain *products.Domain) *Product {
	return &Product{
		ID:           domain.ID,
		Code:         domain.Code,
		Description:  domain.Description,
		Nominal:      domain.Nominal,
		Details:      domain.Details,
		Price:        domain.Price,
		Type:         domain.Type,
		Category:     domain.Category,
		ActivePeriod: domain.ActivePeriod,
		Status:       domain.Status,
		IconUrl:      domain.IconUrl,
		Created:      domain.Created,
		Updated:      domain.Updated,
		Deleted:      domain.Deleted,
	}
}

func FromDomainArray(domain []products.Domain) []Product {
	var res []Product
	for _, value := range domain {
		res = append(res, *FromDomain(&value))
	}
	return res
}

func (p *Product) ToDomain() products.Domain {
	return products.Domain{
		ID:           p.ID,
		Code:         p.Code,
		Description:  p.Description,
		Nominal:      p.Nominal,
		Details:      p.Details,
		Price:        p.Price,
		Type:         p.Type,
		Category:     p.Category,
		ActivePeriod: p.ActivePeriod,
		Status:       p.Status,
		IconUrl:      p.IconUrl,
		Created:    p.Created,
		Updated:    p.Updated,
		Deleted:    p.Deleted,
	}
}

func ToDomainArray(u *[]Product) []products.Domain {
	var res []products.Domain
	for _, value := range *u {
		res = append(res, value.ToDomain())
	}
	return res
}
