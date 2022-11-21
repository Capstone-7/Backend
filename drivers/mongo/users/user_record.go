package users

import (
	"capstone/businesses/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"`
	Status   string             `bson:"status" json:"status"`
	Created  primitive.DateTime `bson:"created" json:"created"`
	Updated  primitive.DateTime `bson:"updated" json:"updated"`
	Deleted  primitive.DateTime `bson:"deleted" json:"deleted"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Role:     domain.Role,
		Status:   domain.Status,
		Created:  domain.Created,
		Updated:  domain.Updated,
		Deleted:  domain.Deleted,
	}
}

func FromDomainArray(domain []users.Domain) []User {
	var users []User
	for _, v := range domain {
		users = append(users, *FromDomain(&v))
	}
	return users
}

func (u *User) ToDomain() *users.Domain {
	return &users.Domain{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
		Status:   u.Status,
		Created:  u.Created,
		Updated:  u.Updated,
		Deleted:  u.Deleted,
	}
}

func ToDomainArray(user []User) []users.Domain {
	var users []users.Domain
	for _, v := range user {
		users = append(users, *v.ToDomain())
	}
	return users
}