package users

import (
	"capstone/businesses/users"
	"capstone/helpers"
	"net/http"
	"strings"

	requests "capstone/controllers/users/requests"
	"capstone/controllers/users/response"
	appjwt "capstone/utils/jwt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Contains the echo handlres for the user

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUseCase users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	// Get request body
	request := requests.UserRegister{}
	c.Bind(&request)

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	user, err := ctrl.UserUseCase.Register(request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, helpers.Response{
		Status:  http.StatusCreated,
		Message: "Success creating user",
		Data:    response.FromDomain(&user),
	})
}

func (ctrl *UserController) Login(c echo.Context) error {
	// Get request body
	request := requests.UserLogin{}
	c.Bind(&request)

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	jwt_token, err := ctrl.UserUseCase.Login(request.ToDomain())
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.Response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success login",
		Data:    map[string]string{"token": jwt_token},
	})
}

func (ctrl *UserController) GetProfile(c echo.Context) error {
	// Get token from header
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id
	user_id := appjwt.GetID(tokenString)

	// Get user profile
	user, err := ctrl.UserUseCase.GetByID(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting user profile",
		Data:    response.FromDomain(&user),
	})
}

func (ctrl *UserController) GetAllUsers(c echo.Context) error {
	users, err := ctrl.UserUseCase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response.FromDomainArray(users),
	})
}

func (ctrl *UserController) GetUserByID(c echo.Context) error {
	user_id := c.Param("id")

	user, err := ctrl.UserUseCase.GetByID(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success getting user profile",
		Data:    response.FromDomain(&user),
	})
}

func (ctrl *UserController) UpdateUserByID(c echo.Context) error {
	user_id := c.Param("id")

	// Get request body
	request := requests.UserUpdateByAdmin{}
	c.Bind(&request)

	// domain
	domain := request.ToDomain()
	ObjID, _ := primitive.ObjectIDFromHex(user_id)
	domain.ID = ObjID

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Update user
	user, err := ctrl.UserUseCase.UpdateByAdmin(domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success updating user profile",
		Data:    response.FromDomain(&user),
	})
}

func (ctrl *UserController) UpdateProfile(c echo.Context) error {
	// Get token from header
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id
	user_id := appjwt.GetID(tokenString)

	// Get request body
	request := requests.UserUpdateProfile{}
	c.Bind(&request)

	// domain
	domain := request.ToDomain()
	ObjID, _ := primitive.ObjectIDFromHex(user_id)
	domain.ID = ObjID

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Update user
	user, err := ctrl.UserUseCase.UpdateProfile(domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success updating user profile",
		Data:    user,
	})
}

func (ctrl *UserController) UpdatePassword(c echo.Context) error {
	// Get token from header
	tokenString := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)

	// Get user id
	user_id := appjwt.GetID(tokenString)

	// Get request body
	request := requests.UserUpdatePassword{}
	c.Bind(&request)

	// domain
	domain := request.ToDomain()
	ObjID, _ := primitive.ObjectIDFromHex(user_id)
	domain.ID = ObjID

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	old := domain
	new := users.Domain{
		Password: request.NewPassword,
	}

	old.Password = request.OldPassword

	// Update user
	_, err := ctrl.UserUseCase.UpdatePassword(old, &new)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func (ctrl *UserController) DeleteUserByID(c echo.Context) error {
	user_id := c.Param("id")

	// Delete user
	user, err := ctrl.UserUseCase.DeleteByAdmin(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    response.FromDomain(&user),
	})
}

func (ctrl *UserController) RequestOTP(c echo.Context) error {
	// Get request body
	request := requests.UserRequestOTP{}
	c.Bind(&request)

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Request OTP
	_, err := ctrl.UserUseCase.RequestOTP(request.Email, request.Scope)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success requesting OTP",
		Data:    nil,
	})
}

func (ctrl *UserController) VerifyEmail(c echo.Context) error {
	// Get request body
	request := requests.UserVerifyEmail{}
	c.Bind(&request)

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Verify email
	_, err := ctrl.UserUseCase.VerifyEmail(request.Email, request.OTP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success verify email",
		Data:    nil,
	})
}

// Reset password
func (ctrl *UserController) ResetPassword(c echo.Context) error {
	// Get request body
	request := requests.UserResetPassword{}
	c.Bind(&request)

	// Validate request
	val_err := request.Validate()
	if val_err != nil {
		return c.JSON(http.StatusNotAcceptable, helpers.Response{
			Status:  http.StatusNotAcceptable,
			Message: "Validation error",
			Data:    val_err,
		})
	}

	// Reset password
	_, err := ctrl.UserUseCase.ResetPassword(request.Email, request.Password, request.OTP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success reset password",
		Data:    nil,
	})
}

// Get Total Users
func (ctrl *UserController) GetTotalUsers(c echo.Context) error {
	// Get total users
	total, err := ctrl.UserUseCase.GetTotalUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, helpers.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    total,
	})
}