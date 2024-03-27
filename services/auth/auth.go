package auth

import "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"

type AuthenServices interface {
	SignIn(signInReq SignInRequest) (*cognitoidentityprovider.InitiateAuthOutput, error)
	SignUp(signUp SignUpRequest) (string, error)
	UserConfirm(string, string) (interface{}, error)
	ResendConfirmCode(string) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error)
	RefreshToken(refreshToken string) (interface{}, error)
	ForgotPassword(userName string) (*cognitoidentityprovider.ForgotPasswordOutput, error)
	ConfirmNewPassword(*UserConfirmNewPassword) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error)
	ChangePassword(changeReq *ChangePasswordReq) (interface{}, error)
	DeleteMyAccount(accessKey string) error
	ConfirmDevice(deviceInput *DeviceConfirmReq) error
}

type (
	// swagger:model SignInRequest
	SignInRequest struct {
		// Username
		// A Username of user authentication
		// in:body
		// Type: String
		// Required: Yes
		Username string `json:"username" form:"username" binding:"required"`
		// Password
		// A Password of user authentication
		// in:body
		// Type: String
		// Required: Yes
		Password string `json:"password" form:"password" binding:"required"`
		// DeviceNo
		// A device no of user authentication
		// in:body
		// Type: String
		// Required: Yes
		DeviceNo string `json:"device_no" form:"device_no" binding:"required"`
	}
	// swagger:model UserConfirm
	UserConfirm struct {
		// Username
		// A Username of confirmation
		// in:body
		// Type: String
		// Required: Yes
		User string `json:"username" validate:"required" binding:"required"`
		// ConfirmationCode
		// A code of user confirmation
		// in:body
		// Type: String
		// Required: Yes
		ConfirmationCode string `json:"confirmationCode" validate:"required" binding:"required"`
	}
	// swagger:model UserDelete
	UserDelete struct {
		// AccessToken
		// A AccessToken of delete user account
		// in:body
		// Type: String
		// Required: Yes
		AccessToken string `json:"accessToken" validate:"required" binding:"required"`
	}
	// swagger:model UserConfirmNewPassWord
	UserConfirmNewPassword struct {
		// UserName
		// A username of confirm new password
		// in:body
		// Type: String
		// Required: Yes
		UserName string `json:"userName" validate:"required" binding:"required"`
		// Password
		// A password of confirm new password
		// in:body
		// Type: String
		// Required: Yes
		Password string `json:"password" validate:"required,max=10,min=1" binding:"required"`
		// ConfirmCode
		//Confirm code sent the confirmation to the client.
		// in:body
		// Type: String
		// Required: Yes
		ConfirmCode string `json:"confirmCode" validate:"required" binding:"required"`
	}
	// swagger:model ChangePassWordReq
	ChangePasswordReq struct {
		// AccessToken
		// A AccessToken of Access to user authorized
		// in:ChangePassWordReq
		// Type: String
		// Required: Yes
		AccessToken string `json:"accessToken" validate:"required" binding:"required"`
		// PreviousPassword
		// A previous password of user authentication
		// in:ChangePassWordReq
		// Type: String
		// Required: Yes
		PreviousPassword string `json:"currentPassword" validate:"required" binding:"required"`
		// ProposePassWord
		// in:ChangePasswordReq
		// Type: String
		// Required: Yes
		ProposePassword string `json:"newPassword" validate:"required" binding:"required"`
	}
	// swagger:model ResendConfirmCode
	ResendConfirmCode struct {
		Username string `json:"username" validate:"required,email" binding:"required,email"`
	}
	// swagger:model ResponseForgotPassword
	ResponseForgotPassword struct {
		// swagger:model CodeDeliveryDetails
		CodeDeliveryDetails struct {
			// AttributeName
			// in:CodeDeliveryDetails
			// Type: String
			AttributeName string `json:"AttributeName"`
			// DeliveryMedium
			// in:CodeDeliveryDetails
			// Type: String
			DeliveryMedium string `json:"DeliveryMedium"`
			// Destination
			// in:CodeDeliveryDetails
			// Type: String
			Destination string `json:"Destination"`
		} `json:"CodeDeliveryDetails"` // @name CodeDeliveryDetails
		// ResultMetadata
		// in:ResponseForgotPassword
		// Type: String
		ResultMetadata struct {
		} `json:"ResultMetadata"`
	}

	// swagger:model SignUpRequest
	SignUpRequest struct {
		// Username is Email of user authentication
		// in:body
		// Type: String
		// Required: Yes
		// example: my-email@mail.com
		Username string `json:"username" binding:"required,email"`
		// Password of user authentication
		// in:body
		// Type: String
		// Required: Yes
		// example: P@assWord1234
		Password string `json:"password" binding:"required"`
		// The mobile number of user
		// in:body
		// Type: String
		// Required: Yes
		// example: 0941234567 for Thailand
		PhoneNo string `json:"phone_no" binding:"required,number,len=10"`
		// The name of user
		// A unique identifier for new user
		// in:body
		// Type: String
		// Required: Yes
		// example: anon
		Name string `json:"name" binding:"required,min=2,max=20"`
		// Last name of user
		// in:body
		// Type: String
		// Required: Yes
		// example: dechpala
		LastName string `json:"lastName" binding:"required,min=2,max=100"`
		// customRole
		// Role of access to our resource
		// in:body
		// Type: String
		// Required: No (option)
		// example: 1 is role default
		CustomRole string `json:"customRole"`
	}
	// swagger:model SignInResponse
	SignInResponse struct {
		// AccessToken
		// The AccessToken of access resource
		// in:body
		// Type: String
		AccessToken *string `json:"access_token,omitempty"`
		// ExpiresIn
		// The expiration period of the authentication result in seconds.
		// in:body
		// Type: Int32
		ExpiresIn int32 `json:"expires_in,omitempty"`
		// idToken
		// The ID token.
		// in:body
		// Type: String
		IdToken *string `json:"id_token,omitempty"`
		// RefreshToken
		// The refresh token.
		// in:body
		// Type: String
		RefreshToken *string `json:"refresh_token,omitempty"`
		// TokenType
		// The token type.
		// in:body
		// Type: String
		TokenType *string `json:"token_type"`
	}
	// swagger:model DeviceConfirmReq
	DeviceConfirmReq struct {
		// AccessToken
		// The AccessToken of access to our resource
		// in:body
		// Type: String
		// Required: Yes
		AccessToken string `json:"accessToken" binding:"required"`
		// DeviceKey
		// The DeviceKey of user device
		// in:body
		// Type: String
		// Required: Yes
		DeviceKey string `json:"deviceKey" binding:"required"`
	}
)
