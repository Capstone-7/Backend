package drivers

import (
	productDomain "capstone/businesses/products"
	"go.mongodb.org/mongo-driver/mongo"

	productDB "capstone/drivers/mongo/products"
)

func NewProductRepository(db *mongo.Database) productDomain.Repository {
	return productDB.NewMongoRepository(db)
}