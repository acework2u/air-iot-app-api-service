package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//var ctx context.Context

type AddressRepositoryDB struct {
	addrCollection *mongo.Collection
	ctx            context.Context
}

func NewAddressRepositoryDB(addrCollection *mongo.Collection, ctx context.Context) AddressRepository {
	return &AddressRepositoryDB{addrCollection, ctx}
}

func (r *AddressRepositoryDB) CreateNewAddress(address *CustomerAddress) (*DBAddress, error) {

	address.UpdateAt = time.Now()

	res, err := r.addrCollection.InsertOne(r.ctx, address)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("address already exits")
		}
		return nil, err
	}

	var newAddress *DBAddress

	query := bson.M{"_id": res.InsertedID}
	if err = r.addrCollection.FindOne(r.ctx, query).Decode(&newAddress); err != nil {
		return nil, err
	}

	return newAddress, nil
}
func (r *AddressRepositoryDB) UpdateAddress(id string, address *UpdateCustomer) (*DBAddress, error) {
	return nil, nil
}
func (r *AddressRepositoryDB) DeleteAddress(id string) error {
	return nil
}
func (r *AddressRepositoryDB) FindAddress() ([]*DBAddress, error) {
	return nil, nil
}
func (r *AddressRepositoryDB) FindAddressId(string) (*DBAddress, error) {
	return nil, nil
}
