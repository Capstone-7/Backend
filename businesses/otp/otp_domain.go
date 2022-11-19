package otp

import "go.mongodb.org/mongo-driver/bson/primitive"

type Domain struct {
	UserID  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Code    string             `json:"code" bson:"code"`
	Expire  primitive.DateTime `json:"expire" bson:"expire"`
	Scope   string             `json:"scope" bson:"scope"`
	Status  string             `json:"status" bson:"status"`
	Created primitive.DateTime `bson:"created" json:"created"`
	Updated primitive.DateTime `bson:"updated" json:"updated"`
	Deleted primitive.DateTime `bson:"deleted" json:"deleted"`
}

type Repository interface {
	GenerateOTP(id, scope string) (*Domain, error)
	VerifyOTP(domain *Domain) (bool, error)
	ConsumeOTP(domain *Domain) (bool, error)
	GetLastByUserID(id string) (*Domain, error)
}