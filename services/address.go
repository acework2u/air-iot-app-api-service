package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Filter struct {
		Id     string `uri:"id" json:"id" binding:"required,uuid" validate:"required"`
		UserId string `json:"userid"`
	}
	CustomerAddress struct {
		CustomerId      string    `json:"customerId,omitempty" bson:"customerId,omitempty"`
		Name            string    `json:"name" bson:"name" binding:"required" `
		LastName        string    `json:"lastName" bson:"lastName" binding:"required"`
		Tel             string    `json:"tel" bson:"tel" binding:"required"`
		Address         string    `json:"address" bson:"address" binding:"required"`
		Zipcode         int       `json:"zipcode" bson:"zipcode" binding:"required"`
		District        string    `json:"district" bson:"district" binding:"required"`
		Amphur          string    `json:"amphur" bson:"amphur" binding:"required"`
		Province        string    `json:"province" json:"province" binding:"required"`
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
		Address         string             `json:"address" bson:"address"`
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

	ResponseAddress struct {
		Id             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name           string             `json:"name" bson:"name"`
		LastName       string             `json:"lastName" bson:"lastName"`
		Tel            string             `json:"tel" bson:"tel"`
		Address        string             `json:"address" bson:"address"`
		Zipcode        int                `json:"zipcode" bson:"zipcode"`
		District       string             `json:"district" bson:"district"`
		Amphur         string             `json:"amphur" bson:"amphur"`
		Province       string             `json:"province" json:"province"`
		Tax            string             `json:"tax" json:"tax"`
		TaxUsed        bool               `json:"tax_used" bson:"tax_used"`
		TaxDefault     bool               `json:"tax_default" bson:"tax_default"`
		AddressDefault bool               `json:"address_default" bson:"address_default"`
	}
)

type AddressService interface {
	NewAddress(address *CustomerAddress) (*DBAddress, error)
	AllAddress(userid string) ([]*ResponseAddress, error)
	DelAddress(id string) error
	UpdateAddress(filter *Filter, addressInfo *CustomerAddress) (*DBAddress, error)
}
