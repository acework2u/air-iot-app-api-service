package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/config"
	cid "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	_ "github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type airthingService struct {
	Cfg       *aws.Config
	StsSvc    *sts.Client
	IotClient *iot.Client
}

func NewAirThingService(cognitoRegion string) AirThinkService {

	cfg, _ := config.LoadDefaultConfig(context.Background(), config.WithRegion(cognitoRegion), config.WithSharedConfigProfile("default"))
	stsClient := sts.NewFromConfig(cfg)
	iotClient := iot.NewFromConfig(cfg)
	return &airthingService{Cfg: &cfg, StsSvc: stsClient, IotClient: iotClient}
}
func (s *airthingService) GetCerts(idToken string) (interface{}, error) {
	//myRoleArn = *aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023")
	//fmt.Println(idToken)
	//stsSvc := sts.NewFromConfig(*s.Cfg)
	//creds := stscreds.NewAssumeRoleProvider(stsSvc, myRoleArn)
	//credens := aws.NewCredentialsCache(creds)

	//assumeRoleInput := &sts.AssumeRoleInput{
	//	RoleArn:         &myRoleArn,
	//	RoleSessionName: aws.String("sessionIot"),
	//}
	//
	//result, err := TakeRole(context.TODO(), s.StsSvc, assumeRoleInput)

	svs := cid.NewFromConfig(*s.Cfg)
	idRes, err := svs.GetId(context.TODO(), &cid.GetIdInput{
		IdentityPoolId: aws.String("ap-southeast-1:4c5dc3d1-cf9d-4980-8fc8-fdd737f6b84b"),
		Logins: map[string]string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_EqxkPGgmk": *aws.String(idToken),
		},
	})

	if err != nil {
		return nil, err
	}
	cresRes, err := svs.GetCredentialsForIdentity(context.TODO(), &cid.GetCredentialsForIdentityInput{
		IdentityId: idRes.IdentityId,
		Logins: map[string]string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_EqxkPGgmk": *aws.String(idToken),
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("Working inService")

	fmt.Println(cresRes)

	iotEndpoint := "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com"
	_ = iotEndpoint

	return cresRes, nil
}
func (s *airthingService) ThingConnect() error {

	return nil
}
