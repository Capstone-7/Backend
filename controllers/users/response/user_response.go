package response

import (
	"capstone/businesses/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contains responses definition

type User struct {
	ID			primitive.ObjectID  `bson:"_id,omitempty" json:"_id"`
	Name		string				`bson:"name" json:"name"`
	Email		string				`bson:"email" json:"email"`
	Role		string		  		`bson:"role" json:"role"`
	Status		string				`bson:"status" json:"status"`
	Created		primitive.DateTime	`bson:"created" json:"created"`
	Updated		primitive.DateTime	`bson:"updated" json:"updated"`
	Deleted		primitive.DateTime	`bson:"deleted" json:"deleted"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
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