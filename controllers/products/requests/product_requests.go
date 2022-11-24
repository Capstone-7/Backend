package requests

import (
	"capstone/businesses/products"
	"capstone/helpers"
	"errors"
	"strings"

	"github.com/fatih/structs"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	Code         string `bson:"code" json:"code" validate:"required,min=3"`
	Description  string `bson:"description" json:"description" validate:"required,min=1"`
	Nominal      string `bson:"nominal" json:"nominal" validate:"required"`
	Details      string `bson:"details" json:"details" validate:"required"`
	Price        int64	`bson:"price" json:"price" validate:"gte=0"`
	Type         string `bson:"type" json:"type" validate:"required"`
	Category     string `bson:"category" json:"category" validate:"required"`
	ActivePeriod int 	`bson:"active_period" json:"active_period" validate:"gte=0"`
	Status       string `bson:"status" json:"status" validate:"required"`
	IconUrl      string `bson:"icon_url" json:"icon_url" validate:"required"`
}

func (p *Product) ToDomain() *products.Domain {
	return &products.Domain{
		Code: p.Code,
		Description: p.Description,
		Nominal: p.Nominal,
		Details: p.Details,
		Price: p.Price,
		Type: p.Type,
		Category: p.Category,
		ActivePeriod: p.ActivePeriod,
		Status: p.Status,
		IconUrl: p.IconUrl,
	}
}

func (p *Product) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(p); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(p)
			out := make([]helpers.ValidationErrors, len(ve))

			for i, e := range ve {
				out[i] = helpers.ValidationErrors{
					Field:   e.Field(),
					Message: helpers.MsgForTag(e.Tag()),
				}

				out[i].Message = strings.Replace(out[i].Message, "[PARAM]", e.Param(), 1)

				// Get field tag
				for _, f := range fields {
					if f.Name() == e.Field() {
						out[i].Field = f.Tag("json")
						break
					}
				}
			}
			return out
		}
	}

	return nil
}