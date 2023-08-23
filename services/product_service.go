package services

import (
	"errors"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"time"
)

type productService struct {
	product repository.ProductRepository
}

func NewProductService(product repository.ProductRepository) ProductService {
	return &productService{product}
}

func (s *productService) GetProduct(serial string) (*ProductResponse, error) {
	if len(serial) < 0 {
		return nil, errors.New("serial is wrong")
	}
	res, err := s.product.GetProduct(serial)
	if err != nil {
		return nil, err
	}
	product := &ProductResponse{
		Serial:          res.Serial,
		Status:          res.Status,
		Active:          res.Active,
		Production:      res.Production,
		DefaultWarranty: res.DefaultWarranty,
		ProductInfo:     (ProductInfo)(res.ProductInfo),
		EWarranty:       (EWarranty)(res.EWarranty),
	}

	return product, err

}
func (s *productService) GetProducts() ([]*ProductResponse, error) {

	products := []*ProductResponse{}

	res, err := s.product.GetProducts()
	if err != nil {
		return nil, err
	}

	for _, item := range res {
		product := &ProductResponse{
			Serial:          item.Serial,
			Status:          item.Status,
			Active:          item.Active,
			ProductInfo:     (ProductInfo)(item.ProductInfo),
			Production:      item.Production,
			DefaultWarranty: item.DefaultWarranty,
			EWarranty:       (EWarranty)(item.EWarranty),
		}

		products = append(products, product)
	}

	return products, nil
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
func (s *productService) UpdateProduct(serial string, productInfo *ProductInfo) (*ProductResponse, error) {

	proInfo := (*repository.ProductInfo)(productInfo)
	updateProduct := &repository.DBProductInfoUpdate{
		ProductInfo: repository.ProductInfo{
			Title:        proInfo.Title,
			Model:        productInfo.Model,
			Sku:          proInfo.Sku,
			Mpn:          proInfo.Mpn,
			ProductImage: proInfo.ProductImage,
		},
	}

	dbProduct, err := s.product.UpdateProductInfo(serial, updateProduct)
	if err != nil {
		return nil, err
	}
	proRes := &ProductResponse{
		Serial:          dbProduct.Serial,
		Active:          dbProduct.Active,
		Status:          dbProduct.Status,
		ProductInfo:     (ProductInfo)(dbProduct.ProductInfo),
		EWarranty:       (EWarranty)(dbProduct.EWarranty),
		Production:      dbProduct.Production,
		DefaultWarranty: dbProduct.DefaultWarranty,
	}

	return proRes, nil
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
