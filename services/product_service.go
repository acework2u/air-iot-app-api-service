package services

import (
	"github.com/acework2u/air-iot-app-api-service/repository"
)

type productService struct {
	product repository.ProductRepository
}

func NewProductService(product repository.ProductRepository) ProductService {
	return &productService{product}
}
func (s *productService) CreateProduct(product *ProductNew) (*ProductInfo, error) {

	productInfo := (ProductInfo)(product.ProductInfo)

	return &productInfo, nil
}
