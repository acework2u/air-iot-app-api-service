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
	"github.com/gorilla/websocket"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
	"time"
)

var ClientAwsMqtt *mqtt.AWSIoTConnection
var upgerder = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	PubSub          *PubSub
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
		airRepo:   cfg.AirRepo,
		Ctx:       context.TODO(),
		PubSub:    &PubSub{},
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
func (s *airIoTService) Ws2AddClient(client Member) *PubSub {

	s.PubSub.Clients = append(s.PubSub.Clients, client)
	payload := []byte("HI " + client.ID)
	client.Conn.WriteMessage(1, payload)
	return s.PubSub
}
func (s *airIoTService) Ws2RemoveClient(client Member) *PubSub {

	for index, sub := range s.PubSub.Subscriptions {
		if client.ID == sub.Client.ID {
			s.PubSub.Subscriptions = append(s.PubSub.Subscriptions[:index], s.PubSub.Subscriptions[index+1:]...)
		}
	}
	// remove client from the list
	for index, c := range s.PubSub.Clients {
		if c.ID == client.ID {
			s.PubSub.Clients = append(s.PubSub.Clients[:index], s.PubSub.Clients[index+1:]...)
		}
	}

	return s.PubSub
}
func (s *airIoTService) Ws2GetSubscriptions(topic string, client *Member) []*Subscription {
	subscriptionList := []*Subscription{}
	for _, subscription := range s.PubSub.Subscriptions {
		if client != nil {
			if subscription.Client.ID == client.ID && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, &subscription)
			} else {
				if subscription.Topic == topic {
					subscriptionList = append(subscriptionList, &subscription)
				}
			}
		}
	}
	return subscriptionList
}
func (s *airIoTService) Ws2Subscribe(client *Member, topic string) *PubSub {

	clientSubs := s.Ws2GetSubscriptions(topic, client)
	if len(clientSubs) > 0 {
		return s.PubSub
	}

	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
	}
	s.PubSub.Subscriptions = append(s.PubSub.Subscriptions, newSubscription)

	fmt.Println("s.PubSub.Subscriptions")
	fmt.Println(s.PubSub.Subscriptions)

	return s.PubSub
}
func (s *airIoTService) Ws2Publish(topic string, message []byte, excludeClient *Member) {

	subscriptions := s.Ws2GetSubscriptions(topic, nil)

	fmt.Println(len(subscriptions))
	fmt.Println(subscriptions)
	for _, sub := range subscriptions {
		fmt.Printf("Sending to Client is %s message is %s", sub.Client.ID, message)
		sub.Client.Ws2Send(message)
	}

}
func (s *airIoTService) Ws2Unsubscribe(client *Member, topic string) *PubSub {

	for index, sub := range s.PubSub.Subscriptions {
		if sub.Client.ID == client.ID && sub.Topic == topic {
			s.PubSub.Subscriptions = append(s.PubSub.Subscriptions[:index], s.PubSub.Subscriptions[index+1:]...)
		}
	}

	return s.PubSub
}
func (s *airIoTService) Ws2HandleReceiveMessage(client Member, messageType int, payload []byte) *PubSub {
	m := WsMessage{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Println("This is not correct message payload")
		return s.PubSub
	}

	switch m.Action {

	case PUBLISH:
		fmt.Println("This is publish new message")
		s.Ws2Publish(m.Topic, m.Message, nil)
		break
	case SUBSCRIBE:
		s.Ws2Subscribe(&client, m.Topic)

		checkAc := s.CheckMyAc(client.ID, m.Topic)
		//fmt.Println("new subscriber to topic ", m.Topic, m.Message, len(s.PubSub.Subscriptions), client.ID)
		if checkAc {
			go func() {
				for {
					serial := string(m.Topic)
					shadowName := "air-users"
					res, err := s.GetShadowsDocument(serial, shadowName)
					if err != nil {
						return
					}

					acInfo, _ := json.Marshal(res)

					client.Conn.WriteMessage(1, []byte(acInfo))
					fmt.Println(s.PubSub.Subscriptions)

					time.Sleep(4 * time.Second)
				}

			}()

		}

		break
	case UNSUBSCRIBE:
		fmt.Println("Client want to unsubscribe the topic", m.Topic, client.ID)
		s.Ws2Unsubscribe(&client, m.Topic)
		break
	default:
		client.Conn.Close()
		break
	}

	return s.PubSub
}
func (s *airIoTService) Airlist(UserId string) ([]*ResponseAir, error) {
	airList := []*ResponseAir{}
	res, err := s.airRepo.Airs(UserId)
	if err != nil {
		return nil, err
	}

	for _, items := range res {
		item := &ResponseAir{
			Serial: items.Serial,
		}

		airList = append(airList, item)
	}

	return airList, err
}
func (s *airIoTService) CheckMyAc(UserId string, serial string) bool {
	res, err := s.Airlist(UserId)
	if err != nil {
		return false
	}
	for _, items := range res {
		if string(items.Serial) == string(serial) {
			return true
		}
	}
	return false
}

func (c *Member) Ws2Send(message []byte) error {
	return c.Conn.WriteMessage(1, message)
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
