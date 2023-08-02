package repository

import (
	"context"
	"errors"
	"time"

	"github.com/acework2u/air-iot-app-api-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CustomerRepositoryDB struct {
	cusCollection *mongo.Collection
	ctx           context.Context
}

func NewCustomerRepositoryDB(cusCollection *mongo.Collection, ctx context.Context) CustomerRepository {
	return &CustomerRepositoryDB{cusCollection, ctx}
}

func (r *CustomerRepositoryDB) CreateCustomer(customer *CreateCustomerRequest) (*DBCustomer, error) {

	customer.CreateAt = time.Now()
	customer.UpdateAt = customer.CreateAt

	res, err := r.cusCollection.InsertOne(r.ctx, customer)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {

			return nil, errors.New("name already exits")
			// return nil, er
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}
	if _, err := r.cusCollection.Indexes().CreateOne(r.ctx, index); err != nil {
		return nil, errors.New("could not create index for name")
	}
	// index2 := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	// if _, err := r.cusCollection.Indexes().CreateOne(r.ctx, index2); err != nil {
	// 	return nil, errors.New("could not create index for email")
	// }

	var newCustomer *DBCustomer

	query := bson.M{"_id": res.InsertedID}
	if err = r.cusCollection.FindOne(r.ctx, query).Decode(&newCustomer); err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (r *CustomerRepositoryDB) NewCustomer(customer *CreateCustomerRequest2) (*DBCustomer2, error) {
	var newCustomer *DBCustomer2

	customer.CreateAt = time.Now()
	customer.UpdateAt = customer.CreateAt

	res, err := r.cusCollection.InsertOne(r.ctx, customer)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {

			return nil, errors.New("name already exits")

		}
		return nil, err
	}

	//opt := options.Index()
	//opt.SetUnique(true)
	//
	//index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}
	//if _, err := r.cusCollection.Indexes().CreateOne(r.ctx, index); err != nil {
	//	return nil, errors.New("could not create index for name")
	//}

	query := bson.M{"_id": res.InsertedID}
	if err = r.cusCollection.FindOne(r.ctx, query).Decode(&newCustomer); err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (r *CustomerRepositoryDB) UpdateCustomer(id string, data *UpdateCustomer) (*DBCustomer, error) {

	doc, err := utils.ToDoc(data)

	if err != nil {
		return nil, err
	}

	//obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.D{{Key: "usersub", Value: id}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := r.cusCollection.FindOneAndUpdate(r.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatCustomer *DBCustomer
	if err := res.Decode(&updatCustomer); err != nil {
		return nil, errors.New("no customer with that Id exists")
	}

	return updatCustomer, nil

}
func (r *CustomerRepositoryDB) FindCustomerById(id string) (*DBCustomer, error) {

	//query := bson.M{"usersub": id}
	//var customers *DBCustomer
	//
	//cursor, err := r.cusCollection.Find(r.ctx, query)
	//fmt.Println("FindCustomer")
	//fmt.Println(err)
	//
	//defer cursor.Close(r.ctx)
	//if err != nil {
	//	if err == mongo.ErrNoDocuments {
	//		return nil, err
	//	}
	//	return nil, err
	//}
	//cursor.Decode(&customers)

	return nil, nil

}
func (r *CustomerRepositoryDB) FindCustomers() ([]*DBCustomer, error) {

	query := bson.M{}

	cursor, err := r.cusCollection.Find(r.ctx, query)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(r.ctx)

	var customers []*DBCustomer

	for cursor.Next(r.ctx) {
		customer := &DBCustomer{}
		err := cursor.Decode(customer)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(customers) == 0 {
		return []*DBCustomer{}, nil
	}

	return customers, nil
}
func (r *CustomerRepositoryDB) DeleteCustomer(id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": objId}

	res, err := r.cusCollection.DeleteOne(r.ctx, query)
	if err != nil {
		return nil
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
func (r *CustomerRepositoryDB) FindCustomerID(uid string) (*DBCustomer2, error) {

	query := bson.M{"usersub": uid}
	var result *DBCustomer2
	err := r.cusCollection.FindOne(r.ctx, query).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
