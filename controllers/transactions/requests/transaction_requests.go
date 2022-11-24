package requests

import (
	"capstone/helpers"
	"errors"
	"strings"

	"github.com/fatih/structs"
	"github.com/go-playground/validator/v10"
)

type SubmitTransactionRequest struct {
	CustomerID  string `json:"customer_id" validate:"required"`
	ProductCode string `json:"product_code" validate:"required"`
	SuccessRedirectURL string `json:"success_redirect_url" validate:"required,url"`
	FailureRedirectURL string `json:"failure_redirect_url" validate:"required,url"`
}

func (r *SubmitTransactionRequest) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(r); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(r)
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

type ReviewTransactionRequest struct {
	CustomerID  string `json:"customer_id" validate:"required"`
	ProductCode string `json:"product_code" validate:"required"`
}

func (r *ReviewTransactionRequest) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(r); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(r)
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