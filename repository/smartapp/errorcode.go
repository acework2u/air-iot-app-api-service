package smartapp

import "go.mongodb.org/mongo-driver/bson/primitive"

type AcErrorCodeDb struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code     string             `bson:"code"`
	Detail   string             `bson:"detail"`
	Title    string             `bson:"title"`
	Unit     string             `bson:"unit"`
	VideoUrl string             `bson:"url_video"`
	WebUrl   string             `bson:"url_web"`
}
type AcErrorCode struct {
	Code     string `bson:"code"`
	Detail   string `bson:"detail"`
	Title    string `bson:"title"`
	Unit     string `bson:"unit"`
	VideoUrl string `bson:"url_video"`
	WebUrl   string `bson:"url_web"`
}

type ErrorCodeRepo interface {
	ErrorCodeList() ([]*AcErrorCodeDb, error)
	AcErrCode(code string) (*AcErrorCodeDb, error)
}
