package services

import (
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"mime/multipart"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ThinksService interface {
	GetCerds() (interface{}, error)
	GetUserCert(*UserReq) (interface{}, error)
	UploadToS3(file *multipart.FileHeader) (interface{}, error)
	ThingRegister(idToken string) (interface{}, error)
	ThingsConnected(idToken string) (*iotdataplane.PublishOutput, error)
	ThingsCert(idToken string) (interface{}, error)
}
