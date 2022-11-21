package users

import (
	"capstone/businesses/users"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) users.Repository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// Create
func (r *UserRepository) Create(domain *users.Domain) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Insert data
	res, err := r.collection.InsertOne(ctx, FromDomain(domain))
	if err != nil {
		return users.Domain{}, err
	}

	// Get inserted data
	var user users.Domain
	err = r.collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&user)

	return user, nil
}

// GetAll
func (r *UserRepository) GetAll() ([]users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user_array []users.Domain
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return []users.Domain{}, err
	}

	err = cursor.All(ctx, &user_array)
	if err != nil {
		return []users.Domain{}, err
	}

	return user_array, nil
}

// GetByEmail
func (r *UserRepository) GetByEmail(email string) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user users.Domain
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return users.Domain{}, err
	}

	return user, nil
}

// GetByID
func (r *UserRepository) GetByID(id primitive.ObjectID) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user users.Domain
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return users.Domain{}, err
	}

	return user, nil
}

// Update
func (r *UserRepository) Update(new *users.Domain) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Update data
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": new.ID}, bson.M{"$set": new})
	if err != nil {
		return users.Domain{}, err
	}

	// Get new data
	var user users.Domain
	err = r.collection.FindOne(ctx, bson.M{"_id": new.ID}).Decode(&user)

	return user, nil
}

// Delete
func (r *UserRepository) Delete(id primitive.ObjectID) (users.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var user users.Domain
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return users.Domain{}, err
	}

	// Shadow delete
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted": primitive.NewDateTimeFromTime(time.Now())}})
	if err != nil {
		return users.Domain{}, err
	}

	return user, nil
}