package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	cid "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
)

var s3client *s3.Client
var Ctx context.Context

type AirCmdToAws struct {
	SerialNumber string     `json:"serialNumber"`
	Data         AirCommand `json:"data"`
}
type AirCommand struct {
	Cmd string `json:"cmd"`
}

type AirPayload struct {
	Message string `json:"message"`
}

type STSAssumeRoleAPI interface {
	AssumeRole(ctx context.Context,
		params *sts.AssumeRoleInput,
		optFns ...func(*sts.Options)) (*sts.AssumeRoleOutput, error)
}

func TakeRole(c context.Context, api STSAssumeRoleAPI, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return api.AssumeRole(c, input)
}

// IotCore

type awsIotCert struct {
	CertificateId             string `json:"certificateId"`
	CertificatePem            string `json:"certificatePem"`
	PrivateKey                string `json:"privateKey"`
	CertificateOwnershipToken string `josn:"certificateOwnershipToken"`
}

type deviceRegister struct {
	CertificateOwnershipToken string      `json:"certificateOwnershipToken"`
	Parameters                deviceParam `json:"parameters"`
}
type deviceParam struct {
	SerialNumber        string `json:"SerialNumber"`
	AWSIoTCertificateId string `json:"AWS::IoT::Certificate::Id"`
}

var certIot awsIotCert
var regPlayload deviceRegister
var myRoleArn = *aws.String("arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role")

// Cognito

type CogClient struct {
	AppClientId string
	UserPoolId  string
	ClientCog   *cip.Client
	IotClient   *iot.Client
	StsSvc      *sts.Client
	Cfg         *aws.Config
	S3client    *s3.Client
	Ctx         context.Context
}

func NewThingClient(cognitoRegion string, userPoolId string, cognitoClientId string) ThinksService {

	// customResolver := aws.EndpointResolverWithOptions(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	if service == sts.ServiceID && region == "us-west-2" {
	// 		return aws.Endpoint{
	// 			PartitionID:   "aws",
	// 			URL:           "https://sts.us-west-2.amazonaws.com",
	// 			SigningRegion: "us-west-2",
	// 		}, nil
	// 	}
	// 	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	// })

	// customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	if service == sts.ServiceID && region == *aws.String("us-west-2") {
	// 		return aws.Endpoint{
	// 			PartitionID:   *aws.String("aws"),
	// 			URL:           *aws.String("https://sts.us-west-2.amazonaws.com"),
	// 			SigningRegion: *aws.String("us-west-2"),
	// 		}, nil
	// 	}
	// 	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	// })

	// fmt.Println("<-------customResolver---------->")
	// fmt.Println(customResolver)

	awsEndpoint := "https://s3.ap-southeast-1.amazonaws.com"
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: *aws.String("ap-southeast-1"),
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	_ = customResolver
	//customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
	//	if awsEndpoint != "" {
	//		return aws.Endpoint{
	//			PartitionID:   "aws",
	//			URL:           awsEndpoint,
	//			SigningRegion: awsRegion,
	//		}, nil
	//	}
	//
	//	// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
	//	return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	//})

	// cfg, err := config.LoadDefaultConfig(context.Background(), config.WithEndpointResolverWithOptions(customResolver), config.WithClientLogMode(1))
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion), config.WithSharedConfigProfile("default"))
	//cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("ap-southeast-1"), config.WithEndpointResolverWithOptions(customResolver))

	// cfg, err := config.LoadDefaultConfig(context.Background(),config.With)
	if err != nil {
		// log.Fatalln("Failed to load AWS config:", err)
		// panic(fmt.Sprintf("Error configuring AWS: %s", err))
	}

	//assumeCnf, _ := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion))

	cognitoIdentityProviderClient := cip.NewFromConfig(cfg)
	stsClient := sts.NewFromConfig(cfg)
	iotClient := iot.NewFromConfig(cfg)
	s3client = s3.NewFromConfig(cfg)

	return &CogClient{
		AppClientId: cognitoClientId,
		UserPoolId:  userPoolId,
		ClientCog:   cognitoIdentityProviderClient,
		StsSvc:      stsClient,
		IotClient:   iotClient,
		Cfg:         &cfg,
		S3client:    s3client,
		Ctx:         context.TODO(),
	}

}

