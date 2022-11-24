package products

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Code         string             `bson:"code" json:"code"`
	Description  string             `bson:"description" json:"description"`
	Nominal      string             `bson:"nominal" json:"nominal"`
	Details      string             `bson:"details" json:"details"`
	Price        int64              `bson:"price" json:"price"`
	Type         string             `bson:"type" json:"type"`
	Category     string             `bson:"category" json:"category"`
	ActivePeriod int             	`bson:"active_period" json:"active_period"`
	Status       string             `bson:"status" json:"status"`
	IconUrl      string             `bson:"icon_url" json:"icon_url"`
	Created      primitive.DateTime `bson:"created" json:"created"`
	Updated      primitive.DateTime `bson:"updated" json:"updated"`
	Deleted      primitive.DateTime `bson:"deleted" json:"deleted"`
}

type UseCase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	CreateProduct(domain *Domain) (Domain, error)
	UpdateProduct(domain *Domain) (Domain, error)
	DeleteProduct(id string) (Domain, error)

	GetCategories() ([]string, error)
	GetCategoriesByProductType(productType string) ([]string, error)
	GetProductsByCategory(category string) ([]Domain, error)
	GetProductsByProductType(productType string) ([]Domain, error)
	GetTotalProducts() (int64, error)
	GetProductByCode(code string) (Domain, error)
}

type Repository interface {
	// New
	Create(domain *Domain) (Domain, error)
	Update( new *Domain) (Domain, error)
	Delete(id primitive.ObjectID) (Domain, error)
	GetByID(id primitive.ObjectID) (Domain, error)
	GetOneByCode(code string, only_not_deleted bool) (Domain, error)
	GetManyByType(code string, only_not_deleted bool) ([]Domain, error)
	GetManyByCategory(code string, only_not_deleted bool) ([]Domain, error)
	GetCategories() ([]string, error)
	GetCategoriesByType(product_type string) ([]string, error)
	GetAll(only_not_deleted bool) ([]Domain, error)
	CountProducts() (int64, error)
}
