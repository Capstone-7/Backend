package users

import (
	"capstone/businesses/otp"
	appjwt "capstone/utils/jwt"
	"errors"
	"fmt"
	"time"

	smtp "capstone/utils/smtp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//  Contains all the business logic for the user

type UserUseCase struct {
	UserRepository Repository
	OTPRepository  otp.Repository
}

func NewUserUseCase(userRepository Repository, otpRepository otp.Repository) UseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		OTPRepository: otpRepository,
	}
}

func (u *UserUseCase) Register(domain *Domain) (Domain, error) {
	// Check if email already exists
	_, err := u.UserRepository.GetByEmail(domain.Email)
	if err == nil {
		return Domain{}, errors.New("email already exists")
	}

	// Encrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.MinCost)
	domain.Password = string(password)
	domain.Status = "not_verified"

	domain.Created = primitive.NewDateTimeFromTime(time.Now())
	domain.Updated = primitive.NewDateTimeFromTime(time.Now())
	domain.Deleted = primitive.NewDateTimeFromTime(time.Time{})

	// Force change role to user
	domain.Role = "user"

	// Create user
	user, err := u.UserRepository.Create(domain)
	return user, err
}

func (u *UserUseCase) Login(domain *Domain) (string, error) {
	// Get user by email
	user, err := u.UserRepository.GetByEmail(domain.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Check is user active
	if user.Status != "verified" {
		return "", errors.New("user is not active")
	}

	// Check is user deleted
	if user.Deleted != primitive.NewDateTimeFromTime(time.Time{}) {
		return "", errors.New("user is deleted")
	}

	// Check password bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(domain.Password))
	if err != nil {
		return "", errors.New("password is incorrect")
	}

	// Generate JWT
	token, err := appjwt.GenerateToken(user.ID.Hex(), user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) GetProfile(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return u.UserRepository.GetByID(ObjId)
}

func (u *UserUseCase) UpdateProfile(domain *Domain) (Domain, error) {
	// Get user by id
	user, err := u.UserRepository.GetByID(domain.ID)
	if err != nil {
		return Domain{}, err
	}

	if domain.Email != user.Email {
		// Check if email already exists
		_, err = u.UserRepository.GetByEmail(domain.Email)
		if err == nil {
			return Domain{}, errors.New("email already exists")
		}

		// Update status to not verified
		user.Status = "not_verified"
	}

	user.Email = domain.Email
	user.Name = domain.Name
	user.Updated = primitive.NewDateTimeFromTime(time.Now())

	return u.UserRepository.Update(&user)
}

func (u *UserUseCase) UpdatePassword(old *Domain, new *Domain) (Domain, error) {
	// Get user by id
	user, err := u.UserRepository.GetByID(old.ID)
	if err != nil {
		return Domain{}, err
	}

	// Check password bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old.Password))
	if err != nil {
		return Domain{}, errors.New("old password is incorrect")
	}

	// Bcrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	user.Password = string(password)

	return u.UserRepository.Update(&user)
}

func (u *UserUseCase) GetAllUsers() ([]Domain, error) {
	return u.UserRepository.GetAll()
}

func (u *UserUseCase) GetByID(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return u.UserRepository.GetByID(ObjId)
}

func (u *UserUseCase) UpdateByAdmin(new *Domain) (Domain, error) {
	// Get user by id
	user, err := u.UserRepository.GetByID(new.ID)
	if err != nil {
		return Domain{}, err
	}

	user.Status = new.Status

	if new.Email != user.Email {
		// Check if email already exists
		_, err = u.UserRepository.GetByEmail(new.Email)
		if err == nil {
			return Domain{}, errors.New("email already exists")
		}

		// Update status to not verified
		user.Status = "not_verified"
	}

	if new.Password != "" {
		// Bcrypt password
		password, _ := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
		user.Password = string(password)
	}

	// Check for is deleted
	theTime := time.Date(1999, time.January, 1, 0, 0, 0, 0, time.UTC)
	if new.Deleted == primitive.NewDateTimeFromTime(theTime) {
		user.Deleted = primitive.NewDateTimeFromTime(time.Time{})
	}

	return u.UserRepository.Update(&user)
}

func (u *UserUseCase) DeleteByAdmin(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return u.UserRepository.Delete(ObjId)
}

func (u *UserUseCase) RequestOTP(email, scope string) (string, error) {
	// Get user by email
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Check is user deleted
	if user.Deleted != primitive.NewDateTimeFromTime(time.Time{}) {
		return "", errors.New("user is deleted")
	}

	// Get Last OTP
	otp, err := u.OTPRepository.GetLastByUserID(user.ID.Hex())
	if err == nil {
		fmt.Println(otp)
		// Check is otp created 1 minute ago
		if time.Now().Sub(otp.Created.Time()).Minutes() < 1 {
			return "", errors.New("please wait 1 minute to request new OTP")
		}
	}

	// Generate OTP
	otp, err = u.OTPRepository.GenerateOTP(user.ID.Hex(), scope)
	if err != nil {
		return "", err
	}

	// Send OTP to email in goroutine
	go func() {
		// Send email
		err = smtp.SendOTP([]string{user.Email}, otp.Code)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return "OTP sent to email "+user.Email, nil
}

func (u *UserUseCase) VerifyEmail(email, code string) (string, error) {
	// Get user by email
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Check is user deleted
	if user.Deleted != primitive.NewDateTimeFromTime(time.Time{}) {
		return "", errors.New("user is deleted")
	}

	// Verify OTP
	req := otp.Domain{
		UserID: user.ID,
		Code: code,
		Scope: "verify-email",
	}

	_, err = u.OTPRepository.VerifyOTP(&req)
	if err != nil {
		return "", errors.New("invalid OTP")
	}

	// Update status to verified
	user.Status = "verified"
	user.Updated = primitive.NewDateTimeFromTime(time.Now())

	_, err = u.UserRepository.Update(&user)

	// Consume OTP
	_, err = u.OTPRepository.ConsumeOTP(&req)

	return "Email verified", nil
}

func (u *UserUseCase) ResetPassword(email, new_password, code string) (Domain, error) {
	// Get user by email
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return Domain{}, errors.New("email not found")
	}

	// Check is user deleted
	if user.Deleted != primitive.NewDateTimeFromTime(time.Time{}) {
		return Domain{}, errors.New("user is deleted")
	}

	// Verify OTP
	req := otp.Domain{
		UserID: user.ID,
		Code: code,
		Scope: "reset-password",
	}

	_, err = u.OTPRepository.VerifyOTP(&req)
	if err != nil {
		return Domain{}, errors.New("invalid OTP")
	}

	// Bcrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.MinCost)
	user.Password = string(password)

	// Update password
	user, err = u.UserRepository.Update(&user)

	// Consume OTP
	_, err = u.OTPRepository.ConsumeOTP(&req)

	return user, nil
}

// Get Total Users
func (u *UserUseCase) GetTotalUsers() (int64, error) {
	return u.UserRepository.Count()
}