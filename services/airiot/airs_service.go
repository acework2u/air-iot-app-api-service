package airiot

import (
	"context"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type airIoTService struct {
	Cfg             *aws.Config
	StsClient       *sts.Client
	IotClient       *iot.Client
	IotData         *iotdataplane.Client
	airRepo         repository.AirRepository
	userPoolId      string
	cognitoClientId string
	region          string
	Ctx             context.Context
}

func NewAirIoTService(cfg *AirIoTConfig) AirIoTService {
	region := "ap-southeast-1"
	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	//awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(cfg.Region), config.WithSharedConfigProfile("production"))

	if err != nil {
		fmt.Println("config is problem")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Config Work-->")

	iotClient := iot.NewFromConfig(awsCfg)
	iotData := iotdataplane.NewFromConfig(awsCfg)

	return &airIoTService{
		Cfg:       &awsCfg,
		IotClient: iotClient,
		IotData:   iotData,
		Ctx:       context.TODO(),
	}
}

func (s *airIoTService) GetIndoorVal(serial string, shadowsName string) (interface{}, error) {

	subTopic := &iotdataplane.GetThingShadowInput{
		ThingName:  aws.String(serial),
		ShadowName: aws.String(shadowsName),
	}
	getThingShadowOutput := &iotdataplane.GetThingShadowOutput{}
	getThingShadowOutput, err := s.IotData.GetThingShadow(s.Ctx, subTopic)
	if err != nil {
		return nil, err
	}

	return getThingShadowOutput, nil
}
func (s *airIoTService) CheckAwsDefault() (interface{}, error) {
	region := "ap-southeast-1"
	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region), config.WithSharedConfigProfile("default"))

	if err != nil {
		fmt.Println("config is problem")
		return nil, err

	}
	data := fmt.Sprintf("AWS load Work : %v", awsCfg)
	fmt.Println(awsCfg)

	return data, err

}
func (s *airIoTService) CheckAwsProduct() (interface{}, error) {

	region := "ap-southeast-1"
	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region), config.WithSharedConfigProfile("production"))

	if err != nil {
		fmt.Println("config is problem")
		return nil, err

	}

	data := fmt.Sprintf("AWS load Work : %v", awsCfg)
	fmt.Println(awsCfg)

	return data, err

}

func (s *airIoTService) CheckAws() (interface{}, error) {
	region := "ap-southeast-1"
	awsCfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))

	if err != nil {
		fmt.Println("config is problem")
		return nil, err

	}

	data := fmt.Sprintf("AWS load Work : %v", awsCfg)
	fmt.Println(awsCfg)

	return data, err
}
