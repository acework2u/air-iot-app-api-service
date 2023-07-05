package clientcoginto

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-playground/validator/v10"
)

type ClientSignUp struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ClientSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
	ConfirmSignUp(string, string) (string, error)
	SignIn(string, string) (string, *cognito.InitiateAuthOutput, error)
	GetUserPoolId() (string, error)
}
