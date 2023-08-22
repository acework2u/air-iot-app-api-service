package services

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

type ProductService interface {
	CreateProduct(product *ProductNew) (*ProductInfo, error)
}
