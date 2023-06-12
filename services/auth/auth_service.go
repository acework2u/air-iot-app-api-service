package auth

import (
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	*cip.Client
	AppClientId string
	UserPoolId  string
}

func NewCognitoClient(cognitoRegion string, cognitoClientId string) AuthService {

	// cfg, err := awsconfig.LoadDefaultConfig(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// return &CognitoClient{
	// 	AppClientId: cognitoClientId,
	// 	UserPoolId: "",
	// 	cip.NewFromConfig(cfg),
	// }

	return nil

}

func (s *CognitoClient) SignIn(string, string) (string, error) {
	return "", nil
}
func (s *CognitoClient) SignUp(string, string) (string, error) {
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
