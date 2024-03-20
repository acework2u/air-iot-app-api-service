package auth

import (
	"context"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var ctx = context.TODO()

// var customerCollection = configs.GetCollection(mongoclient, "customers")
// var customerRepository = repository.NewCustomerRepositoryDB(customerCollection, ctx)

type CognitoClient struct {
	AppClientId   string
	UserPoolId    string
	ClientCognito *cip.Client
	cusRepo       repository.CustomerRepository
}

func NewCognitoClient(cognitoRegion string, userPoolId string, cognitoClientId string, cusRepo repository.CustomerRepository) AuthenServices {

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion))
	if err != nil {
		panic(err)
	}

	return &CognitoClient{
		AppClientId:   cognitoClientId,
		UserPoolId:    userPoolId,
		ClientCognito: cip.NewFromConfig(cfg),
		cusRepo:       cusRepo,
	}

}

func (s *CognitoClient) SignIn(signInReq SignInRequest) (*cip.InitiateAuthOutput, error) {

	// Work
	flow := aws.String("USER_PASSWORD_AUTH")
	//flow := aws.String("USER_SRP_AUTH")
	params := map[string]string{
		"USERNAME":   *aws.String(signInReq.Username),
		"PASSWORD":   *aws.String(signInReq.Password),
		"DEVICE_KEY": *aws.String(signInReq.DeviceNo),
	}

	signInInput := &cip.InitiateAuthInput{
		AuthFlow:       types.AuthFlowType(*flow),
		AuthParameters: params,
		ClientId:       &s.AppClientId,
	}

	result, err := s.ClientCognito.InitiateAuth(ctx, signInInput)

	if err != nil {
		return nil, err
	}

	//return *res.Session, nil
	return result, nil

}

func (s *CognitoClient) SignUp(signUpReq SignUpRequest) (string, error) {

	userSignUp := &cip.SignUpInput{
		ClientId: &s.AppClientId,
		Username: aws.String(signUpReq.Username),
		Password: aws.String(signUpReq.Password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(signUpReq.Username),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(signUpReq.PhoneNo),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(signUpReq.Name),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(signUpReq.LastName),
			},
			{
				Name:  aws.String("custom:role"),
				Value: aws.String("1"),
			},
		},
	}

	result, err := s.ClientCognito.SignUp(ctx, userSignUp)

	if err != nil {
		return "", err
	}

	// Register Customer success
	userInfo := &repository.CreateCustomerRequest2{
		UserSub:       *result.UserSub,
		Name:          signUpReq.Name,
		Lastname:      signUpReq.LastName,
		Email:         signUpReq.Username,
		UserConfirmed: result.UserConfirmed,
		Mobile:        signUpReq.PhoneNo,
		Role:          signUpReq.CustomRole,
	}

	_, ok := s.cusRepo.NewCustomer(userInfo)

	msgSuccess := fmt.Sprintf("ลงทะเบียนสำเร็จ กรุณายืนยันข้อมูลที่ email:  %v", *result.CodeDeliveryDetails.Destination)

	//_ = msgSuccess

	if ok != nil {
		return "", err
	}

	return msgSuccess, nil

}

func (s *CognitoClient) UserConfirm(username string, confirmCode string) (interface{}, error) {

	confirmSignUpInput := &cip.ConfirmSignUpInput{
		Username:         aws.String(username),
		ConfirmationCode: aws.String(confirmCode),
		ClientId:         aws.String(s.AppClientId),
	}

	result, err := s.ClientCognito.ConfirmSignUp(ctx, confirmSignUpInput)

	if err != nil {
		return "", err
	}

	return result, nil

}

func (s *CognitoClient) ResendConfirmCode(username string) (*cip.ResendConfirmationCodeOutput, error) {

	resendConfirmCodeInput := &cip.ResendConfirmationCodeInput{
		ClientId: aws.String(s.AppClientId),
		Username: aws.String(username),
	}

	resConfOut, err := s.ClientCognito.ResendConfirmationCode(ctx, resendConfirmCodeInput)
	if err != nil {
		return nil, err
	}

	return resConfOut, nil

}

