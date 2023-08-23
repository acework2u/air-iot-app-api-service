package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ProductRepositoryDB struct {
	productCollection *mongo.Collection
	ctx               context.Context
	client            *mongo.Client
}

func NewProductRepositoryDB(ctx context.Context, productCollection *mongo.Collection) ProductRepository {
	return &ProductRepositoryDB{ctx: ctx, productCollection: productCollection}
}
func (r *ProductRepositoryDB) GetProduct(serial string) (*DBProduct, error) {

	query := bson.M{"serial": serial}
	productInfo := &DBProduct{}
	if err := r.productCollection.FindOne(r.ctx, query).Decode(productInfo); err != nil {
		return nil, err
	}
	return productInfo, nil

}
func (r *ProductRepositoryDB) CreateProduct(product *Product) (*DBProduct, error) {
	now := time.Now()
	product.Production = now
	product.DefaultWarranty = now.AddDate(1, 0, 0)

	res, err := r.productCollection.InsertOne(r.ctx, product)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("product already exits")
		}
		return nil, err
	}
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"serial": 1}, Options: opt}
	if _, err := r.productCollection.Indexes().CreateOne(r.ctx, index); err != nil {
		return nil, err
	}

	newCustomer := &DBProduct{}
	query := bson.M{"_id": res.InsertedID}

	if err = r.productCollection.FindOne(r.ctx, query).Decode(newCustomer); err != nil {
		return nil, err
	}

	return newCustomer, nil
}
func (r *ProductRepositoryDB) DeleteProduct(serial string) error {

	query := bson.M{"serial": serial}

	delProduct, err := r.productCollection.DeleteOne(r.ctx, query)
	if err != nil {
		return err
	}
	if delProduct.DeletedCount == 0 {
		return errors.New("no product with that serial exists")
	}
	return nil
}
