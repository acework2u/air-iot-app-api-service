package smartapp

type ProductInfo struct {
}

type Products struct {
}

type ProductWarranty struct {
}

type ProductRepository interface {
	NewProduct(info ProductInfo) (*ProductInfo, error)
	UpdateProduct(id int, info ProductInfo) (bool, error)
	DeleteProduct(id int) error
	RegisterWarranty(productWarrantyInfo ProductWarranty) (*ProductWarranty, error)
	ProductsWarranty(serialNo string) (*ProductWarranty, error)
}
