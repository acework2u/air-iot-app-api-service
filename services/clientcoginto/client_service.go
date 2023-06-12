package clientcoginto

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type cognitoService struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
}

func NewCognitoService(cognitoRegion string, cognitoClientId string) ClientCognito {
	conf := &aws.Config{
		Region: aws.String(cognitoRegion),
	}

	sess, err := session.NewSession(conf)
	client := cognito.New(sess)

	if err != nil {
		panic(err)
	}

	return &cognitoService{
		cognitoClient: client,
		appClientId:   cognitoClientId,
	}
}

// SigUp implements ClientCognito
func (sc *cognitoService) SignUp(emil string, password string) (string, error) {

	//phone_no := 0945968514

	user := &cognito.SignUpInput{
		ClientId: &sc.appClientId,
		Username: aws.String(emil),
		Password: aws.String(password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("phone_number"),
				Value: aws.String("+66945968514"),
			},
		},
	}

	result, err := sc.cognitoClient.SignUp(user)

	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func (sc *cognitoService) ConfirmeSignUp(email string, code string) (string, error) {

	//phone_no := 0945968514

	confirmSignUpInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
		ClientId:         aws.String(sc.appClientId),
	}

	result, err := sc.cognitoClient.ConfirmSignUp(confirmSignUpInput)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(result)

	return result.String(), nil
}

func (sc *cognitoService) SignIn(email string, password string) (string, *cognito.InitiateAuthOutput, error) {

	flow := aws.String("USER_PASSWORD_AUTH")

	params := map[string]*string{
		"USERNAME": aws.String(email),
		"PASSWORD": aws.String(password),
	}

	authTry := &cognito.InitiateAuthInput{
		AuthFlow:       flow,
		AuthParameters: params,
		ClientId:       &sc.appClientId,
	}

	res, err := sc.cognitoClient.InitiateAuth(authTry)

	if err != nil {
		return "", nil, err
	}

	fmt.Println("res.AuthenticationResult")

	return res.String(), res, nil

}
