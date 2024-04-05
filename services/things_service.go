package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	cid "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
	"log"
	"mime/multipart"
	"os"
	"time"
)

var s3client *s3.Client
var Ctx context.Context
var ClientAwsMqtt *mqtt.AWSIoTConnection

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
	CertificateOwnershipToken string `json:"certificateOwnershipToken"`
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
	IotData     *iotdataplane.Client
	StsSvc      *sts.Client
	Cfg         *aws.Config
	S3client    *s3.Client
	Ctx         context.Context
	iotWs       *mqtt.AWSIoTConnection
}

func NewThingClient(cognitoRegion string, userPoolId string, cognitoClientId string) ThinksService {

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
	iotData := iotdataplane.NewFromConfig(cfg)

	return &CogClient{
		AppClientId: cognitoClientId,
		UserPoolId:  userPoolId,
		ClientCog:   cognitoIdentityProviderClient,
		StsSvc:      stsClient,
		IotClient:   iotClient,
		IotData:     iotData,
		Cfg:         &cfg,
		S3client:    s3client,
		Ctx:         context.TODO(),
		iotWs:       ClientAwsMqtt,
	}

}

func (s *CogClient) GetCerts() (interface{}, error) {
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
func (s *CogClient) ThingsConnected(idToken string, things string) (*iotdataplane.PublishOutput, error) {

	client := iotdataplane.NewFromConfig(*s.Cfg)

	//var getThingShadowOutput *iotdataplane.GetThingShadowOutput

	//go func() {
	//	var err error
	//	getThingShadowOutput, err = client.GetThingShadow(context.TODO(), &iotdataplane.GetThingShadowInput{
	//		ThingName:  aws.String(things),
	//		ShadowName: aws.String("air-users"),
	//	})
	//
	//	if err != nil {
	//		fmt.Println("Shadow Error")
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println("Shadow")
	//	rep := map[string]interface{}{}
	//	_ = json.Unmarshal([]byte(getThingShadowOutput.Payload), &rep)
	//	fmt.Println(rep["metadata"])
	//
	//}()

	payload := &AirPayload{
		Message: idToken,
	}
	bytes, _ := json.Marshal(payload)
	pubTopic := fmt.Sprintf("%v/CD/%v", things, things)
	fmt.Println("pubTopic")
	fmt.Println(pubTopic)

	publishOutput, err := client.Publish(context.TODO(), &iotdataplane.PublishInput{
		Topic:   aws.String(pubTopic),
		Payload: bytes,
	})

	if err != nil {
		return nil, err
	}
	return publishOutput, err

	////Shadow Sub
	//getThingShadowOutput, err := client.GetThingShadow(context.TODO(), &iotdataplane.GetThingShadowInput{
	//	ThingName:  aws.String("2300F15050017"),
	//	ShadowName: aws.String("air-users"),
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//fmt.Println("getThingShadowOutput")
	//fmt.Println(getThingShadowOutput)
	//
	//_ = getThingShadowOutput
	//
	//
	//
	//fmt.Println("publishOutput")
	//fmt.Println(publishOutput)
	//
	//if err != nil {
	//	fmt.Println("Service Error")
	//	fmt.Println(err)
	//}
	////return getThingShadowOutput, nil
	//return publishOutput, nil

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
func (s *CogClient) ThinksShadows(idToken string, res string) (*ShadowsValue, error) {

	shadowsVal := &ShadowsValue{}

	err := s.NewAwsMqttConnect(idToken)
	if err != nil {
		return nil, err
	}
	//fmt.Println("AWS IoT Mqtt connection")
	shadowsAcceptTopic := fmt.Sprintf("$aws/things/%s/shadow/name/air-users/update/accepted", res)
	fmt.Println(shadowsAcceptTopic)

	revMsg := make(chan *ShadowsAccepted)
	deviceValShadows := ShadowsAccepted{}

	_ = revMsg
	_ = deviceValShadows

	go func(msg chan<- *ShadowsAccepted) {
		//	fmt.Println("Work in Go Routine")
		ok := s.iotWs.SubscribeWithHandler(shadowsAcceptTopic, 0, func(client MQTT.Client, message MQTT.Message) {
			msgPayload := fmt.Sprintf(`%v`, string(message.Payload()))
			fmt.Println("Work in Handle")
			fmt.Println(msgPayload)
			shadowDoc := &ShadowsAccepted{}
			err := json.Unmarshal(message.Payload(), shadowDoc)
			if err != nil {
				panic(err)
			}
			revMsg <- shadowDoc

		})
		if ok != nil {
			fmt.Println("Error Sub")
			fmt.Println(ok.Error())
		}

	}(revMsg)

	select {
	case <-revMsg:

		mesg := <-revMsg

		msg := mesg.State.Reported.Message
		decodeShadow, err := utils.GetClaimsFromToken(msg)
		if err != nil {
			return nil, err
		}

		acValReq := utils.DecodeValAcShadow(decodeShadow)

		acVal := utils.NewGetAcVal(acValReq)
		ac1000 := acVal.Ac1000()

		//acData := (*IndoorInfo)(ac1000)
		acData := &IndoorInfo{
			Power: ac1000.Power,
		}

		log.Println(acData)
		log.SetFlags(log.Ldate | log.Lshortfile)

	default:
		fmt.Println("no chan value")
		//log.Panicln("no chan val")

	}
	return shadowsVal, nil
}
func (s *CogClient) PubGetShadows(thinkName string, shadowName string) (*IndoorInfo, error) {

	subTopic := &iotdataplane.GetThingShadowInput{
		ThingName:  aws.String(thinkName),
		ShadowName: aws.String("air-users"),
	}
	getThingShadowOutput, err := s.IotData.GetThingShadow(s.Ctx, subTopic)
	if err != nil {
		return nil, err
	}
	shadowVal := &ShadowsValue{}
	err = json.Unmarshal(getThingShadowOutput.Payload, shadowVal)

	if err != nil {
		return nil, err
	}

	decodeShadow, err := utils.GetClaimsFromToken(shadowVal.State.Reported.Message)
	if err != nil {
		return nil, err
	}
	acValReq := utils.DecodeValAcShadow(decodeShadow)
	acVal := utils.NewGetAcVal(acValReq)
	ac1000 := acVal.Ac1000()

	pubTopic := fmt.Sprintf("$aws/things/%v/shadow/name/air-users/get", thinkName)
	_, err = s.IotData.Publish(s.Ctx, &iotdataplane.PublishInput{Topic: aws.String(pubTopic)})
	if err != nil {
		return nil, err
	}

	//acData := (*IndoorInfo)(ac1000)
	acData := &IndoorInfo{
		Power:    ac1000.Power,
		Mode:     ac1000.Mode,
		Temp:     ac1000.Temp,
		RoomTemp: ac1000.RoomTemp,
		RhSet:    ac1000.RhSet,
		RhRoom:   ac1000.RhRoom,
		FanSpeed: ac1000.FanSpeed,
		Louver:   ac1000.Louver,
		Aps:      ac1000.Aps,
		OzoneGen: ac1000.OzoneGen,
		Pm25Info: (Pm25Info)(ac1000.Pm25Info),
	}

	return acData, nil

}
func (s *CogClient) PubUpdateShadows(thinkName string, payload string) (*IndoorInfo, error) {

	shadowsCmd := &ShadowsCommand{}
	shadowsCmd.State.Desired.Cmd = payload
	shadowsPayload, err := json.Marshal(shadowsCmd)
	if err != nil {
		return nil, err
	}

	_, err = s.IotData.UpdateThingShadow(s.Ctx, &iotdataplane.UpdateThingShadowInput{
		Payload:    shadowsPayload,
		ShadowName: aws.String("air-users"),
		ThingName:  aws.String(thinkName),
	})
	if err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)

	shadowOutput, ok := s.PubGetShadows(thinkName, "")
	if ok != nil {
		return nil, ok
	}

	return shadowOutput, nil
}
func (s *CogClient) NewAwsMqttConnect(cognitoIdentityId string) error {
	var err error
	clientMq, err := mqtt.NewConnection(mqtt.Config{
		KeyPath:  "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-private.pem.key",
		CertPath: "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-certificate.pem.crt",
		CAPath:   "./certs/cert7/AmazonRootCA1.pem",
		ClientId: *aws.String(cognitoIdentityId),
		Endpoint: "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com",
	})

	if err != nil {
		return err
	}

	s.iotWs = clientMq

	return err
}

func iotSub(topic string, result chan<- *ShadowsValue) {
	shadowsVal := &ShadowsValue{}

	go func() {
		ClientAwsMqtt.SubscribeWithHandler(topic, 0, func(client MQTT.Client, message MQTT.Message) {
			//msgPayload := fmt.Sprintf(`%v`, string(message.Payload()))
			err := json.Unmarshal(message.Payload(), &shadowsVal)
			if err != nil {
				fmt.Println("err")
				fmt.Println(err)
			}
			//fmt.Println("shadowsVal")
			//fmt.Println(shadowsVal)
			result <- shadowsVal
		})
	}()

}

func NewAwsMqttConnect(cognitoIdentityId string) (*mqtt.AWSIoTConnection, error) {
	var err error
	clientMq, err := mqtt.NewConnection(mqtt.Config{
		KeyPath:  "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-private.pem.key",
		CertPath: "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-certificate.pem.crt",
		CAPath:   "./certs/cert7/AmazonRootCA1.pem",
		ClientId: *aws.String(cognitoIdentityId),
		Endpoint: "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com",
	})

	if err != nil {
		return nil, err
	}

	return clientMq, err
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
