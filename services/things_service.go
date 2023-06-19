package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	_ "github.com/aws/aws-sdk-go/service/iot"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type CogClient struct {
	AppClientId string
	UserPoolId  string
	ClientCog   *cip.Client
	IotClient   *iot.Client
	StsSvc      *sts.Client
	Cfg         *aws.Config
}

func NewThingClient(cognitoRegion string, userPoolId string, cognitoClientId string) ThinksService {

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion), config.WithSharedConfigProfile("air_iot_dev"))
	if err != nil {
		log.Fatalln("Failed to load AWS config:", err)
	}
	cognitoIdentityProviderClient := cip.NewFromConfig(cfg)
	stsClient := sts.NewFromConfig(cfg)
	iotClient := iot.NewFromConfig(cfg)

	return &CogClient{
		AppClientId: cognitoClientId,
		UserPoolId:  userPoolId,
		ClientCog:   cognitoIdentityProviderClient,
		StsSvc:      stsClient,
		IotClient:   iotClient,
		Cfg:         &cfg,
	}

}

func (s *CogClient) GetCerds() (interface{}, error) {

	username := "tidosi6511@vaband.com"
	password := "Pass@word2020"

	// Authenticate the user and retrieve the Cognito ID token
	authResult, err := s.ClientCog.InitiateAuth(context.TODO(), &cip.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
		ClientId: aws.String(s.AppClientId),
		// UserPoolId: aws.String(s.UserPoolId),
	})
	if err != nil {
		log.Fatalln("Failed to authenticate user:", err)
	}

	//  Create a Cognito Identity client
	// Create a Security Token Service (STS) client

	myRoleArn := aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023")

	assumeRoleResult, err := s.StsSvc.AssumeRoleWithWebIdentity(context.TODO(), &sts.AssumeRoleWithWebIdentityInput{
		RoleArn:          myRoleArn,
		RoleSessionName:  aws.String("session"),
		WebIdentityToken: authResult.AuthenticationResult.IdToken,
	})

	// if err != nil {
	// 	println(err)
	// 	panic(err)
	// }

	// creds := stscreds.NewAssumeRoleProvider(s.StsSvc, *myRoleArn)
	// webId := stscreds.NewWebIdentityRoleProvider(s.StsSvc, *myRoleArn)

	// _ = webId

	println("Cert")
	println(assumeRoleResult.Credentials)
	// println(creds)

	// s.Cfg.Credentials = aws.NewCredentialsCache(creds)

	// cerds := stscreds.NewWebIdentityRoleProvider(s.StsSvc, *aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023"))
	// s.Cfg.Credentials = aws.NewCredentialsCache(cerds)

	// if err != nil {
	// 	log.Fatalln("Failed to assume role with web identity:", err)
	// }

	// println("Assumrole")
	// println(assumeRoleResult.Credentials)
	// _ = assumeRoleResult.Credentials
	// Set the temporary credentials in the AWS config
	// s.Cfg.Credentials = aws.NewCredentialsCache(assumeRoleResult.Credentials)

	// s.Cfg.Credentials = aws.NewCredentialsCache(assumeRoleResult.Credentials)
	return authResult.AuthenticationResult, nil
}

func GetClientId(idToken string) (string, error) {

	pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
	formattedURL := fmt.Sprintf(pubKeyURL, os.Getenv("AWS_REGION"), os.Getenv("USER_POOL_ID"))
	keySet, err := jwk.Fetch(context.TODO(), formattedURL)
	if err != nil {
		return "", nil
	}
	token, err := jwt.Parse(
		[]byte(idToken),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
	)
	if err != nil {
		return "", nil
	}

	username, _ := token.Get("cognito:username")
	cognitoIdentityId := username.(string)

	return cognitoIdentityId, nil
}