func (s *CogClient) GetCerds() (interface{}, error) {
	username := "wowoy73603@camplvad.com"
	password := "J@e2262527"

	// _ = creds
	authResult, err := s.ClientCog.InitiateAuth(context.TODO(), &cip.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		AuthParameters: map[string]string{
			"USERNAME": *aws.String(username),
			"PASSWORD": *aws.String(password),
		},
		ClientId: aws.String(s.AppClientId),
	})
	if err != nil {
		log.Fatalln("Failed to authenticate user:", err)
	}
	fmt.Println("&authResult.AuthenticationResult.IdToken")

	svs := cid.NewFromConfig(*s.Cfg)

	fmt.Println(svs)

	idRes, err := svs.GetId(context.TODO(), &cid.GetIdInput{
		IdentityPoolId: aws.String("ap-southeast-1:8f60452e-9616-4914-bdbf-d8f149ca8dfa"),
		Logins: map[string]string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_yW7AZdShx": *authResult.AuthenticationResult.IdToken,
		},
	})

	if err != nil {
		fmt.Println(err.Error())

		fmt.Println("This Error Block")
	}

	fmt.Println("<--- idRes --->")
	// fmt.Println(idRes)

	cresRes, err := svs.GetCredentialsForIdentity(context.TODO(), &cid.GetCredentialsForIdentityInput{
		IdentityId: idRes.IdentityId,
		Logins: map[string]string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_yW7AZdShx": *authResult.AuthenticationResult.IdToken,
		},
	})

	if err != nil {
		fmt.Println("cresRes Error")
		fmt.Println(err.Error())

	}

	// iotClient := iot.NewFromConfig(*s.Cfg)

	// certList, err := iotClient.GetPolicy(context.TODO(), &iot.GetPolicyInput{
	// 	PolicyName: aws.String("AirThingPolicy"),
	// })

	certList, err := s.IotClient.ListAttachedPolicies(context.TODO(), &iot.ListAttachedPoliciesInput{
		Target: aws.String("arn:aws:iot:ap-southeast-1:513310385702:cert/ffe2384c236d4b639c830b18d578f8a35f97eac3c8b88b6f420d795428b9ab85"),
	})

	if err != nil {
		fmt.Println("Error IOT")
		fmt.Println(err)
	}

	fmt.Println("IoT Cert List")
	fmt.Println(certList)
	_ = certList
	// fmt.Println(certList)

	// myArn := "arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role"
	// client := sts.NewFromConfig(*s.Cfg)
	// newCreds := stscreds.NewAssumeRoleProvider(client, myArn)

	// _ = newCreds

	// newCreds := credentials.NewStaticCredentialsProvider(*cresRes.Credentials.AccessKeyId, *cresRes.Credentials.SecretKey, *cresRes.Credentials.SessionToken)

	/*


		_ = cresRes

		myArn := "arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role"
		client := sts.NewFromConfig(*s.Cfg)

		creds := stscreds.NewAssumeRoleProvider(client, myArn)

		fmt.Println("stscreds")
		fmt.Println(creds)

		aws.NewCredentialsCache(creds)

		Credentials := aws.NewCredentialsCache(creds)

	*/

	// input := &sts.AssumeRoleInput{
	// 	RoleArn:         aws.String("arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role"),
	// 	RoleSessionName: aws.String("sessionIot"),
	// }

	// result, err := TakeRole(context.TODO(), client, input)

	// if err != nil {
	// 	fmt.Println("Error assuming the role")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(result.AssumedRoleUser)

	// fmt.Println(cresRes)

	// return cresRes, nil

	_ = authResult
	_ = cresRes

	return cresRes, nil

}
func (s *CogClient) GetUserCert(user *UserReq) (interface{}, error) {

	fmt.Println("User Login")
	fmt.Println(user)

	fmt.Println("Working in Service")

	return user, nil

}
func (s *CogClient) UploadToS3(file *multipart.FileHeader) (interface{}, error) {

	fmt.Println("Working in Service")
	fmt.Printf("type of c is %T\n", file)

	f, openErr := file.Open()
	if openErr != nil {
		return "", openErr
	}
	uploader := manager.NewUploader(s.S3client)

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("airiotbucket"),
		Key:    aws.String("image/" + file.Filename),
		Body:   f,
		//ACL:    "public-read",
	})

	if uploadErr != nil {
		return "", uploadErr
	}

	return result, nil
}
func (s *CogClient) ThingRegister(idToken string) (interface{}, error) {

	cognitoIdentityId := "ap-southeast-1:4c5dc3d1-cf9d-4980-8fc8-fdd737f6b84b"
	//#AttachIotPolicyToIdentity
	attachPolicyOutput, err := s.IotClient.AttachPolicy(context.TODO(), &iot.AttachPolicyInput{
		PolicyName: aws.String("AirThingPolicy"),
		Target:     aws.String(cognitoIdentityId),
	})

	fmt.Println("attachPolicyOutput")
	fmt.Println(attachPolicyOutput)
	//
	//s.IotClient.UpdateThing()

	//fmt.Println("ClientID")
	//fmt.Println(clientId)

	//attachThingPrincipalOutput, err := s.IotClient.AttachThingPrincipal(context.TODO(), &iot.AttachThingPrincipalInput{
	//	Principal: aws.String("arn:aws:cognito-identity:ap-southeast-1:513310385702:identitypool/ap-southeast-1:4c5dc3d1-cf9d-4980-8fc8-fdd737f6b84b"),
	//	ThingName: aws.String("23F05110000126"),
	//})

	//fmt.Println("attachPolicyOutput")
	//fmt.Println(attachPolicyOutput)

	/*
		connection, err := mqtt.NewConnection(mqtt.Config{
			KeyPath:  "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-private.pem.key",
			CertPath: "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-certificate.pem.crt",
			CAPath:   "./certs/cert7/AmazonRootCA1.pem",
			ClientId: *aws.String(cognitoIdentityId),
			Endpoint: "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com",
		})
	*/
	if err != nil {

		fmt.Println(err)
		panic(err)

	}

	//_ = connection

	return attachPolicyOutput, nil
}
func (s *CogClient) ThingsConnected(idToken string) (*iotdataplane.PublishOutput, error) {

	client := iotdataplane.NewFromConfig(*s.Cfg)

	//getThingShadowOutput, err := client.GetThingShadow(context.TODO(), &iotdataplane.GetThingShadowInput{
	//	ThingName:  aws.String("23F05110000126"),
	//	ShadowName: aws.String("air-users"),
	//})
	//AirCon := map[string]int{
	//	"setTemp":  55,
	//	"roomTemp": 50,
	//	"co2":      45,
	//}
	//bytes, _ := json.Marshal(AirCon)

	airPayload := make([]byte, 40)
	copy(airPayload[0:], string(1))
	copy(airPayload[1:], string(3))
	copy(airPayload[2:], string(60))
	copy(airPayload[3:], string(120))
	copy(airPayload[4:], string(1))
	copy(airPayload[14:], string(1))

	reg2000 := make([]byte, 40)
	reg3000 := make([]byte, 160)
	reg4000 := make([]byte, 40)

	airCon := map[string][]byte{
		"reg1000": airPayload,
		"reg2000": reg2000,
		"reg3000": reg3000,
		"reg4000": reg4000,
	}
	_ = airCon
	cmd := "on"

	//dataFrame := utils.AirPower(cmd)
	dataFrame2 := utils.AirPower2(cmd)

	payload := &AirCmdToAws{
		SerialNumber: string("2300F15050017"),
		Data:         AirCommand{Cmd: fmt.Sprintf("%x", dataFrame2)},
	}

	fmt.Println("Data Frame")
	fmt.Println(dataFrame2)

	bytes, _ := json.Marshal(payload)

	var publishOutput *iotdataplane.PublishOutput

	publishOutput, err := client.Publish(context.TODO(), &iotdataplane.PublishInput{
		Topic:   aws.String("2300F15050017/CD/2300F15050017"),
		Payload: bytes,
	})

	fmt.Println("publishOutput")
	fmt.Println(publishOutput)

	if err != nil {
		fmt.Println("Service Error")
		fmt.Println(err)
	}
	//return getThingShadowOutput, nil
	return publishOutput, nil

}
func (s *CogClient) ThingsCert(idToken string) (interface{}, error) {

	cogClient := cid.NewFromConfig(*s.Cfg)
	IdToken := &idToken
	idUser, err := cogClient.GetId(context.TODO(), &cid.GetIdInput{
		IdentityPoolId: aws.String("ap-southeast-1:4c5dc3d1-cf9d-4980-8fc8-fdd737f6b84b"),
		Logins: map[string]string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_EqxkPGgmk": *IdToken,
		},
	})
	if err != nil {
		return nil, err
	}

	// Get Cert AssumeRole
	myRoleArn = *aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023")
	assumeRoleOutput, err := s.StsSvc.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
		RoleArn:         &myRoleArn,
		RoleSessionName: aws.String("cogCert"),
	})

	// Get CredentialForIdentity
	svs := cid.NewFromConfig(*s.Cfg)
	cresRes, err := svs.GetCredentialsForIdentity(context.TODO(), &cid.GetCredentialsForIdentityInput{
		IdentityId: idUser.IdentityId,
		Logins: map[string]string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_EqxkPGgmk": *IdToken,
		},
	})

	if err != nil {
		return nil, err
	}

	credsNew := stscreds.NewAssumeRoleProvider(s.StsSvc, myRoleArn)
	certNew := aws.NewCredentialsCache(credsNew)

	airCon := map[string]interface{}{
		"CredentIden": cresRes,
		"CertAssume":  assumeRoleOutput.Credentials,
		"CertNew":     certNew,
	}

	return airCon, nil
}

