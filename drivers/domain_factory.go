package drivers

import (
	otpDomain "capstone/businesses/otp"
	userDomain "capstone/businesses/users"
	otpDB "capstone/drivers/mongo/otp"
	userDB "capstone/drivers/mongo/users"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) userDomain.Repository {
	return userDB.NewUserRepository(db)
}

func NewOTPRepository(db *mongo.Database) otpDomain.Repository {
	return otpDB.NewOTPRepository(db)
}