func (s *CognitoClient) RefreshToken(refreshToken string) (interface{}, error) {

	param := map[string]string{
		"REFRESH_TOKEN": *aws.String(refreshToken),
	}

	refreshTokenInput := &cip.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeRefreshToken,
		ClientId:       aws.String(s.AppClientId),
		AuthParameters: param,
	}
	initiateAuthOutput, err := s.ClientCognito.InitiateAuth(ctx, refreshTokenInput)

	//param := map[string]string{
	//	"REFRESH_TOKEN": *aws.String(refreshToken),
	//	"DEVICE_KEY":    *aws.String("R3FR3SHT0K3N"),
	//}
	//adminRefreshTokenInput := &cip.AdminInitiateAuthInput{
	//	AuthFlow:       types.AuthFlowTypeRefreshTokenAuth,
	//	ClientId:       aws.String(s.AppClientId),
	//	UserPoolId:     aws.String(s.UserPoolId),
	//	AuthParameters: parama,
	//}
	//initiateAuthOutput, err := s.ClientCognito.AdminInitiateAuth(ctx, adminRefreshTokenInput)
	if err != nil {
		return nil, err
	}
	return initiateAuthOutput, nil
}

func (s *CognitoClient) ChangePassword(changeReq *ChangePasswordReq) (interface{}, error) {

	changePwInput := &cip.ChangePasswordInput{
		AccessToken:      aws.String(changeReq.AccessToken),
		PreviousPassword: aws.String(changeReq.PreviousPassword),
		ProposedPassword: aws.String(changeReq.ProposePassword),
	}
	changePasswordOutput, err := s.ClientCognito.ChangePassword(ctx, changePwInput)
	if err != nil {
		return nil, err
	}
	return changePasswordOutput, nil
}

func (s *CognitoClient) ForgotPassword(userName string) (*cip.ForgotPasswordOutput, error) {
	// ClientId: aws.String(s.AppClientId),
	forgotPasswordInput := &cip.ForgotPasswordInput{
		Username: aws.String(userName),
		ClientId: aws.String(s.AppClientId)}

	forgotPasswordOutput, err := s.ClientCognito.ForgotPassword(ctx, forgotPasswordInput)
	if err != nil {
		return nil, err
	}
	return forgotPasswordOutput, nil
}

func (s *CognitoClient) ConfirmNewPassword(userConfirm *UserConfirmNewPassword) (*cip.ConfirmForgotPasswordOutput, error) {

	confirmForgotPasswordInput := &cip.ConfirmForgotPasswordInput{
		ClientId:         aws.String(s.AppClientId),
		ConfirmationCode: aws.String(userConfirm.ConfirmCode),
		Username:         aws.String(userConfirm.UserName),
		Password:         aws.String(userConfirm.Password),
	}

	cfpPwOutput, err := s.ClientCognito.ConfirmForgotPassword(ctx, confirmForgotPasswordInput)

	if err != nil {
		return nil, err
	}

	return cfpPwOutput, nil
}

func (s *CognitoClient) DeleteMyAccount(accessKey string) error {
	_, err := s.ClientCognito.DeleteUser(ctx, &cip.DeleteUserInput{AccessToken: aws.String(accessKey)})
	if err != nil {
		return err
	}
	return nil
}

func (s *CognitoClient) ConfirmDevice(deviceInput *DeviceConfirmReq) error {
	var err error

	deviceConfirm := &cip.ConfirmDeviceInput{
		AccessToken: aws.String(deviceInput.AccessToken),
		DeviceKey:   aws.String(deviceInput.DeviceKey),
	}

	confirmDeviceOut, err := s.ClientCognito.ConfirmDevice(ctx, deviceConfirm)

	if err != nil {
		return err
	}
	// Confirm Success
	if confirmDeviceOut.UserConfirmationNecessary {
		return nil
	}
	return err
}