// func (s *CogClient) GetCerds() (interface{}, error) {

// 	username := "wowoy73603@camplvad.com"
// 	password := "J@e2262527"

// 	//Ser

// 	// stsSvc := sts.NewFromConfig(*s.Cfg)
// 	// creds := stscreds.NewAssumeRoleProvider(stsSvc, myRoleArn)

// 	// credens := aws.NewCredentialsCache(creds)

// 	fmt.Println("<------credens---------->")
// 	// fmt.Println(credens.Retrieve(context.TODO()))

// 	// _ = creds
// 	authResult, err := s.ClientCog.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
// 		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
// 		AuthParameters: map[string]string{
// 			"USERNAME": *aws.String(username),
// 			"PASSWORD": *aws.String(password),
// 		},
// 		ClientId: aws.String(s.AppClientId),
// 	})
// 	if err != nil {
// 		log.Fatalln("Failed to authenticate user:", err)
// 	}
// 	fmt.Println("&authResult.AuthenticationResult.IdToken")

// 	fmt.Println("<---iotSession----<")

// 	// 	fmt.Println(webId)
// 	// webId := stscreds.NewWebIdentityRoleProvider(s.StsSvc, *myRoleArn)
// 	// fmt.Println("Web Idd")
// 	// fmt.Println(webId)

