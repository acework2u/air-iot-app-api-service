package services

import "github.com/acework2u/air-iot-app-api-service/repository"

type productService struct {
	product repository.ProductRepositoryDB
}

func NewProductService(product repository.ProductRepositoryDB) ProductService {
	return &productService{product: product}
}
