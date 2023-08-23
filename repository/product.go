package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Product struct {
		Serial          string      `json:"serial" bson:"serial"`
		Status          bool        `json:"status" bson:"status"`
		Active          bool        `json:"active" bson:"active"`
		ProductInfo     ProductInfo `json:"productInfo" bson:"productInfo"`
		Production      time.Time   `json:"production,omitempty" bson:"production,omitempty"`
		DefaultWarranty time.Time   `json:"defaultWarranty" bson:"defaultWarranty"`
		EWarranty       EWarranty   `json:"EWarranty" bson:"EWarranty"`
	}
	DBProduct struct {
		Id              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Serial          string             `json:"serial" bson:"serial"`
		Status          bool               `json:"status" bson:"status"`
		Active          bool               `json:"active" bson:"active"`
		ProductInfo     ProductInfo        `json:"productInfo" bson:"productInfo"`
		Production      time.Time          `json:"production,omitempty" bson:"production,omitempty"`
		DefaultWarranty time.Time          `json:"defaultWarranty" bson:"defaultWarranty"`
		EWarranty       EWarranty          `json:"EWarranty" bson:"EWarranty"`
	}

	EWarranty struct {
		EWarranty  time.Time `json:"EWarranty" bson:"EWarranty"`
		ActiveDate time.Time `json:"activeDate" bson:"activeDate"`
	}

	ProductInfo struct {
		Title        string `json:"title" bson:"title"`
		Model        string `json:"model" bson:"model"`
		Sku          string `json:"sku" bson:"sku"`
		Mpn          string `json:"mpn" bson:"mpn"`
		ProductImage string `json:"productImage" bson:"productImage"`
	}
)

type ProductRepository interface {
	GetProduct(serial string) (*DBProduct, error)
	GetProducts() ([]*DBProduct, error)
	CreateProduct(product *Product) (*DBProduct, error)
	DeleteProduct(serial string) error
}