// 	// Create assomeRole
// 	// asSumeRoleInput := sts.AssumeRoleInput{
// 	// 	RoleArn:         &myRoleArn,
// 	// 	RoleSessionName: aws.String("session"),
// 	// 	TokenCode:       authResult.AuthenticationResult.IdToken,
// 	// }

// 	// credsOut, err := stsSvc.AssumeRole(context.TODO(), &asSumeRoleInput, func(o *sts.Options) {
// 	// 	o.Region = *aws.String("us-east-1")
// 	// })

// 	// if err != nil {
// 	// 	log.Fatal("can not assomeRole :", err)
// 	// }

// 	// fmt.Println(credsOut)

// 	// Create a Cognito Identity client
// 	// cognitoIdentityClient := cognitoidentity.NewFromConfig(*s.Cfg)

// 	// // // arn:aws:cognito-idp:ap-southeast-1:513310385702:userpool/ap-southeast-1_yW7AZdShx

// 	// getIdentityResult, err := cognitoIdentityClient.GetId(context.TODO(), &cognitoidentity.GetIdInput{
// 	// 	IdentityPoolId: aws.String(s.UserPoolId),
// 	// 	Logins: map[string]string{
// 	// 		"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_yW7AZdShx": *authResult.AuthenticationResult.IdToken,
// 	// 	},
// 	// })

