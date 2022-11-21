package products

import (
	"capstone/businesses/products"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Code         string             `bson:"username" json:"username"`
	Description  string             `bson:"description" json:"description"`
	Nominal      string             `bson:"nominal" json:"nominal"`
	Details      string             `bson:"details" json:"details"`
	Price        int             `bson:"price" json:"price"`
	Type         string             `bson:"type" json:"type"`
	ActivePeriod string             `bson:"active_period" json:"active_period"`
	Status       string             `bson:"status" json:"status"`
	IconUrl      string             `bson:"icon_url" json:"icon_url"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at" json:"updated_at"`
	DeletedAt    primitive.DateTime `bson:"deleted_at" json:"deleted_at"`
}

func FromDomain(domain *products.Domain) *Product {
	return &Product{
		Id:           domain.Id,
		Code:         domain.Code,
		Description:  domain.Description,
		Nominal:      domain.Nominal,
		Details:      domain.Details,
		Price:        domain.Price,
		Type:         domain.Type,
		ActivePeriod: domain.ActivePeriod,
		Status:       domain.Status,
		IconUrl:      domain.IconUrl,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
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
		Id:           p.Id,
		Code:         p.Code,
		Description:  p.Description,
		Nominal:      p.Nominal,
		Details:      p.Details,
		Price:        p.Price,
		Type:         p.Type,
		ActivePeriod: p.ActivePeriod,
		Status:       p.Status,
		IconUrl:      p.IconUrl,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		DeletedAt:    p.DeletedAt,
	}
}

func ToDomainArray(u *[]Product) []products.Domain {
	var res []products.Domain
	for _, value := range *u {
		res = append(res, value.ToDomain())
	}
	return res
}
