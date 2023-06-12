package auth

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var ctx = context.TODO()

type CognitoClient struct {
	AppClientId string
	UserPoolId  string

	ClientCog *cip.Client
}

func NewCognitoClient(cognitoRegion string, userPoolId string, cognitoClientId string) AuthService {

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion))
	if err != nil {
		panic(err)
	}

	return &CognitoClient{
		AppClientId: cognitoClientId,
		UserPoolId:  userPoolId,
		ClientCog:   cip.NewFromConfig(cfg),
	}

}

func (s *CognitoClient) SignIn(email string, password string) (string, error) {

	// flow := aws.String("ADMIN_USER_PASSWORD_AUTH")
	// params := map[string]string{
	// 	"USERNAME": *aws.String(email),
	// 	"PASSWPRD": *aws.String(password),
	// }

	// signInInput := &cip.AdminInitiateAuthInput{

	// 	AuthFlow:       types.AuthFlowType(*flow),
	// 	AuthParameters: params,
	// 	ClientId:       &s.AppClientId,
	// 	UserPoolId:     &s.UserPoolId,
	// }

	// res, err := s.ClientCog.AdminInitiateAuth(ctx, signInInput)

	// if err != nil {
	// 	return "", err
	// }

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

	res, err := s.ClientCog.InitiateAuth(ctx, signInInput)

	if err != nil {
		return "Error DB Conncetion", err
	}

	fmt.Println(res)

	return "Connect DB Success", nil
}
func (s *CognitoClient) SignUp(email string, passeorf string) (string, error) {
	return "", nil
}

// func init() *CognitoClient {
// 	cfg, err := config.LoadDefaultConfig(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &CognitoClient{
// 		os.Getenv("COGNITO_APP_CLIENT_ID"),
// 		os.Getenv("COGNITO_USER_POOL_ID"),
// 		cip.NewFromConfig(cfg),
// 	}
// }
