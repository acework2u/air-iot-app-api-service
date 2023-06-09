package clientcoginto

import "github.com/go-playground/validator/v10"

type ClientSignUp struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type (
	User struct {
		Username string `json:"username" validate:"required" db:"username"`
		Password string `json:"password" validate:"required"`
	}
	UserForgot struct {
		Username string `json:"username" validate:"required"`
	}
	UserConfirm struct {
		ConfirmationCode string `json:"confirmationcode" validate:"required" binding:"required"`
		User             string `json:"username" validate:"required" binding:"required" `
	}
	UserRegister struct {
		Email string `json:"email" validate:"required"`
		User  User   `json:"user" validate:"required"`
	}
	OTP struct {
		Username string `json:"username"`
		OTP      string `json:"otp"`
	}
	Response struct {
		Error error `json:"error"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

type ClientCognito interface {
	SignUp(string, string) (string, error)
	ConfirmeSignUp(string, string) (string, error)
}
