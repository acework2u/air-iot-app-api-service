package auth

type AuthService interface {
	SignIn(string, string) (string, error)
	SignUp(string, string) (string, error)
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
