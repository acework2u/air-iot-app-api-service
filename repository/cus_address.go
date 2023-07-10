package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	CustomerAddress struct {
		CustomerId      string    `json:"customerId" bson:"customerId"`
		Name            string    `json:"name" bson:"name"`
		LastName        string    `json:"lastName" bson:"lastName"`
		Tel             string    `json:"tel" bson:"tel"`
		Address_1       string    `json:"address_1" bson:"address_1"`
		Zipcode         int       `json:"zipcode" bson:"zipcode"`
		District        string    `json:"district" bson:"district"`
		Amphur          string    `json:"amphur" bson:"amphur"`
		Province        string    `json:"province" json:"province"`
		Tax             string    `json:"tax" json:"tax"`
		Tax_used        bool      `json:"tax_used" bson:"tax_used"`
		Tax_default     bool      `json:"tax_default" bson:"tax_default"`
		Address_default bool      `json:"address_default" bson:"address_default"`
		UpdateAt        time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	}

	UpdateAddress struct {
		CustomerId      string    `json:"customerId" bson:"customerId"`
		Name            string    `json:"name" bson:"name"`
		LastName        string    `json:"lastName" bson:"lastName"`
		Tel             string    `json:"tel" bson:"tel"`
		Address_1       string    `json:"address_1" bson:"address_1"`
		Zipcode         int       `json:"zipcode" bson:"zipcode"`
		District        string    `json:"district" bson:"district"`
		Amphur          string    `json:"amphur" bson:"amphur"`
		Province        string    `json:"province" json:"province"`
		Tax             string    `json:"tax" json:"tax"`
		Tax_used        bool      `json:"tax_used" bson:"tax_used"`
		Tax_default     bool      `json:"tax_default" bson:"tax_default"`
		Address_default bool      `json:"address_default" bson:"address_default"`
		UpdateAt        time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	}
	DBAddress struct {
		Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		CustomerId      string             `json:"customerId" bson:"customerId"`
		Name            string             `json:"name" bson:"name"`
		LastName        string             `json:"lastName" bson:"lastName"`
		Tel             string             `json:"tel" bson:"tel"`
		Address_1       string             `json:"address_1" bson:"address_1"`
		Zipcode         int                `json:"zipcode" bson:"zipcode"`
		District        string             `json:"district" bson:"district"`
		Amphur          string             `json:"amphur" bson:"amphur"`
		Province        string             `json:"province" json:"province"`
		Tax             string             `json:"tax" json:"tax"`
		Tax_used        bool               `json:"tax_used" bson:"tax_used"`
		Tax_default     bool               `json:"tax_default" bson:"tax_default"`
		Address_default bool               `json:"address_default" bson:"address_default"`
		UpdateAt        time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	}
)

type AddressRepository interface {
	CreateNewAddress(address *CustomerAddress) (*DBAddress, error)
	UpdateAddress(string, *UpdateCustomer) (*DBAddress, error)
	DeleteAddress(string) error
	FindAddress() ([]*DBAddress, error)
	FindAddressId(string) (*DBAddress, error)
}
