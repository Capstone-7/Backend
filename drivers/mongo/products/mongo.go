package products

import (
	"capstone/businesses/products"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type productRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database) products.Repository{
	return &productRepository{
		collection: db.Collection("products"),
	}
}

func (r *productRepository) Create(domain *products.Domain) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Generate ID
	domain.Id = primitive.NewObjectID()

	// Insert data
	_, err := r.collection.InsertOne(ctx, domain)
	if err != nil {
		return products.Domain{}, err
	}

	return *domain, nil
}

func (r productRepository) Update(new *products.Domain) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Update data
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": new.Id}, bson.M{"$set": new})
	if err != nil {
		return products.Domain{}, err
	}

	// Get new data
	var product products.Domain
	err = r.collection.FindOne(ctx, bson.M{"_id": new.Id}).Decode(&product)

	return product, nil
}

func (r productRepository) Delete(id primitive.ObjectID) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var product products.Domain
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return products.Domain{}, err
	}

	// Shadow delete
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted": primitive.NewDateTimeFromTime(time.Now())}})
	if err != nil {
		return products.Domain{}, err
	}

	return product, nil
}

func (r *productRepository) GetByID(id primitive.ObjectID) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var product products.Domain
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return products.Domain{}, err
	}

	return product, nil
}

func (r *productRepository) GetAll() ([]products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var product []Product

	// Get all products
	cursor, err := r.collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return ToDomainArray(&product), err
	}

	// Decode cursor to products
	err = cursor.All(ctx, &product)
	if err != nil {
		return ToDomainArray(&product), err
	}

	return ToDomainArray(&product), err
}
