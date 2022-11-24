package drivers

import (
	otpDomain "capstone/businesses/otp"
	productDomain "capstone/businesses/products"
	userDomain "capstone/businesses/users"
	otpDB "capstone/drivers/mongo/otp"
	productDB "capstone/drivers/mongo/products"
	userDB "capstone/drivers/mongo/users"

	transactionDomain "capstone/businesses/transactions"
	transactionDB "capstone/drivers/mongo/transactions"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewProductRepository(db *mongo.Database) productDomain.Repository {
	return productDB.NewMongoRepository(db)
}

func NewUserRepository(db *mongo.Database) userDomain.Repository {
	return userDB.NewUserRepository(db)
}

func NewOTPRepository(db *mongo.Database) otpDomain.Repository {
	return otpDB.NewOTPRepository(db)
}

func NewTransactionRepository(db *mongo.Database) transactionDomain.Repository {
	return transactionDB.NewTransactionRepository(db)
}