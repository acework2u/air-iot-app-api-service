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

	DBProductInfoUpdate struct {
		Serial          string      `json:"serial" bson:"serial,omitempty"`
		Status          bool        `json:"status" bson:"status,omitempty"`
		Active          bool        `json:"active" bson:"active,omitempty"`
		ProductInfo     ProductInfo `json:"productInfo" bson:"productInfo,omitempty"`
		Production      time.Time   `json:"production,omitempty" bson:"production,omitempty"`
		DefaultWarranty time.Time   `json:"defaultWarranty" bson:"defaultWarranty,omitempty"`
		EWarranty       EWarranty   `json:"EWarranty" bson:"EWarranty,omitempty"`
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
<<<<<<< HEAD
	GetProducts() ([]*DBProduct, error)
=======
>>>>>>> ad1f98be097d983c078b0925f74ee2be200245ae
	CreateProduct(product *Product) (*DBProduct, error)
	UpdateProduct(serial string, product *Product) (*DBProduct, error)
	UpdateProductInfo(serial string, productInfo *DBProductInfoUpdate) (*DBProduct, error)
	DeleteProduct(serial string) error
	UpdateEWarranty(serial string) (*DBProductInfoUpdate, error)
}
