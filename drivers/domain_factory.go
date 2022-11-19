package drivers

import (
	userDomain "capstone/businesses/users"
	userDB "capstone/drivers/mongo/users"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) userDomain.Repository {
	return userDB.NewUserRepository(db)
}