package repository

import (
	"context"
	"errors"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
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
<<<<<<< HEAD
func (r *ProductRepositoryDB) GetProducts() ([]*DBProduct, error) {
	filter := bson.D{{}}
	return r.filterProduct(filter)
}
=======
>>>>>>> ad1f98be097d983c078b0925f74ee2be200245ae
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
func (r *ProductRepositoryDB) UpdateProduct(serial string, product *Product) (*DBProduct, error) {

	doc, err := utils.ToDoc(product)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "serial", Value: serial}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := r.productCollection.FindOneAndUpdate(r.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	productInfo := &DBProduct{}
	if err = res.Decode(productInfo); err != nil {
		return nil, err
	}

	return productInfo, nil
}
func (r *ProductRepositoryDB) UpdateProductInfo(serial string, productInfo *DBProductInfoUpdate) (*DBProduct, error) {

	productUp := (*DBProductInfoUpdate)(productInfo)
	doc, err := utils.ToDoc(productUp)
	if err != nil {
		return nil, err
	}

	query := bson.D{{Key: "serial", Value: serial}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := r.productCollection.FindOneAndUpdate(r.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	product := &DBProduct{}
	if err := res.Decode(product); err != nil {
		return nil, errors.New("no product with serial exists")
	}

	return product, nil
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
func (r *ProductRepositoryDB) UpdateEWarranty(serial string) (*DBProductInfoUpdate, error) {

	chkActive, err := r.checkActive(serial)
	if err != nil {
		return nil, err
	}
	if chkActive {
		return nil, errors.New("This product has previously been registered for warranty.")
	}

	now := time.Now()
	activeDate := now.Local()
	ewarranty := activeDate.AddDate(1, 0, 0)

	productWarranty := EWarranty{EWarranty: ewarranty, ActiveDate: activeDate}

	productInfo := &DBProductInfoUpdate{
		Serial:    serial,
		Active:    true,
		EWarranty: productWarranty,
	}

	doc, err := utils.ToDoc(productInfo)
	if err != nil {
		return nil, err
	}
	query := bson.D{{Key: "serial", Value: serial}}
	update := bson.D{{Key: "$set", Value: doc}}

	_, err = r.productCollection.UpdateOne(r.ctx, query, update)
	if err != nil {
		return nil, err
	}

	return productInfo, nil
}
func (r *ProductRepositoryDB) filterProduct(filter interface{}) ([]*DBProduct, error) {

	products := []*DBProduct{}
	//var products []*DBProduct
	cur, err := r.productCollection.Find(r.ctx, filter)
	if err != nil {
		return products, err
	}
	for cur.Next(r.ctx) {
		t := DBProduct{}
		err := cur.Decode(&t)
		if err != nil {
			return products, err
		}
		products = append(products, &t)
	}
	if err := cur.Err(); err != nil {
		return products, err
	}
	cur.Close(r.ctx)

	if len(products) == 0 {
		return products, mongo.ErrNoDocuments
	}
	return products, nil
}
func (r *ProductRepositoryDB) checkActive(serial string) (bool, error) {
	result := &DBProduct{}
	query := bson.D{{Key: "serial", Value: strings.ToUpper(serial)}}
	err := r.productCollection.FindOne(r.ctx, query).Decode(result)
	if err != nil {
		return false, err
	}
	return result.Active, nil
}
