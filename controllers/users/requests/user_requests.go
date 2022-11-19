package requests

import (
	"capstone/businesses/users"
	"capstone/helpers"
	"errors"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contains requests definition

type UserRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=!@#$%^&*,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
	Name	 string `json:"name" validate:"required,min=3,max=50"`
}

func (u *UserRegister) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

func (u *UserRegister) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
	}
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=!@#$%^&*,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}

func (u *UserLogin) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

func (u *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    u.Email,
		Password: u.Password,
	}
}

type UserUpdateByAdmin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=!@#$%^&*,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
	Name	 string `json:"name" validate:"required,min=3,max=50"`
	Status	 string `json:"status" validate:"required,min=3,max=50"`
	IsActivated bool `json:"is_activated"`
}

func (u *UserUpdateByAdmin) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
			out := make([]helpers.ValidationErrors, 0)

			for i, e := range ve {
				if e.Field() == "Password" && u.Password == "" {
					continue
				}

				out = append(out, helpers.ValidationErrors{
					Field:   e.Field(),
					Message: helpers.MsgForTag(e.Tag()),
				})

				out[i].Message = strings.Replace(out[i].Message, "[PARAM]", e.Param(), 1)

				// Get field tag
				for _, f := range fields {
					if f.Name() == e.Field() {
						out[i].Field = f.Tag("json")
						break
					}
				}
			}

			if len(out) == 0 {
				return nil
			}

			return out
		}
	}

	return nil
}

func (u *UserUpdateByAdmin) ToDomain() *users.Domain {
	// Domain
	domain := &users.Domain{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
		Status:   u.Status,
	}

	if u.IsActivated {
		// Custom Time
		theTime := time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)
		domain.Deleted = primitive.NewDateTimeFromTime(theTime)
	}

	return domain
}

type UserUpdatePassword struct {
	OldPassword string `json:"old_password" validate:"required,min=8,max=12,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=!@#$%^&*,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=12,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=!@#$%^&*,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}

func (u *UserUpdatePassword) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

func (u *UserUpdatePassword) ToDomain() *users.Domain {
	return &users.Domain{
		Password: u.NewPassword,
	}
}

type UserUpdateProfile struct {
	Name	 string `json:"name" validate:"required,min=3,max=50"`
	Email	string `json:"email" validate:"required,email"`
}

func (u *UserUpdateProfile) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

func (u *UserUpdateProfile) ToDomain() *users.Domain {
	return &users.Domain{
		Name: u.Name,
		Email: u.Email,
	}
}

type UserRequestOTP struct {
	Email string `json:"email" validate:"required,email"`
	Scope string `json:"scope" validate:"required"`
}

func (u *UserRequestOTP) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

func (u *UserRequestOTP) ToDomain() *users.Domain {
	return &users.Domain{
		Email: u.Email,
	}
}

type UserVerifyEmail struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"	" validate:"required,len=6"`
}

func (u *UserVerifyEmail) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

type UserResetPassword struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required,len=6"`
	Password string `json:"password" validate:"required,min=8,max=12,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=!@#$%^&*,containsany=abcdefghijklmnopqrstuvwxyz,containsany=0123456789"`
}

func (u *UserResetPassword) Validate() []helpers.ValidationErrors {
	var ve validator.ValidationErrors

	if err := validator.New().Struct(u); err != nil {
		if errors.As(err, &ve) {
			fields := structs.Fields(u)
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

func (u *UserResetPassword) ToDomain() *users.Domain {
	return &users.Domain{
		Email: u.Email,
		Password: u.Password,
	}
}