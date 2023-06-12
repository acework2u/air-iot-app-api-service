package auth

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/aws"
)

type CognitoClient struct {
	AppClientId string
	UserPoolId  string

	ClientCog *cip.Client
}

func NewCognitoClient(cognitoRegion string, userPoolId string, cognitoClientId string) AuthService {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}
	cfg.Region = *aws.String(cognitoRegion)

	// return &CognitoClient{
	// 	AppClientId: cognitoClientId,
	// 	UserPoolId:  userPoolId,
	// 	cip.NewFromConfig(cfg),
	// }

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

	// res, err := s.ClientCog.AdminInitiateAuth(context.Background(), signInInput)

	// if err != nil {
	// 	return "", err
	// }

	return "", nil
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
