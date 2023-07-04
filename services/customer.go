package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateCustomerRequest struct {
	Name     string    `json:"name" bson:"name" binding:"required"`
	Lastname string    `json:"last_name" binding:"required"`
	Tel      string    `json:"tel" bson:"tel" binding:"required"`
	Email    string    `json:"email" bson:"email"`
	CreateAt time.Time `json:"created_date,omitempty" bson:"created_date,omitempty"`
	UpdateAt time.Time `json:"updated_date,omitempty" bson:"updated_date,omtiempty"`
}

type CreateCustomerRequest2 struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserSub       string             `json:"usersub" bson:"usersub"`
	Name          string             `json:"name" bson:"name"`
	Lastname      string             `json:"last_name" bson:"last_name"`
	Email         string             `json:"email" bson:"email"`
	Mobile        string             `json:"mobile_no" bson:"mobile"`
	UserConfirmed bool               `json:"UserConfirmed" bson:"UserConfirmed"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DbCustomerResponse2 struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserSub       string             `json:"usersub" bson:"usersub"`
	Name          string             `json:"name" bson:"name"`
	Lastname      string             `json:"last_name" bson:"last_name"`
	Email         string             `json:"email" bson:"email"`
	Mobile        string             `json:"mobile_no" bson:"mobile"`
	UserConfirmed bool               `json:"UserConfirmed" bson:"UserConfirmed"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBCustomer struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Lastname string             `json:"last_name" bson:"last_name"`
	Tel      string             `json:"tel" bson:"tel"`
	Email    string             `json:"email" bson:"email"`
	CreateAt time.Time          `json:"created_date,omitempty" bson:"created_date,omitempty"`
	UpdateAt time.Time          `json:"updated_date,omitempty" bson:"updated_date,omtiempty"`
}

type CustomerResponse struct {
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
type UpdateCustomer struct {
	Name     string    `json:"name" bson:"name"`
	Lastname string    `json:"last_name"`
	Tel      string    `json:"tel" bson:"tel"`
	Email    string    `json:"email" bson:"emil"`
	UpdateAt time.Time `json:"updated_date,omitempty" bson:"updated_date,omtiempty"`
}

type CustomerService interface {
	CreateNewCustomer(*CreateCustomerRequest) (*DBCustomer, error)
	AllCustomers() ([]*DBCustomer, error)
	UpdateCustomer(string, *UpdateCustomer) (*DBCustomer, error)
	DeleteCustomer(string) error
}