// 	pubKeyURL := "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"
// 	formattedURL := fmt.Sprintf(pubKeyURL, os.Getenv("AWS_REGION"), s.UserPoolId)
// 	keySet, err := jwk.Fetch(context.TODO(), formattedURL)
// 	if err != nil {
// 		return "", nil
// 	}

// 	idToken := authResult.AuthenticationResult.IdToken

// 	token, err := jwt.Parse(
// 		[]byte(*idToken),
// 		jwt.WithKeySet(keySet),
// 		jwt.WithValidate(true),
// 	)
// 	if err != nil {
// 		return "", nil
// 	}

// 	if err != nil {
// 		fmt.Println("Failed to get Cognito identity ID:")
// 		fmt.Println(err.Error())
// 		log.Fatalln("Failed to get Cognito identity ID:", err)
// 	}

// 	userInfo, _ := token.Get("cognito:username")
// 	cognitoIdentityId := userInfo.(string)

// 	fmt.Println("cognitoIdentityId")
// 	fmt.Println(cognitoIdentityId)

// 	// myRoleArn := aws.String("arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role")
// 	// assomeRoleResult, err := s.StsSvc.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
// 	// 	RoleArn:         myRoleArn,
// 	// 	RoleSessionName: aws.String("session"),
// 	// 	TokenCode:       authResult.AuthenticationResult.IdToken,
// 	// })

// 	// if err != nil {
// 	// 	fmt.Println("Can not Assumrole")
// 	// 	fmt.Println(err.Error())
// 	// }

// 	// fmt.Println(assomeRoleResult)

// 	// stsClient := sts.NewFromConfig(*s.Cfg)

// 	// // Assume an IAM role with the Cognito identity as the role's principal
// 	// myRoleArn := aws.String("arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role")

// 	// assumeRoleResult, err := s.StsSvc.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
// 	// 	RoleArn:         &myRoleArn,
// 	// 	RoleSessionName: aws.String("iot-access-air"),
// 	// 	DurationSeconds: aws.Int32(900),
// 	// })

// 	/*
// 	assumeRoleResult, err := s.StsSvc.AssumeRoleWithWebIdentity(context.TODO(), &sts.AssumeRoleWithWebIdentityInput{
// 		RoleArn:          &myRoleArn,
// 		RoleSessionName:  aws.String("session"),
// 		WebIdentityToken: authResult.AuthenticationResult.IdToken,
// 	})

// 	if err != nil {
// 		// log.Fatalln("Failed to assume role with web identity:", err)
// 		fmt.Println(err)
// 	}

// 	fmt.Println(assumeRoleResult)
// */

// 	svs := cid.NewFromConfig(*s.Cfg)

// 	idRes,_ := svs.GetId(context.TODO(),&cid.GetIdInput{
// 		IdentityPoolId: aws.String("ap-southeast-1:8f60452e-9616-4914-bdbf-d8f149ca8dfa"),
// 		Logins: map[string]string{
// 			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_yW7AZdShx:":assumeRoleResult.AuthenticationResult.IdToken
// 		},
// 	})

// 	fmt.Println("<--- idRes --->")
// 	// fmt.Println(idRes)

