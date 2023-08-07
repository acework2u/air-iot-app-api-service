package services

import (
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"mime/multipart"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ShadowsValue struct {
	State struct {
		Desired  Desired  `json:"desired"`
		Reported Reported `json:"reported"`
	} `json:"state"`
}
type ShadowsCommand struct {
	State struct {
		Desired Desired `json:"desired"`
	} `json:"state"`
}

type Desired struct {
	Cmd string `json:"cmd"`
}
type Reported struct {
	Message string `json:"message"`
}

type ThinksService interface {
	GetCerds() (interface{}, error)
	GetUserCert(*UserReq) (interface{}, error)
	UploadToS3(file *multipart.FileHeader) (interface{}, error)
	ThingRegister(idToken string) (interface{}, error)
	ThingsConnected(idToken string, thing string) (*iotdataplane.PublishOutput, error)
	ThingsCert(idToken string) (interface{}, error)
	ThinksShadows(idToken string, rs string) (interface{}, error)
}
