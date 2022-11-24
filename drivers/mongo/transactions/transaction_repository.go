package transactions

import (
	"capstone/businesses/transactions"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) transactions.Repository {
	return &TransactionRepository{
		collection: db.Collection("transactions"),
	}
}

// Create
func (t *TransactionRepository) Create(domain *transactions.Domain) (transactions.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Insert data
	res, err := t.collection.InsertOne(ctx, FromDomain(domain))
	if err != nil {
		return transactions.Domain{}, err
	}

	// Get inserted data
	var transaction transactions.Domain
	err = t.collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&transaction)

	return transaction, err
}

// Update
func (t *TransactionRepository) Update(domain *transactions.Domain) (transactions.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Update data
	_, err := t.collection.UpdateOne(ctx, bson.M{"_id": domain.ID}, bson.M{"$set": FromDomain(domain)})
	if err != nil {
		return transactions.Domain{}, err
	}

	// Get updated data
	var transaction transactions.Domain
	err = t.collection.FindOne(ctx, bson.M{"_id": domain.ID}).Decode(&transaction)

	return transaction, err
}

// Get By ID
func (t *TransactionRepository) GetByID(id primitive.ObjectID) (transactions.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var transaction transactions.Domain
	err := t.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)

	return transaction, err
}

// Get By User ID
func (t *TransactionRepository) GetByUserID(userID primitive.ObjectID) ([]transactions.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var transaction_array []transactions.Domain
	cursor, err := t.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return []transactions.Domain{}, err
	}

	err = cursor.All(ctx, &transaction_array)
	if err != nil {
		return []transactions.Domain{}, err
	}

	return transaction_array, err
}

// Get All
func (t *TransactionRepository) GetAll(filter map[string]interface{}) ([]transactions.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var transaction_array []transactions.Domain
	cursor, err := t.collection.Find(ctx, filter)
	if err != nil {
		return []transactions.Domain{}, err
	}

	err = cursor.All(ctx, &transaction_array)
	if err != nil {
		return []transactions.Domain{}, err
	}

	return transaction_array, err
}

// Get By Xendit Invoice ID
func (t *TransactionRepository) GetByXenditInvoiceID(xenditInvoiceID string) (transactions.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var transaction transactions.Domain
	err := t.collection.FindOne(ctx, bson.M{"xendit_invoice_id": xenditInvoiceID}).Decode(&transaction)

	return transaction, err
}