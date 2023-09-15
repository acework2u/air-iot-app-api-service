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
	Bg           string    `bson:"bg" json:"bg"`
	Widgets      AirWidget `json:"widgets" bson:"widgets"`
	RegisterDate time.Time `json:"registerDate" bson:"registerDate"`
	UpdatedDate  time.Time `bson:"updatedDate" bson:"updatedDate"`
}
type UpdateAirInfo struct {
	Serial      string    `json:"serial" bson:"serial"`
	UserId      string    `json:"userId" bson:"userId"`
	Title       string    `json:"title" bson:"title"`
	Bg          string    `bson:"bg" json:"bg"`
	Status      bool      `json:"status,omitempty" bson:"status,omitempty"`
	Widgets     AirWidget `json:"widgets" bson:"widgets"`
	UpdatedDate time.Time `bson:"updatedDate" bson:"updatedDate"`
}
type FilterUpdate struct {
	Id     string `json:"id" bson:"_id"`
	UserId string `json:"userId" bson:"userId"`
}
type DBAirInfo struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Serial       string             `json:"serial" bson:"serial"`
	UserId       string             `json:"userId" bson:"userId"`
	Title        string             `json:"title" bson:"title"`
	Bg           string             `bson:"bg" json:"bg"`
	Status       bool               `json:"status" bson:"status"`
	Widgets      AirWidget          `json:"widgets" bson:"widgets"`
	RegisterDate time.Time          `json:"registerDate" bson:"registerDate,omitempty"`
	UpdatedDate  time.Time          `bson:"updatedDate" bson:"updatedDate,omitempty"`
}
type AirWidget struct {
	Swing             bool `json:"swing" bson:"swing" default:"true"`
	Mode              bool `json:"mode" bson:"mode" default:"true"`
	FanSpeed          bool `json:"fanSpeed" bson:"fanSpeed" default:"true"`
	Schedule          bool `json:"schedule" bson:"schedule"`
	Engineer          bool `json:"engineer" bson:"engineer"`
	Energy            bool `json:"energy" bson:"energy"`
	UltrafineParticle bool `json:"ultrafineParticle" bson:"ultrafineParticle"`
	Ewarranty         bool `json:"ewarranty" bson:"ewarranty"`
	Filter            bool `json:"filter"bson:"filter"`
	Sleep             bool `json:"sleep"bson:"sleep"`
}

type AirRepository interface {
	RegisterAir(info *AirInfo) (*DBAirInfo, error)
	UpdateAir(filter *FilterUpdate, info *UpdateAirInfo) (*DBAirInfo, error)
	Airs(userId string) ([]*DBAirInfo, error)
}
