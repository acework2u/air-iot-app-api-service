package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	Name      string    `json:"name" bson:"name"`
	Title     string    `json:"title" bson:"title"`
	Model     string    `json:"model" bson:"model"`
	SerialNo  string    `json:"serialNo" bson:"serialNo"`
	Warranty  string    `json:"warranty" bson:"warranty"`
	UserId    string    `json:"userId" bson:"userId"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type DBDevice struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId    string             `json:"userId" bson:"userId"`
	Name      string             `json:"name" bson:"name"`
	Title     string             `json:"title" bson:"title"`
	Model     string             `json:"model" bson:"model"`
	SerialNo  string             `json:"serialNo" bson:"serialNo"`
	Warranty  string             `json:"warranty" bson:"warranty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type DeviceRequest struct {
	UserId string `json:"userid" bson:"userId"`
}

type DevicesRepository interface {
	CreateDevice(device *Device) (*DBDevice, error)
	FindDevices(request *DeviceRequest) ([]*DBDevice, error)
	CheckDupDevice(userId string, serialNo string) (int64, error)
}
