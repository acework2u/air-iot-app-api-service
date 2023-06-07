package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Things struct {
	Name      string `json:"name" bson:"name"`
	SeriailNo string `json:"serial_no" bson:"serial_no"`
	Model     string `json:"model" bson:"model"`
	Certs     string `json:"thing_cert" bson:"thing_cert"`
	ImgThing  string `json:"thing_img" bson:"thing_img"`
}

type ThingsRegister struct {
	Name      string    `json:"name" bson:"name"`
	SeriailNo string    `json:"serial_no" bson:"serial_no"`
	Model     string    `json:"model" bson:"model"`
	Certs     string    `json:"thing_cert" bson:"thing_cert"`
	ImgThing  string    `json:"thing_img" bson:"thing_img"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ThingsDB struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	SeriailNo string             `json:"serial_no" bson:"serial_no"`
	Model     string             `json:"model" bson:"model"`
	Certs     string             `json:"thing_cert" bson:"thing_cert"`
	ImgThing  string             `json:"thing_img" bson:"thing_img"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ThingsUpdate struct {
	Name      string    `json:"name" bson:"name"`
	SeriailNo string    `json:"serial_no" bson:"serial_no"`
	Model     string    `json:"model" bson:"model"`
	Certs     string    `json:"thing_cert" bson:"thing_cert"`
	ImgThing  string    `json:"thing_img" bson:"thing_img"`
	UpdateAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type ThingsRepositoty interface {
	RegisterThings(*ThingsRegister) (*ThingsDB, error)
	UpdateThings(string, *ThingsUpdate) (*ThingsDB, error)
}
