package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCustomerRequest struct {
	Name     string    `json:"name" bson:"name"`
	Lastname string    `json:"last_name" bson:"last_name"`
	Tel      string    `json:"tel" bson:"tel"`
	Email    string    `json:"email" bson:"email"`
	CreateAt time.Time `json:"created_date,omitempty" bson:"created_date,omitempty"`
	UpdateAt time.Time `json:"updated_date,omitempty" bson:"updated_date,omtiempty"`
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
type UpdateCustomer struct {
	Name     string    `json:"name" bson:"name"`
	Lastname string    `json:"last_name" bson:"last_name"`
	Tel      string    `json:"tel" bson:"tel"`
	Email    string    `json:"email" bson:"email"`
	UpdateAt time.Time `json:"updated_date,omitempty" bson:"updated_date,omtiempty"`
}

type CustomerRepository interface {
	CreateCustomer(*CreateCustomerRequest) (*DBCustomer, error)
	UpdateCustomer(string, *UpdateCustomer) (*DBCustomer, error)
	FindCustomerById(string) (*DBCustomer, error)
	FindCustomers() ([]*DBCustomer, error)
	DeleteCustomer(string) error
}
