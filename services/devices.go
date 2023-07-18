package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	Name     string `json:"name" validate:"required"`
	UserId   string `json:"userId" bson:"userId"`
	SerialNo string `json:"serialNo" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Model    string `json:"model" validate:"required"`
	Warranty string `json:"warranty" validate:"required"`
}

type ReqUpdateDevice struct {
	Name      string    `json:"name" validate:"required"`
	UserId    string    `json:"userId" bson:"userId"`
	SerialNo  string    `json:"serialNo" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Model     string    `json:"model" validate:"required"`
	Warranty  string    `json:"warranty" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type DeviceRequest struct {
	UserId string `json:"userid" binding:"required"`
}

type DeviceFilter struct {
	UserId string `json:"userId" bson:"userId,omitempty"`
	Id     string `json:"id" bson:"_id,omitempty"`
}

type ResponseDevice struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	SerialNo string             `json:"serialNo"`
	Title    string             `json:"title"`
	Model    string             `json:"model"`
	Warranty string             `json:"warranty"`
}

type DevicesService interface {
	NewDevice(*Device) (*ResponseDevice, error)
	ListDevice(*DeviceRequest) ([]*ResponseDevice, error)
	CheckDup(string, string) int32
	UpdateDevice(string, *ReqUpdateDevice) (*ResponseDevice, error)
	DeleteDevice(filter *DeviceFilter) (bool, error)
}
