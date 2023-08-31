package airiot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
)

var ClientAwsMqtt *mqtt.AWSIoTConnection

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
	getThingShadowOutput, err := s.IotData.GetThingShadow(s.Ctx, subTopic)
	if err != nil {
		return nil, err
	}
	dataAc := &utils.ShadowAcceptStrut{}
	err = json.Unmarshal([]byte(getThingShadowOutput.Payload), &dataAc)
	if err != nil {
		return nil, err
	}
	if len(dataAc.State.Reported.Message) < 0 {
		return nil, err
	}

	acVal := &dataAc.State.Reported.Message

	decodeShadow, _ := utils.GetClaimsFromToken(*acVal)

	return decodeShadow, nil
}
func (s *airIoTService) GetShadowsDocument(serial string, shadowsName string) (interface{}, error) {

	subTopic := &iotdataplane.GetThingShadowInput{
		ThingName:  aws.String(serial),
		ShadowName: aws.String(shadowsName),
	}
	getThingShadowOutput, err := s.IotData.GetThingShadow(s.Ctx, subTopic)
	if err != nil {
		return nil, err
	}
	dataAc := &utils.ShadowAcceptStrut{}
	err = json.Unmarshal([]byte(getThingShadowOutput.Payload), &dataAc)
	if err != nil {
		return nil, err
	}
	if len(dataAc.State.Reported.Message) < 0 {
		return nil, err
	}

	acVal := &dataAc.State.Reported.Message

	decodeShadow, _ := utils.GetClaimsFromToken(*acVal)

	_ = decodeShadow

	return acVal, nil
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
func (s *airIoTService) ShadowsAir(clientId string, serial string) (interface{}, error) {

	if len(clientId) < 0 {
		return nil, errors.New("userId is empty")
	}
	if len(serial) < 0 {
		return nil, errors.New("serial is empty")
	}
	var err error
	ClientAwsMqtt, err = NewAwsMqttConn(clientId)
	if err != nil {
		fmt.Println(err.Error())

		panic(err)

		return nil, err

	}

	subTopic := fmt.Sprintf("$aws/things/%v/shadow/name/air-users/update/documents", serial)
	shadowsVal := &ShadowsValue{}

	result := make(chan *ShadowsValue)
	go iotSub(subTopic, result)
	shadowsVal = <-result

	return shadowsVal, nil
}

func NewAwsMqttConn(clientId string) (*mqtt.AWSIoTConnection, error) {
	clientMq := &mqtt.AWSIoTConnection{}
	clientMq, err := mqtt.NewConnection(mqtt.Config{
		KeyPath:  "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-private.pem.key",
		CertPath: "./certs/cert7/57c2a591aca1a833d146cb9283ce66770ed9d65a4be0cd90a754ec8f92679371-certificate.pem.crt",
		CAPath:   "./certs/cert7/AmazonRootCA1.pem",
		ClientId: *aws.String(clientId),
		Endpoint: "a18xth5rea73tz-ats.iot.ap-southeast-1.amazonaws.com",
	})
	if err != nil {
		return nil, err
	}
	return clientMq, err

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
			fmt.Println("shadowsVal")
			fmt.Println(shadowsVal)
			result <- shadowsVal
		})
	}()

}
