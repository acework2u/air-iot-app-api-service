package services

import (
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	mqtt "github.com/tech-sumit/aws-iot-device-sdk-go"
	"mime/multipart"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type (
	ShadowsValue struct {
		State struct {
			Desired  Desired  `json:"desired"`
			Reported Reported `json:"reported"`
		} `json:"state"`
	}

	ShadowsCommand struct {
		State struct {
			Desired Desired `json:"desired"`
		} `json:"state"`
	}
	Desired struct {
		Cmd string `json:"cmd"`
	}
	Reported struct {
		Message string `json:"message"`
	}
	IndoorInfo struct {
		Power    string `json:"power"`
		Mode     string `json:"mode"`
		Temp     string `json:"temp"`
		RoomTemp string `json:"roomTemp"`
		RhSet    string `json:"rhSet"`
		RhRoom   string `json:"RhRoom"`
		FanSpeed string `json:"fanSpeed"`
		Louver   string `json:"louver"`
	}
)

type ThinksService interface {
	GetCerts() (interface{}, error)
	GetUserCert(*UserReq) (interface{}, error)
	UploadToS3(file *multipart.FileHeader) (interface{}, error)
	ThingRegister(idToken string) (interface{}, error)
	ThingsConnected(idToken string, thing string) (*iotdataplane.PublishOutput, error)
	ThingsCert(idToken string) (interface{}, error)
	ThinksShadows(idToken string, rs string) (*ShadowsValue, error)
	NewAwsMqttConnect(cognitoId string) (*mqtt.AWSIoTConnection, error)
	PubGetShadows(thinkName string, shadowName string) (*IndoorInfo, error)
	PubUpdateShadows(thinkName string, payload string) (*IndoorInfo, error)
}
