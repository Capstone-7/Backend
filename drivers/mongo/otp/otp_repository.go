package otp

import (
	"capstone/businesses/otp"
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)	

type OTPRepository struct {
	collection *mongo.Collection
}

func NewOTPRepository(db *mongo.Database) otp.Repository {
	return &OTPRepository{
		collection: db.Collection("otp"),
	}
}

// GenerateOTP
func (r *OTPRepository) GenerateOTP(id, scope string) (*otp.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Generate OTP
	ObjID, _ := primitive.ObjectIDFromHex(id)

	// Generate random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999-100000) + 100000
	// Generate expire time

	domain := otp.Domain{
		UserID: ObjID,
		Code: fmt.Sprint(code), // Convert int to string
		Expire: primitive.NewDateTimeFromTime(time.Now().Add(10 * time.Minute)),
		Scope: scope,
		Status: "active",
		Created: primitive.NewDateTimeFromTime(time.Now()),
		Updated: primitive.NewDateTimeFromTime(time.Now()),
		Deleted: primitive.NewDateTimeFromTime(time.Time{}),
	}
	
	// Insert data
	_, err := r.collection.InsertOne(ctx, domain)
	if err != nil {
		return nil, err
	}

	return &domain, nil
}

// VerifyOTP
func (r *OTPRepository) VerifyOTP(req *otp.Domain) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Find OTP Not Expired
	var domain otp.Domain
	err := r.collection.FindOne(ctx, bson.M{
		"user_id": req.UserID,
		"code": req.Code,
		"expire": bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now())},
		"scope": req.Scope,
		"status": "active",
	}).Decode(&domain)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Consume OTP
func (r *OTPRepository) ConsumeOTP(req *otp.Domain) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(ctx, bson.M{
		"user_id": req.UserID,
		"code": req.Code,
		"expire": bson.M{"$gt": primitive.NewDateTimeFromTime(time.Now())},
		"scope": req.Scope,
		"status": "active",
	}, bson.M{
		"$set": bson.M{
			"status": "consumed",
			"updated": primitive.NewDateTimeFromTime(time.Now()),
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

// Get last OTP by user id
func (r *OTPRepository) GetLastByUserID(id string) (*otp.Domain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Sort by _id
	ObjID, _ := primitive.ObjectIDFromHex(id)
	cursor, err := r.collection.Find(ctx, bson.D{
		{Key: "user_id", Value: ObjID},
	}, options.Find().SetSort(bson.D{{Key: "_id", Value: -1}}))

	if err != nil {
		return nil, err
	}

	var domain otp.Domain
	if cursor.Next(ctx) {
		err := cursor.Decode(&domain)
		if err != nil {
			return nil, err
		}
	}
	
	return &domain, nil
}