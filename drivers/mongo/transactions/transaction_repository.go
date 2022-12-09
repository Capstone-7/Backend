package transactions

import (
	"capstone/businesses/transactions"
	"capstone/controllers/transactions/response"
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

// Get Transaction History By ID
func (t *TransactionRepository) GetTransactionHistoryByID(id primitive.ObjectID) (response.TransactionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var transaction response.TransactionResponse
	// Get transaction with aggregation
	cursor, err := t.collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "user_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "user"},
				},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "products"},
					{Key: "localField", Value: "product_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "product"},
				},
			},
		},
		bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "user_email",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$user.email",
									0,
								},
							},
						},
					},
					{Key: "product_code",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.code",
									0,
								},
							},
						},
					},
					{Key: "product_description",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.description",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "user_id", Value: 0},
					{Key: "product_id", Value: 0},
					{Key: "deleted", Value: 0},
					{Key: "user", Value: 0},
					{Key: "product", Value: 0},
				},
			},
		},
	})

	if err != nil {
		return response.TransactionResponse{}, err
	}

	// Decode data only one
	if cursor.Next(ctx) {
		err = cursor.Decode(&transaction)
	}

	return transaction, nil
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

// Get Transaction History By User ID
func (t *TransactionRepository) GetAllTransactionHistoryByUserID(userID primitive.ObjectID) ([]response.TransactionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var transaction_array []response.TransactionResponse
	cursor, err := t.collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userID}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "user_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "user"},
				},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "products"},
					{Key: "localField", Value: "product_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "product"},
				},
			},
		},
		bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "user_email",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$user.email",
									0,
								},
							},
						},
					},
					{Key: "product_code",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.code",
									0,
								},
							},
						},
					},
					{Key: "product_description",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.description",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "user_id", Value: 0},
					{Key: "product_id", Value: 0},
					{Key: "deleted", Value: 0},
					{Key: "user", Value: 0},
					{Key: "product", Value: 0},
				},
			},
		},
	})

	if err != nil {
		return []response.TransactionResponse{}, err
	}

	err = cursor.All(ctx, &transaction_array)
	return transaction_array, err
}

// GetAllTransaction(page, limit int64, status string)
func (t *TransactionRepository) GetAllTransaction(page, limit int64, status string) ([]response.TransactionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	search := bson.D{}
	if status != "" {
		search = bson.D{{Key: "$match", Value: bson.D{{Key: "status", Value: status}}}}
	}else{
		// search all
		search = bson.D{{Key: "$match", Value: bson.D{{Key: "status", Value: bson.D{{Key: "$ne", Value: "DELETED"}}}}}}
	}

	var transaction_array []response.TransactionResponse
	cursor, err := t.collection.Aggregate(ctx, bson.A{
		search,
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "user_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "user"},
				},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "products"},
					{Key: "localField", Value: "product_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "product"},
				},
			},
		},
		bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "user_email",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$user.email",
									0,
								},
							},
						},
					},
					{Key: "product_code",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.code",
									0,
								},
							},
						},
					},
					{Key: "product_description",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.description",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "user_id", Value: 0},
					{Key: "product_id", Value: 0},
					{Key: "deleted", Value: 0},
					{Key: "user", Value: 0},
					{Key: "product", Value: 0},
				},
			},
		},
		bson.D{
			{Key: "$skip", Value: (page - 1) * limit},
		},
		bson.D{
			{Key: "$limit", Value: limit},
		},
	})

	if err != nil {
		return []response.TransactionResponse{}, err
	}

	err = cursor.All(ctx, &transaction_array)
	return transaction_array, err
}

// Count all transaction
func (t *TransactionRepository) Count() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	count, err := t.collection.CountDocuments(ctx, bson.D{})
	return count, err
}

// Get Top Products By Category
func (t *TransactionRepository) GetTopProductsByCategory() (map[string]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	top_product_array := map[string]int{}
	cursor, err := t.collection.Aggregate(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "status", Value: "SUCCESS"}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "products"},
					{Key: "localField", Value: "product_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "product"},
				},
			},
		},
		bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "product_category",
						Value: bson.D{
							{Key: "$arrayElemAt",
								Value: bson.A{
									"$product.category",
									0,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "$product_category"},
					{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
				},
			},
		},
	})

	if err != nil {
		return map[string]int{}, err
	}

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return map[string]int{}, err
		}

		top_product_array[result["_id"].(string)] = int(result["count"].(int32))
	}
	return top_product_array, err
}