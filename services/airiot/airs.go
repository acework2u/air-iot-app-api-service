package airiot

import (
	"encoding/json"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type PubSub struct {
	Clients       []Member
	Subscriptions []Subscription
}
type Member struct {
	ID   string
	Conn *websocket.Conn
}

type WsMessage struct {
	Action  string          `json:"action"`
	Topic   string          `json:"topic"`
	Message json.RawMessage `json:"message"`
}

type Subscription struct {
	Topic  string
	Client *Member
}

type (
	AirInfo struct {
		Serial       string    `json:"serial" binding:"required"`
		UserId       string    `json:"userId"`
		Title        string    `json:"title" binding:"required"`
		RegisterDate time.Time `json:"registerDate"`
		UpdatedDate  time.Time `bson:"updatedDate"`
	}

	AirRq struct {
		Serial string `json:"serial"`
	}

	DBAirInfo struct {
		Id           primitive.ObjectID `json:"id"`
		Serial       string             `json:"serial"`
		UserId       string             `json:"userId"`
		Title        string             `json:"title"`
		RegisterDate time.Time          `json:"registerDate"`
		UpdatedDate  time.Time          `bson:"updatedDate"`
	}
	ResponseAir struct {
		Id     primitive.ObjectID `json:"id,omitempty"`
		Serial string             `json:"serial,omitempty"`
		Title  string             `json:"title,omitempty"`
		Bg     string             `json:"bg,omitempty"`
		Indoor *IndoorInfo        `json:"indoor,omitempty"`
	}
	AirThingConfig struct {
		Region          string `json:"region"`
		UserPoolId      string `json:"userPoolId"`
		CognitoClientId string `json:"cognitoClientId"`
	}
	IndoorInfo struct {
		Power    string `json:"power,omitempty"`
		Mode     string `json:"mode,omitempty"`
		Temp     string `json:"temp,omitempty"`
		RoomTemp string `json:"roomTemp,omitempty"`
		RhSet    string `json:"rhSet,omitempty"`
		RhRoom   string `json:"RhRoom,omitempty"`
		FanSpeed string `json:"fanSpeed,omitempty"`
		Louver   string `json:"louver,omitempty"`
	}

	AirIoTConfig struct {
		Region          string                   `json:"region"`
		UserPoolId      string                   `json:"userPoolId"`
		CognitoClientId string                   `json:"cognitoClientId"`
		AirRepo         repository.AirRepository `json:"airRepo"`
	}
	ShadowsValue struct {
		State struct {
			Desired  Desired  `json:"desired,omitempty"`
			Reported Reported `json:"reported,omitempty"`
		} `json:"state"`
	}
	Desired struct {
		Cmd string `json:"cmd,omitempty"`
	}
	Reported struct {
		Message string `json:"message,omitempty"`
	}
)

type AirIoTService interface {
	GetIndoorVal(serial string, shadowsName string) (interface{}, error)
	GetShadowsDocument(serial string, shadowsName string) (interface{}, error)
	CheckAwsDefault() (interface{}, error)
	CheckAwsProduct() (interface{}, error)
	CheckAws() (interface{}, error)
	ShadowsAir(clientId string, serial string) (interface{}, error)
	Ws2AddClient(client Member) *PubSub
	Ws2RemoveClient(client Member) *PubSub
	Ws2GetSubscriptions(topic string, client *Member) []*Subscription
	Ws2Subscribe(client *Member, topic string) *PubSub
	Ws2Publish(topic string, message []byte, excludeClient *Member)
	Ws2HandleReceiveMessage(client Member, messageType int, payload []byte) *PubSub
	Airlist(UserId string) ([]*ResponseAir, error)
	CheckMyAc(UserId string, serial string) bool
}
