package airiot

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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
}
