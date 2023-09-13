package auth

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type AuthenServices interface {
	SignIn(string, string) (*cognitoidentityprovider.InitiateAuthOutput, error)
	SignUp(string, string, string) (string, error)
	UserConfirm(string, string) (interface{}, error)
	ResendConfirmCode(string) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	RefreshToken(refreshToken string) (interface{}, error)
	ForgotPassword(userName string) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ConfirmNewPassword(*UserConfirmNewPassword) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ChangePassword(changeReq *ChangePasswordReq) (interface{}, error)
	DeleteMyAccount(accessKey string) error
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
	UserDelete struct {
		AccessToken string `json:"accessToken" validate:"required" binding:"required"`
	}
	UserConfirmNewPassword struct {
		UserName    string `json:"userName" validate:"required" binding:"required"`
		Password    string `json:"password" validate:"required,max=10,min=1" binding:"required"`
		ConfirmCode string `json:"confirmCode" validate:"required" binding:"required"`
	}
	ChangePasswordReq struct {
		AccessToken      string `json:"accessToken" validate:"required" binding:"required"`
		PreviousPassword string `json:"currentPassword" validate:"required" binding:"required"`
		ProposePassword  string `json:"newPassword" validate:"required" binding:"required"`
	}

	ResendConfirmCode struct {
		Username string `json:"username" validate:"required" binding:"required"`
	}

	ResponseForgotPassword struct {
		CodeDeliveryDetails struct {
			AttributeName  string `json:"AttributeName"`
			DeliveryMedium string `json:"DeliveryMedium"`
			Destination    string `json:"Destination"`
		} `json:"CodeDeliveryDetails"`
		ResultMetadata struct {
		} `json:"ResultMetadata"`
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
