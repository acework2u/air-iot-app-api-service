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

type CreateCustomerRequest2 struct {
	UserSub       string    `json:"usersub" bson:"usersub"`
	Name          string    `json:"name" bson:"name"`
	Lastname      string    `json:"last_name" bson:"last_name"`
	Email         string    `json:"email" bson:"email"`
	Mobile        string    `json:"mobile_no" bson:"mobile"`
	UserConfirmed bool      `json:"UserConfirmed" bson:"UserConfirmed"`
	CreateAt      time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt      time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBCustomer2 struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserSub       string             `json:"usersub" bson:"usersub"`
	Name          string             `json:"name" bson:"name"`
	Lastname      string             `json:"last_name" bson:"last_name"`
	Email         string             `json:"email" bson:"email"`
	Mobile        string             `json:"mobile" bson:"mobile"`
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

type UpdateCustomer struct {
	Name     string    `json:"name" bson:"name"`
	Lastname string    `json:"lastname" bson:"last_name"`
	Mobile   string    `json:"mobile" bson:"mobile"`
	UpdateAt time.Time `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
}

type CustomerRepository interface {
	CreateCustomer(*CreateCustomerRequest) (*DBCustomer, error)
	NewCustomer(*CreateCustomerRequest2) (*DBCustomer2, error)
	UpdateCustomer(string, *UpdateCustomer) (*DBCustomer, error)
	FindCustomerById(string) (*DBCustomer, error)
	FindCustomers() ([]*DBCustomer, error)
	DeleteCustomer(string) error
	FindCustomerID(string) (*DBCustomer2, error)
}
