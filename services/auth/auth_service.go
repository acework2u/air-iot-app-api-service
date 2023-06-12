package auth

type CognitoClient struct {
	AppClientId string
	UserPoolId  string
	// *cip.Client
}

func NewCognitoClient(userPoolId string, cognitoClientId string) AuthService {

	// cfg, err := config.LoadDefaultConfig(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// return &CognitoClient{
	// 	AppClientId: cognitoClientId,
	// 	UserPoolId:  userPoolId,
	// 	cip.NewFromConfig(cfg),
	// }

	return &CognitoClient{
		AppClientId: cognitoClientId,
		UserPoolId:  userPoolId,
	}

}

func (s *CognitoClient) SignIn(email string, password string) (string, error) {

	return s.AppClientId, nil
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
