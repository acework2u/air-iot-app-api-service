package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AirInfo struct {
	Serial       string    `json:"serial" bson:"serial" binding:"required"`
	UserId       string    `json:"userId" bson:"userId"`
	Title        string    `json:"title" bson:"title" binding:"required"`
	RegisterDate time.Time `json:"registerDate" bson:"registerDate"`
	UpdatedDate  time.Time `bson:"updatedDate" bson:"updatedDate"`
}
type DBAirInfo struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Serial       string             `json:"serial" bson:"serial"`
	UserId       string             `json:"userId" bson:"userId"`
	Title        string             `json:"title" bson:"title"`
	RegisterDate time.Time          `json:"registerDate" bson:"registerDate,omitempty"`
	UpdatedDate  time.Time          `bson:"updatedDate" bson:"updatedDate,omitempty"`
}

type AirThinkService interface {
	GetCerts(string2 string) (interface{}, error)
	ThingConnect(idToken string) (interface{}, error)
	AddAir(info *AirInfo) (*DBAirInfo, error)
}
