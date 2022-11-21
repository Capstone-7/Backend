package products

import (
	"capstone/businesses/products"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	fmt.Println("domain", FromDomain(domain))

	// Insert data
	res, err := r.collection.InsertOne(ctx, FromDomain(domain))
	if err != nil {
		return products.Domain{}, err
	}

	// Get inserted data
	var product products.Domain
	err = r.collection.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&product)

	return product, nil
}

func (r *productRepository) Update(new *products.Domain) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Update data
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": new.ID}, bson.M{"$set": new})
	if err != nil {
		return products.Domain{}, err
	}

	// Get new data
	var product products.Domain
	err = r.collection.FindOne(ctx, bson.M{"_id": new.ID}).Decode(&product)

	return product, nil
}

func (r *productRepository) Delete(id primitive.ObjectID) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Shadow delete
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"deleted": primitive.NewDateTimeFromTime(time.Now())}})
	if err != nil {
		return products.Domain{}, err
	}

	// Get deleted data
	var product products.Domain
	err = r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)

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

func (r *productRepository) GetAll(only_not_deleted bool) ([]products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var product []Product
	// Get all products
	if only_not_deleted {
		cursor, err := r.collection.Find(ctx, bson.M{"deleted": primitive.NewDateTimeFromTime(time.Time{})})
		if err != nil {
			return []products.Domain{}, err
		}

		// Decode cursor to products
		err = cursor.All(ctx, &product)
		if err != nil {
			return []products.Domain{}, err
		}

	} else {
		cursor, err := r.collection.Find(ctx, bson.M{})
		if err != nil {
			return []products.Domain{}, err
		}

		// Decode cursor to products
		err = cursor.All(ctx, &product)
		if err != nil {
			return []products.Domain{}, err
		}

	}

	return ToDomainArray(&product), nil
}

// Get One Product by Code
func (r *productRepository) GetOneByCode(code string, only_not_deleted bool) (products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var product Product

	// Get one product not deleted
	if only_not_deleted {
		err := r.collection.FindOne(ctx, bson.M{"code": code, "deleted": primitive.NewDateTimeFromTime(time.Time{})}).Decode(&product)
		if err != nil {
			return product.ToDomain(), err
		}
	} else {
		err := r.collection.FindOne(ctx, bson.M{"code": code}).Decode(&product)
		if err != nil {
			return product.ToDomain(), err
		}
	}

	return product.ToDomain(), nil
}

// Get Many Product by Type
func (r *productRepository) GetManyByType(product_type string, only_not_deleted bool) ([]products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var out_products []Product

	// Get many product not deleted
	if only_not_deleted {
		cursor, err := r.collection.Find(ctx, bson.M{"type": product_type, "deleted": primitive.NewDateTimeFromTime(time.Time{})})
		if err != nil {
			return []products.Domain{}, err
		}

		// Decode cursor to products
		err = cursor.All(ctx, &out_products)
		if err != nil {
			return []products.Domain{}, err
		}

	} else {
		cursor, err := r.collection.Find(ctx, bson.M{"type": product_type})
		if err != nil {
			return []products.Domain{}, err
		}

		// Decode cursor to products
		err = cursor.All(ctx, &out_products)
		if err != nil {
			return []products.Domain{}, err
		}

	}

	return ToDomainArray(&out_products), nil
}

// Get Many Product by Category
func (r *productRepository) GetManyByCategory(category string, only_not_deleted bool) ([]products.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var out_products []Product

	// Get many product not deleted
	if only_not_deleted {
		cursor, err := r.collection.Find(ctx, bson.M{"category": category, "deleted": primitive.NewDateTimeFromTime(time.Time{})})
		if err != nil {
			return []products.Domain{}, err
		}
		
		// Decode cursor to products
		err = cursor.All(ctx, &out_products)
		if err != nil {
			return []products.Domain{}, err
		}

	} else {
		cursor, err := r.collection.Find(ctx, bson.M{"category": category})
		if err != nil {
			return []products.Domain{}, err
		}

		// Decode cursor to products
		err = cursor.All(ctx, &out_products)
		if err != nil {
			return []products.Domain{}, err
		}

	}

	return ToDomainArray(&out_products), nil
}

// Get Categories Only
func (r *productRepository) GetCategories() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var out_categories []string

	// Get categories only
	cursor, err := r.collection.Distinct(ctx, "category", bson.M{"deleted": primitive.NewDateTimeFromTime(time.Time{})})
	if err != nil {
		return []string{}, err
	}

	// loop cursor to categories
	for _, category := range cursor {
		out_categories = append(out_categories, category.(string))
	}

	return out_categories, nil
}

// Get Categories By Type
func (r *productRepository) GetCategoriesByType(product_type string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var out_categories []string

	// Get categories only
	cursor, err := r.collection.Distinct(ctx, "category", bson.M{"type": product_type, "deleted": primitive.NewDateTimeFromTime(time.Time{})})
	if err != nil {
		return []string{}, err
	}

	// loop cursor to categories
	for _, category := range cursor {
		out_categories = append(out_categories, category.(string))
	}

	return out_categories, nil
}

func (r *productRepository) CountProducts() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Count all products not deleted
	count, err := r.collection.CountDocuments(ctx, bson.M{"deleted": primitive.NewDateTimeFromTime(time.Time{})})
	if err != nil {
		return 0, err
	}

	return count, nil
}