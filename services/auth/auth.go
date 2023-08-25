package auth

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type AuthenServices interface {
	SignIn(string, string) (*cognitoidentityprovider.InitiateAuthOutput, error)
	SignUp(string, string, string) (string, error)
	UserConfirm(string, string) (interface{}, error)
	ResendConfirmCode(string) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	RefreshToken(refreshToken string) (interface{}, error)
}

type (
	SignInRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserConfirm struct {
		ConfirmationCode string `json:"confirmationCode" validate:"required" binding:"required"`
		User             string `json:"username" validate:"required" binding:"required"`
	}

	ResendConfirmCode struct {
		Username string `json:"username" validate:"required" binding:"required"`
	}

	SignUpRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		PhoneNo  string `json:"phone_no" binding:"required"`
	}
	SignInResponse struct {
		// The access token.
		AccessToken *string `json:"access_token,omitempty"`

		// The expiration period of the authentication result in seconds.
		ExpiresIn int32 `json:"expires_in,omitempty"`

		// The ID token.
		IdToken *string `json:"id_token,omitempty"`

		// The refresh token.
		RefreshToken *string `json:"refresh_token,omitempty"`

		// The token type.
		TokenType *string `json:"token_type"`
	}
)
