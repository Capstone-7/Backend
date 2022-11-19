package otp

import (
	"capstone/businesses/otp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func FromDomain(domain *otp.Domain) *Domain {
	return &Domain{
		UserID:  domain.UserID,
		Code:    domain.Code,
		Expire:  domain.Expire,
		Scope:   domain.Scope,
		Status:  domain.Status,
		Created: domain.Created,
		Updated: domain.Updated,
		Deleted: domain.Deleted,
	}
}

func FromDomainArray(domain []otp.Domain) []Domain {
	var otps []Domain
	for _, v := range domain {
		otps = append(otps, *FromDomain(&v))
	}
	return otps
}

func (u *Domain) ToDomain() *otp.Domain {
	return &otp.Domain{
		UserID:  u.UserID,
		Code:    u.Code,
		Expire:  u.Expire,
		Scope:   u.Scope,
		Status:  u.Status,
		Created: u.Created,
		Updated: u.Updated,
		Deleted: u.Deleted,
	}
}