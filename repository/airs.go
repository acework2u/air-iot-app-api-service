package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AirInfo struct {
	Serial       string    `json:"serial" bson:"serial"`
	UserId       string    `json:"userId" bson:"userId"`
	Title        string    `json:"title" bson:"title"`
	Status       bool      `json:"status" bson:"status"`
	RegisterDate time.Time `json:"registerDate" bson:"registerDate"`
	UpdatedDate  time.Time `bson:"updatedDate" bson:"updatedDate"`
}
type DBAirInfo struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Serial       string             `json:"serial" bson:"serial"`
	UserId       string             `json:"userId" bson:"userId"`
	Title        string             `json:"title" bson:"title"`
	Status       bool               `json:"status" bson:"status"`
	RegisterDate time.Time          `json:"registerDate" bson:"registerDate,omitempty"`
	UpdatedDate  time.Time          `bson:"updatedDate" bson:"updatedDate,omitempty"`
}

type AirRepository interface {
	RegisterAir(info *AirInfo) (*DBAirInfo, error)
	UpdateAir(userId string, info *AirInfo) (*DBAirInfo, error)
	Airs(userId string) ([]*DBAirInfo, error)
}
