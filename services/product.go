package services

import "time"

type ProductNew struct {
	Serial          string      `json:"serial",binding:"required"`
	ProductInfo     ProductInfo `json:"productInfo" bson:"productInfo"`
	Production      time.Time   `json:"production,omitempty"`
	DefaultWarranty time.Time   `json:"defaultWarranty,omitempty"`
}

type ProductInfo struct {
	Title        string `json:"title",binding:"required"`
	Model        string `json:"model"`
	Sku          string `json:"sku"`
	Mpn          string `json:"mpn"`
	ProductImage string `json:"productImage"`
}
