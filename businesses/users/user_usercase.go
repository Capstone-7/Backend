package users

import (
	appjwt "capstone/utils/jwt"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

//  Contains all the business logic for the user

type UserUseCase struct {
	UserRepository Repository
}

func NewUserUseCase(userRepository Repository) UseCase {
	return &UserUseCase{
		UserRepository: userRepository,
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
	new.Password = string(password)

	return u.UserRepository.Update(&user)
}

func (u *UserUseCase) RequestOTP(domain *Domain) (string, error) {
	panic("implement me")
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