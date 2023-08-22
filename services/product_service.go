package services

import (
	"github.com/acework2u/air-iot-app-api-service/repository"
	"time"
)

type productService struct {
	product repository.ProductRepository
}

func NewProductService(product repository.ProductRepository) ProductService {
	return &productService{product}
}
func (s *productService) CreateProduct(product *ProductNew) (*ProductResponse, error) {

	now := time.Now()
	productInfo := (repository.ProductInfo)(product.ProductInfo)
	productNew := &repository.Product{
		Serial:          product.Serial,
		Status:          true,
		Active:          false,
		ProductInfo:     productInfo,
		Production:      now,
		DefaultWarranty: now.AddDate(1, 0, 0),
	}

	newProduct, err := s.product.CreateProduct(productNew)

	if err != nil {
		return nil, err
	}

	response := &ProductResponse{
		Serial:          newProduct.Serial,
		Status:          newProduct.Status,
		Active:          newProduct.Active,
		ProductInfo:     (ProductInfo)(newProduct.ProductInfo),
		Production:      newProduct.Production,
		DefaultWarranty: newProduct.DefaultWarranty,
		EWarranty:       (EWarranty)(newProduct.EWarranty),
	}
	return response, nil
}
func (s *productService) DeleteProduct(serial string) error {

	var err error
	if len(serial) > 0 {
		err = s.product.DeleteProduct(serial)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