// 	// Connect Iot Core

// 	// s.Cfg.Credentials = aws.NewCredentialsCache(assumeRoleResult)

// 	return authResult, nil
// }

// func (s *CogClient) GetCerds() (interface{}, error) {

// 	// username := "tidosi6511@vaband.com"
// 	// password := "Pass@word2020"
// 	username := "wowoy73603@camplvad.com"
// 	password := "J@e2262527"

// 	// Authenticate the user and retrieve the Cognito ID token
// 	authResult, err := s.ClientCog.InitiateAuth(context.TODO(), &cip.InitiateAuthInput{
// 		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
// 		AuthParameters: map[string]string{
// 			"USERNAME": username,
// 			"PASSWORD": password,
// 		},
// 		ClientId: aws.String(s.AppClientId),
// 		// UserPoolId: aws.String(s.UserPoolId),
// 	})

// 	// authResult, err := s.ClientCog.AdminInitiateAuth(context.TODO(), &cip.AdminInitiateAuthInput{
// 	// 	AuthFlow: types.AuthFlowTypeAdminNoSrpAuth,
// 	// 	AuthParameters: map[string]string{
// 	// 		"USERNAME": username,
// 	// 		"PASSWORD": password,
// 	// 	},
// 	// 	ClientId:   aws.String(s.AppClientId),
// 	// 	UserPoolId: aws.String(s.UserPoolId),
// 	// })

// 	if err != nil {
// 		log.Fatalln("Failed to authenticate user:", err)
// 	}

// 	//  Create a Cognito Identity client
// 	// Create a Security Token Service (STS) client

// 	myRoleArn := aws.String("arn:aws:iam::513310385702:role/Cognito_aws_iotUnauth_Role")
// 	// myRoleArn := aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023")

// 	// assumeRoleResult, err := s.StsSvc.AssumeRoleWithWebIdentity(context.TODO(), &sts.AssumeRoleWithWebIdentityInput{
// 	// 	RoleArn:          myRoleArn,
// 	// 	RoleSessionName:  aws.String("session"),
// 	// 	WebIdentityToken: authResult.AuthenticationResult.IdToken,
// 	// })

// 	// if err != nil {
// 	// 	println(err)
// 	// 	panic(err)
// 	// }

// 	// _ = assumeRoleResult

// 	creds := stscreds.NewAssumeRoleProvider(s.StsSvc, *myRoleArn)

// 	idenInputParams := sts.AssumeRoleWithWebIdentityInput{
// 		RoleArn:          myRoleArn,
// 		RoleSessionName:  aws.String("session"),
// 		WebIdentityToken: authResult.AuthenticationResult.IdToken,
// 	}

// 	webId, err := s.StsSvc.AssumeRoleWithWebIdentity(context.TODO(), &idenInputParams)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	println(creds)
// 	fmt.Println(webId)
// 	// webId := stscreds.NewWebIdentityRoleProvider(s.StsSvc, *myRoleArn)

// 	// _ = webId
// 	// _ = myRoleArn
// 	// println("Cert")
// 	// fmt.Println(assumeRoleResult)
// 	// attParams := &iot.AttachPolicyInput{
// 	// 	PolicyName: aws.String("Cognito_aws_iotUnauth_Role"),
// 	// 	Target:     aws.String("arn:aws:iot:ap-southeast-1:513310385702:thinggroup/air_iot"),
// 	// }
// 	// attachPolicyOutput, err := s.IotClient.AttachPolicy(context.TODO(), attParams)

// 	// principalParams := iot.AttachPrincipalPolicyInput{
// 	// 	PolicyName: aws.String("AirThingPolicy"),
// 	// 	Principal:  aws.String("ap-southeast-1:8f60452e-9616-4914-bdbf-d8f149ca8dfa"),
// 	// }

// 	// _ = principalParams
// 	// principalOut, err := s.IotClient.AttachPrincipalPolicy(context.TODO(), &principalParams)

