package products

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Code         string             `bson:"code" json:"code"`
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

type UseCase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	CreateProduct(domain *Domain) (Domain, error)
	UpdateProduct(domain *Domain) (Domain, error)
	DeleteProduct(id string) (Domain, error)
}

type Repository interface {
	// New
	Create(domain *Domain) (Domain, error)
	Update( new *Domain) (Domain, error)
	Delete(id primitive.ObjectID) (Domain, error)
	GetByID(id primitive.ObjectID) (Domain, error)
	GetAll() ([]Domain, error)
}
