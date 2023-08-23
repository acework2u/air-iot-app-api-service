package services

import (
	"time"
)

type ProductNew struct {
	Serial      string      `json:"serial" binding:"required"`
	ProductInfo ProductInfo `json:"productInfo" binding:"required"`
}

type ProductInfo struct {
	Title        string `json:"title" binding:"required"`
	Model        string `json:"model" binding:"required"`
	Sku          string `json:"sku" binding:"required"`
	Mpn          string `json:"mpn" binding:"required"`
	ProductImage string `json:"image,omitempty"`
}
type ProductUpdate struct {
	Serial          string      `json:"serial,omitempty"`
	Status          bool        `json:"status,omitempty"`
	Active          bool        `json:"active,omitempty"`
	ProductInfo     ProductInfo `json:"productInfo,omitempty"`
	Production      time.Time   `json:"production,omitempty"`
	DefaultWarranty time.Time   `json:"defaultWarranty,omitempty"`
	EWarranty       EWarranty   `json:"EWarranty,omitempty"`
}

type ProductResponse struct {
	Serial          string      `json:"serial,omitempty"`
	Status          bool        `json:"status,omitempty"`
	Active          bool        `json:"active,omitempty"`
	ProductInfo     ProductInfo `json:"productInfo,omitempty"`
	Production      time.Time   `json:"production,omitempty"`
	DefaultWarranty time.Time   `json:"defaultWarranty,omitempty"`
	EWarranty       EWarranty   `json:"EWarranty,omitempty"`
}

type ProductWarranty struct {
	EWarranty  string `json:"EWarranty" binding:"required"`
	ActiveDate string `json:"activeDate" binding:"required"`
}

type EWarranty struct {
	EWarranty  time.Time `json:"EWarranty"`
	ActiveDate time.Time `json:"activeDate"`
}

type ProductService interface {
	GetProduct(serial string) (*ProductResponse, error)
	GetProducts() ([]*ProductResponse, error)
	CreateProduct(product *ProductNew) (*ProductResponse, error)
	UpdateProduct(serial string, productInfo *ProductInfo) (*ProductResponse, error)
	UpdateEWarranty(serial string, warranty *EWarranty) (*ProductResponse, error)
	DeleteProduct(serial string) error
}
