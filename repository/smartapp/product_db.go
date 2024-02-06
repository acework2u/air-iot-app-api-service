package smartapp

import (
	"context"
	"gorm.io/gorm"
)

type productsRepo struct {
	db  *gorm.DB
	ctx context.Context
}

func NewProductRepository(ctx context.Context, db *gorm.DB) ProductRepository {
	return &productsRepo{db: db, ctx: ctx}
}

func (r *productsRepo) NewProduct(info ProductInfo) (*ProductInfo, error) {
	return nil, nil
}
func (r *productsRepo) UpdateProduct(id int, info ProductInfo) (bool, error) {
	return false, nil
}
func (r *productsRepo) DeleteProduct(id int) error {
	return nil
}
func (r *productsRepo) RegisterWarranty(productWarrantyInfo ProductWarranty) (*ProductWarranty, error) {
	return nil, nil
}
func (r *productsRepo) ProductsWarranty(serialNo string) (*ProductWarranty, error) {
	return nil, nil
}
