package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AirInfo struct {
	Serial       string    `json:"serial" binding:"required"`
	UserId       string    `json:"userId"`
	Title        string    `json:"title" binding:"required"`
	RegisterDate time.Time `json:"registerDate"`
	UpdatedDate  time.Time `bson:"updatedDate"`
}
type DBAirInfo struct {
	Id           primitive.ObjectID `json:"id"`
	Serial       string             `json:"serial"`
	UserId       string             `json:"userId"`
	Title        string             `json:"title"`
	RegisterDate time.Time          `json:"registerDate"`
	UpdatedDate  time.Time          `bson:"updatedDate"`
}
type ResponseAir struct {
	Id     primitive.ObjectID `json:"id"`
	Serial string             `json:"serial"`
	Title  string             `json:"title"`
}

type AirThinkService interface {
	GetCerts(string2 string) (interface{}, error)
	ThingConnect(idToken string) (interface{}, error)
	AddAir(info *AirInfo) (*DBAirInfo, error)
	GetAirs(userId string) ([]*ResponseAir, error)
}
