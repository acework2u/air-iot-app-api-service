package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	Name      string    `json:"name" bson:"name"`
	UserId    string    `json:"userId" bson:"userId"`
	SerialNo  string    `json:"serialNo" bson:"serialNo"`
	Warranty  string    `json:"warranty" bson:"warranty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type DBDevice struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId    string             `json:"userId" bson:"userId"`
	Name      string             `json:"name" bson:"name"`
	SerialNo  string             `json:"serialNo" bson:"serialNo"`
	Warranty  string             `json:"warranty" bson:"warranty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type DevicesRepository interface {
	CreateDevice(device *Device) (*DBDevice, error)
	FindDevices() ([]*DBDevice, error)
}
