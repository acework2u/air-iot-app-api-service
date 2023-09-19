package services

import (
	"context"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	_ "github.com/aws/aws-sdk-go-v2/config"
	cid "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	_ "github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"time"
)

type airthingService struct {
	Cfg             *aws.Config
	StsSvc          *sts.Client
	IotClient       *iot.Client
	airRepo         repository.AirRepository
	userPoolId      string
	cognitoClientId string
	region          string
}

func NewAirThingService(airconfig *AirThingConfig, airRepo repository.AirRepository) AirThinkService {

	cfg, _ := config.LoadDefaultConfig(context.Background(), config.WithRegion(airconfig.Region), config.WithSharedConfigProfile("default"))
	stsClient := sts.NewFromConfig(cfg)
	iotClient := iot.NewFromConfig(cfg)
	return &airthingService{Cfg: &cfg, StsSvc: stsClient, IotClient: iotClient, airRepo: airRepo, userPoolId: airconfig.UserPoolId, cognitoClientId: airconfig.CognitoClientId, region: airconfig.Region}
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

	iotEndpoint := "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com"
	_ = iotEndpoint

	return cresRes, nil
}
func (s *airthingService) ThingConnect(idToken string) (interface{}, error) {

	myRoleArn = *aws.String("arn:aws:iam::513310385702:role/service-role/customer_air_iot_2023")
	assumeRoleOutput, err := s.StsSvc.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
		RoleArn:         &myRoleArn,
		RoleSessionName: aws.String("cogCert"),
	})

	if err != nil {
		return nil, err
	}
	fmt.Println("assumeRoleOutput")

	return assumeRoleOutput, nil
}
func (s *airthingService) AddAir(info *AirInfo) (*DBAirInfo, error) {

	now := time.Now()
	airInfo := &repository.AirInfo{
		Serial:       info.Serial,
		UserId:       info.UserId,
		Title:        info.Title,
		Bg:           info.Bg,
		RegisterDate: now.Local(),
		UpdatedDate:  now.Local(),
	}

	res, err := s.airRepo.RegisterAir(airInfo)

	if err != nil {
		return nil, err
	}
	//newAirRegInfo := (*DBAirInfo)(res)
	newAirRegInfo := &DBAirInfo{
		Id:      res.Id,
		Serial:  res.Serial,
		Title:   res.Title,
		Bg:      res.Bg,
		Widgets: (AirWidget)(res.Widgets),
		Status:  res.Status,
	}

	return newAirRegInfo, nil
}
func (s *airthingService) GetAirs(userId string) ([]*ResponseAir, error) {

	var airList []*ResponseAir
	res, err := s.airRepo.Airs(userId)
	if err != nil {
		return nil, err
	}

	thingsServ := NewThingClient(s.region, s.userPoolId, s.cognitoClientId)
	_ = thingsServ

	for _, items := range res {

		indInfo := &IndoorInfo{}
		result := make(chan *IndoorInfo)

		//indInfo, _ = thingsServ.PubGetShadows(items.Serial, "")
		go func(thingsName string, result chan<- *IndoorInfo) {
			res, _ := thingsServ.PubGetShadows(thingsName, "")
			result <- res
		}(items.Serial, result)

		indInfo = <-result

		item := &ResponseAir{
			Id:      items.Id,
			Serial:  items.Serial,
			Title:   items.Title,
			Indoor:  indInfo,
			Bg:      items.Bg,
			Widgets: (AirWidget)(items.Widgets),
		}
		airList = append(airList, item)
	}
	return airList, nil
}
func (s *airthingService) UpdateAir(filter *FilterUpdate, info *UpdateAirInfo) (*DBAirInfo, error) {
	fmt.Println(filter)
	fmt.Println(info)

	fil := (*repository.FilterUpdate)(filter)
	updateInfo := &repository.UpdateAirInfo{
		Serial:      info.Serial,
		UserId:      info.UserId,
		Title:       info.Title,
		Bg:          info.Bg,
		Widgets:     (repository.AirWidget)(info.Widgets),
		UpdatedDate: time.Now(),
	}
	airInfo, err := s.airRepo.UpdateAir(fil, updateInfo)
	if err != nil {
		return nil, err
	}

	resAirInfo := &DBAirInfo{
		Id:          airInfo.Id,
		Serial:      airInfo.Serial,
		Title:       airInfo.Title,
		Status:      airInfo.Status,
		Widgets:     (AirWidget)(airInfo.Widgets),
		UpdatedDate: airInfo.UpdatedDate,
	}

	return resAirInfo, nil
}
func (s *airthingService) DeleteAir(id string, userId string) error {
	filter := &repository.FilterUpdate{}
	filter.Id = id
	filter.UserId = userId

	err := s.airRepo.DelAir(filter)
	if err != nil {
		return err
	}
	return nil
}
