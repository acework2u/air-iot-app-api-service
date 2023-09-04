package airiot

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type PubSub struct {
	Clients       []Client
	Subscriptions []Subscription
}

type WsMessage struct {
	Action  string          `json:"action"`
	Topic   string          `json:"topic"`
	Message json.RawMessage `json:"message"`
}

type Subscription struct {
	Topic  string
	Client *Client
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
		Id     primitive.ObjectID `json:"id"`
		Serial string             `json:"serial"`
		Title  string             `json:"title"`
		Bg     string             `json:"bg"`
		Indoor *IndoorInfo        `bson:"indoor"`
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
		Region          string `json:"region"`
		UserPoolId      string `json:"userPoolId"`
		CognitoClientId string `json:"cognitoClientId"`
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
	Ws2AddClient(client Client) *PubSub
	Ws2RemoveClient(client Client) *PubSub
	Ws2GetSubscriptions(topic string, client *Client) []Subscription
	Ws2Subscribe(client *Client, topic string) *PubSub
	Ws2Publish(topic string, message []byte, excludeClient *Client)
	Ws2HandleReceiveMessage(client Client, messageType int, payload []byte) *PubSub
}
