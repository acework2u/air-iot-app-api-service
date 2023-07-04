package clientcoginto

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoService struct {
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

	return &CognitoService{
		cognitoClient: client,
		appClientId:   cognitoClientId,
	}
}

// SigUp implements ClientCognito
func (sc *CognitoService) SignUp(email string, password string) (string, error) {

	//phone_no := 0945968514

	userName := strings.Split(email, "@")

	user := &cognito.SignUpInput{
		ClientId: &sc.appClientId,
		Username: aws.String(userName[0]),
		Password: aws.String(password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(userName[0]),
			},
		},
	}

	result, err := sc.cognitoClient.SignUp(user)

	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func (sc *CognitoService) ConfirmeSignUp(email string, code string) (string, error) {

	//phone_no := 0945968514

	confirmSignUpInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(email),
		ConfirmationCode: aws.String(code),
		ClientId:         aws.String(sc.appClientId),
	}

	result, err := sc.cognitoClient.ConfirmSignUp(confirmSignUpInput)

	if err != nil {
		//	fmt.Println(err)
		return "", err
	}

	//fmt.Println(result)

	return result.String(), nil
}

func (sc *CognitoService) SignIn(email string, password string) (string, *cognito.InitiateAuthOutput, error) {

	flow := aws.String("USER_PASSWORD_AUTH")

	userName := strings.Split(email, "@")
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

		fmt.Println(err.Error())

		return "Authentication Error" + userName[0], nil, err
	}

	// fmt.Println("res.AuthenticationResult")
	// fmt.Println(res.AuthenticationResult)
	fmt.Println(res)

	return res.String(), res, nil

}

func (sc *CognitoService) GetUserPoolId() (string, error) {
	return "", nil
}
