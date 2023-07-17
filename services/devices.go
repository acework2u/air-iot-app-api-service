package services

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	Name     string `json:"name" validate:"required"`
	UserId   string `json:"userId" bson:"userId"`
	SerialNo string `json:"serialNo" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Model    string `json:"model" validate:"required"`
	Warranty string `json:"warranty" validate:"required"`
}

type responseDevice struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	SerialNo string             `json:"serialNo"`
	Title    string             `json:"title"`
	Model    string             `json:"model"`
	Warranty string             `json:"warranty"`
}

type DevicesService interface {
	NewDevice(*Device) (*responseDevice, error)
	//RegisterDevice()
	//UpdateDevice()
	//DeleteDevice()
	//ConnectDevice()
}
