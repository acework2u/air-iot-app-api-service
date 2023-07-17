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
func (r *deviceRepositoryDB) FindDevices(request *DeviceRequest) ([]*DBDevice, error) {

	query := bson.M{"userId": request.UserId}
	cursor, err := r.devicesCollection.Find(r.ctx, query)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var devices []*DBDevice

	for cursor.Next(r.ctx) {
		device := &DBDevice{}
		err := cursor.Decode(device)

		if err != nil {
			return nil, err
		}

		devices = append(devices, device)

	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		return []*DBDevice{}, nil
	}

	return devices, err

}
