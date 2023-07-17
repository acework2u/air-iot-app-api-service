package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type deviceRepositoryDB struct {
	devicesCollection *mongo.Collection
	ctx               context.Context
}

func NewDeviceRepositoryDB(ctx context.Context, devicesCollection *mongo.Collection) DevicesRepository {

	return &deviceRepositoryDB{
		ctx: ctx, devicesCollection: devicesCollection,
	}
}

func (r *deviceRepositoryDB) CreateDevice(device *Device) (*DBDevice, error) {

	device.CreatedAt = time.Now()
	device.UpdatedAt = device.CreatedAt

	res, err := r.devicesCollection.InsertOne(r.ctx, device)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Serial No already exits")
		}
		return nil, err
	}

	var newDevice *DBDevice
	query := bson.M{"_id": res.InsertedID}
	if err = r.devicesCollection.FindOne(r.ctx, query).Decode(&newDevice); err != nil {
		return nil, err
	}

	return newDevice, nil
}
func (r *deviceRepositoryDB) FindDevices() ([]*DBDevice, error) {
	return nil, nil
}
