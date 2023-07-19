package services

import "mime/multipart"

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ThinksService interface {
	GetCerds() (interface{}, error)
	GetUserCert(*UserReq) (interface{}, error)
	UploadToS3(file *multipart.FileHeader) (interface{}, error)
}
