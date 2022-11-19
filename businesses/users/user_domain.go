package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contains users usecase definition

type Domain struct {
	ID			primitive.ObjectID  `bson:"_id,omitempty" json:"_id"`
	Name		string				`bson:"name" json:"name"`
	Email		string				`bson:"email" json:"email"`
	Password	string				`bson:"password" json:"password"`
	Role		string		  		`bson:"role" json:"role"`
	Status		string				`bson:"status" json:"status"`
	Created		primitive.DateTime	`bson:"created" json:"created"`
	Updated		primitive.DateTime	`bson:"updated" json:"updated"`
	Deleted		primitive.DateTime	`bson:"deleted" json:"deleted"`
}

type UseCase interface {
	Register(domain *Domain) (Domain, error)
	Login(domain *Domain) (string, error)
	UpdateProfile(domain *Domain) (Domain, error)
	UpdatePassword(old *Domain, new *Domain) (Domain, error)
	ResetPassword(email, new_password, otp string) (Domain, error)
	RequestOTP(email, scope string) (string, error)
	VerifyEmail(email, code string) (string, error)
	GetAllUsers() ([]Domain, error)
	GetByID(id string) (Domain, error)
	UpdateByAdmin(new *Domain) (Domain, error)
	DeleteByAdmin(id string) (Domain, error)
}

type Repository interface {
	Create(domain *Domain) (Domain, error)
	Update(new *Domain) (Domain, error)
	Delete(id primitive.ObjectID) (Domain, error)
	GetByID(id primitive.ObjectID) (Domain, error)
	GetByEmail(email string) (Domain, error)
	GetAll() ([]Domain, error)
}