package auth

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type AuthService interface {
	SignIn(string, string) (string, error)
	SignUp(string, string, string) (*cognitoidentityprovider.SignUpOutput, error)
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone_no string `json:"phone_no" binding:"required"`
}
type SignInResponse struct {
	// The access token.
	AccessToken *string `json:"access_token"`

	// The expiration period of the authentication result in seconds.
	ExpiresIn int32 `json:"expires_in"`

	// The ID token.
	IdToken *string `json:"id_token"`

	// The refresh token.
	RefreshToken *string `json:"refresh_token"`

	// The token type.
	TokenType *string `json:"token_type"`
}
