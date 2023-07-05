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

func (s *CognitoClient) SignIn(email string, password string) (*cip.InitiateAuthOutput, error) {

	//params := map[string]string{
	//	"USERNAME": *aws.String(email),
	//	"PASSWPRD": *aws.String(password),
	//}

	//signInInput := &cip.AdminInitiateAuthInput{
	//
	//	AuthFlow:       types.AuthFlowTypeAdminUserPasswordAuth,
	//	AuthParameters: params,
	//	ClientId:       &s.AppClientId,
	//	UserPoolId:     &s.UserPoolId,
	//}
	//
	//res, err := s.ClientCognito.AdminInitiateAuth(ctx, signInInput)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return "Notwork", err
	//}
	//fmt.Println(res)
	//
	//return "work", nil

	// Work

	flow := aws.String("USER_PASSWORD_AUTH")
	params := map[string]string{
		"USERNAME": *aws.String(email),
		"PASSWORD": *aws.String(password),
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
func (s *CognitoClient) SignUp(email string, password string, phoneNo string) (string, error) {

	userSignUp := &cip.SignUpInput{
		ClientId: &s.AppClientId,
		Username: aws.String(email),
		Password: aws.String(password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(phoneNo),
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
		Email:         email,
		UserConfirmed: result.UserConfirmed,
		Mobile:        phoneNo,
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
