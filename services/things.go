package services

import (
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"mime/multipart"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type (
	ShadowsValue struct {
		State struct {
			Desired  Desired  `json:"desired,omitempty"`
			Reported Reported `json:"reported,omitempty"`
		} `json:"state"`
	}

	ShadowsCommand struct {
		State struct {
			Desired Desired `json:"desired,omitempty"`
		} `json:"state"`
	}
	Desired struct {
		Cmd string `json:"cmd,omitempty"`
	}
	Reported struct {
		Message string `json:"message,omitempty"`
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
		Aps      string `json:"aps,omitempty"`
		OzoneGen string `json:"ozoneGen"`
	}
	ShadowsAccepted struct {
		State struct {
			Reported struct {
				Message string `json:"message"`
			} `json:"reported"`
		} `json:"state"`
		Metadata struct {
			Reported struct {
				Message struct {
					Timestamp int `json:"timestamp"`
				} `json:"message"`
			} `json:"reported"`
		} `json:"metadata"`
		Version   int `json:"version"`
		Timestamp int `json:"timestamp"`
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
	NewAwsMqttConnect(cognitoId string) error
	PubGetShadows(thinkName string, shadowName string) (*IndoorInfo, error)
	PubUpdateShadows(thinkName string, payload string) (*IndoorInfo, error)
}