// 	// if err != nil {
// 	// 	fmt.Println("Not attach work")
// 	// 	fmt.Println(err.Error())
// 	// }

// 	// fmt.Println(principalOut)

// 	// if err != nil {
// 	// 	println("AttachPolicy not work")
// 	// 	panic(err)

// 	// }

// 	// println("attachPolicyOutput")
// 	// println(principalOut)
// 	// signer := v4.NewSigner(assumeRoleResult.Credentials)
// 	// _ = signer
// 	// println(assumeRoleResult.Credentials)
// 	// println(creds)

// 	// s.Cfg.Credentials = aws.NewCredentialsCache(creds)

// 	// cerds := stscreds.NewWebIdentityRoleProvider(s.StsSvc, *aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023"))
// 	// s.Cfg.Credentials = aws.NewCredentialsCache(cerds)

// 	// if err != nil {
// 	// 	log.Fatalln("Failed to assume role with web identity:", err)
// 	// }

// 	// println("Assumrole")
// 	// println(assumeRoleResult.Credentials)
// 	// _ = assumeRoleResult.Credentials
// 	// Set the temporary credentials in the AWS config
// 	// s.Cfg.Credentials = aws.NewCredentialsCache(assumeRoleResult.Credentials)

// 	// s.Cfg.Credentials = aws.NewCredentialsCache(assumeRoleResult.Credentials)
// 	return authResult.AuthenticationResult, nil
// }

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

func iotConn(cognitoIdentityId string) {
	connection, err := mqtt.NewConnection(mqtt.Config{
		KeyPath:  "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-private.pem.key",
		CertPath: "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-certificate.pem.crt",
		CAPath:   "./certs/cert7/AmazonRootCA1.pem",
		ClientId: *aws.String(cognitoIdentityId),
		Endpoint: "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com",
	})

	if err != nil {
		panic(err)
	}
	go func() {
		err = connection.SubscribeWithHandler("$aws/certificates/create/json/accepted", 0, func(client MQTT.Client, message MQTT.Message) {
			//print(string(message.Payload()))
			fmt.Println("<!-----Certificate Create Accepted--->")
			msgPayload := fmt.Sprintf(`%v`, string(message.Payload()))
			//fmt.Println(msgPayload)

			ok := json.Unmarshal([]byte(msgPayload), &certIot)

			if ok != nil {
				fmt.Println(err.Error())
				//json: Unmarshal(non-pointer main.Request)
			} else {

				pubPayload := deviceRegister{
					CertificateOwnershipToken: certIot.CertificateOwnershipToken,
					Parameters: deviceParam{
						SerialNumber:        "23F05110000126",
						AWSIoTCertificateId: string(certIot.CertificateId),
					},
				}

				fmt.Println(certIot.CertificateId)
				data, _ := json.Marshal(pubPayload)

				fmt.Println(string(data))

				regDev := connection.Publish("$aws/provisioning-templates/AirIotProvisionTemplate/provision/json", data, 0)
				if regDev != nil {
					fmt.Println(err.Error())
				}

			}

		})

	}()
	if err != nil {
		panic(err)
	}
	go func() {
		err = connection.SubscribeWithHandler("$aws/provisioning-templates/AirIotProvisionTemplate/provision/json/accepted", 0, func(client MQTT.Client, message MQTT.Message) {
			fmt.Println("<!-----Provision Acceped--->")
			print(string(message.Payload()))

		})

	}()

	go func() {
		err = connection.SubscribeWithHandler("$aws/provisioning-templates/AirIotProvisionTemplate/provision/json/rejected", 0, func(client MQTT.Client, message MQTT.Message) {
			fmt.Println("<!-----Provision Rejected--->")
			print(string(message.Payload()))

		})

	}()
	if err != nil {
		panic(err)
	}

	err = connection.Publish("$aws/certificates/create/json", "", 0)
	if err != nil {
		panic(err)
	}

	for {

		fmt.Printf("IOT AC =%v \n", time.Now())
		time.Sleep(4 * time.Second)
	}
} // end of func
