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

func NewCognitoClient(cognitoRegion string, userPoolId string, cognitoClientId string) AuthService {

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion))
	if err != nil {
		panic(err)
	}

	return &CognitoClient{
		AppClientId:   cognitoClientId,
		UserPoolId:    userPoolId,
		ClientCognito: cip.NewFromConfig(cfg),
	}

}

func (s *CognitoClient) SignIn(email string, password string) (string, error) {

	params := map[string]string{
		"USERNAME": *aws.String(email),
		"PASSWPRD": *aws.String(password),
	}

	signInInput := &cip.AdminInitiateAuthInput{

		AuthFlow:       types.AuthFlowTypeAdminUserPasswordAuth,
		AuthParameters: params,
		ClientId:       &s.AppClientId,
		UserPoolId:     &s.UserPoolId,
	}

	res, err := s.ClientCognito.AdminInitiateAuth(ctx, signInInput)

	if err != nil {
		fmt.Println(err)
		return "Notwork", err
	}
	fmt.Println(res)

	return "work", nil

	// Work

	// flow := aws.String("USER_PASSWORD_AUTH")
	// params := map[string]string{
	// 	"USERNAME": *aws.String(email),
	// 	"PASSWORD": *aws.String(password),
	// }

	// signInInput := &cip.InitiateAuthInput{
	// 	AuthFlow:       types.AuthFlowType(*flow),
	// 	AuthParameters: params,
	// 	ClientId:       &s.AppClientId,
	// }

	// res, err := s.ClientCog.InitiateAuth(ctx, signInInput)

	// if err != nil {
	// 	return "Error DB Conncetion", err
	// }

	// fmt.Println(res)

	// return *res.Session, nil

}
func (s *CognitoClient) SignUp(email string, password string, phoneNo string) (*cip.SignUpOutput, error) {

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
		return nil, err
	}

	return result, nil

}
