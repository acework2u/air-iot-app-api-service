package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AirInfo struct {
	Serial       string    `json:"serial" binding:"required"`
	UserId       string    `json:"userId"`
	Title        string    `json:"title" binding:"required"`
	Status       bool      `json:"status" bson:"status"`
	Bg           string    `json:"bg,omitempty"`
	RegisterDate time.Time `json:"registerDate,omitempty"`
	UpdatedDate  time.Time `bson:"updatedDate,omitempty"`
}
type UpdateAirInfo struct {
	Serial      string    `json:"serial" validate:"required" binding:"required"`
	UserId      string    `json:"userId" validate:"required"`
	Title       string    `json:"title" validate:"required" binding:"required"`
	Bg          string    `json:"bg,omitempty"`
	UpdatedDate time.Time `bson:"updatedDate,omitempty"`
	Widgets     AirWidget `json:"widgets"`
}
type FilterUpdate struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
}
type DBAirInfo struct {
	Id           primitive.ObjectID `json:"id"`
	Serial       string             `json:"serial"`
	UserId       string             `json:"userId,omitempty"`
	Title        string             `json:"title"`
	Bg           string             `json:"bg"`
	Status       bool               `json:"status"`
	Widgets      AirWidget          `json:"widgets"`
	RegisterDate time.Time          `json:"registerDate,omitempty"`
	UpdatedDate  time.Time          `bson:"updatedDate,omitempty"`
}
type ResponseAir struct {
	Id      primitive.ObjectID `json:"id"`
	Serial  string             `json:"serial"`
	Title   string             `json:"title"`
	Bg      string             `json:"bg"`
	Indoor  *IndoorInfo        `json:"indoor"`
	Widgets AirWidget          `json:"widgets"`
}
type AirThingConfig struct {
	Region          string `json:"region"`
	UserPoolId      string `json:"userPoolId"`
	CognitoClientId string `json:"cognitoClientId"`
}
type AirWidget struct {
	Swing             bool `json:"swing,omitempty" default:"true"`
	Mode              bool `json:"mode,omitempty" default:"true"`
	FanSpeed          bool `json:"fanSpeed,omitempty" default:"true"`
	Schedule          bool `json:"schedule,omitempty"`
	Engineer          bool `json:"engineer,omitempty"`
	Energy            bool `json:"energy,omitempty"`
	UltrafineParticle bool `json:"ultrafineParticle,omitempty"`
	Ewarranty         bool `json:"ewarranty,omitempty" default:"true"`
	Filter            bool `json:"filter,omitempty"`
	Sleep             bool `json:"sleep,omitempty"`
}

type AirThinkService interface {
	GetCerts(string2 string) (interface{}, error)
	ThingConnect(idToken string) (interface{}, error)
	AddAir(info *AirInfo) (*DBAirInfo, error)
	GetAirs(userId string) ([]*ResponseAir, error)
	UpdateAir(filter *FilterUpdate, info *UpdateAirInfo) (*DBAirInfo, error)
	DeleteAir(id string, userId string) error
}
