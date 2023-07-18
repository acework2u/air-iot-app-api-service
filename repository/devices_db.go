package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
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

	check, _ := r.CheckDupDevice(device.UserId, device.SerialNo)

	if check > 0 {

		err := errors.New("Your device is a duplicate.")

		return nil, err

	}

	res, err := r.devicesCollection.InsertOne(r.ctx, device)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Serial No. already exists.")
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
func (r *deviceRepositoryDB) CheckDupDevice(userId string, serialNo string) (int64, error) {
	cursor, err := r.devicesCollection.CountDocuments(r.ctx, bson.M{"userId": userId, "serialNo": strings.ToUpper(serialNo)})
	if err != nil {
		return 0, err
	}
	return cursor, nil
}
func (r *deviceRepositoryDB) UpdateDevice(userid string, device *DeviceUpdateReq) (*DBDevice, error) {

	objId, _ := primitive.ObjectIDFromHex(userid)

	doc, err := utils.ToDoc(device)
	if err != nil {
		return nil, err
	}
	query := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{{Key: "$set", Value: doc}}

	fmt.Println("Working in Repo")
	fmt.Println(userid)
	fmt.Println(query)

	res := r.devicesCollection.FindOneAndUpdate(r.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var deviceInfo *DBDevice
	if err := res.Decode(&deviceInfo); err != nil {
		return nil, errors.New("no device with that Id exists")
	}
	return deviceInfo, nil
}
func (r *deviceRepositoryDB) DeleteDevice(filter *DeviceFilter) (bool, error) {

	userID := filter.UserId
	objId, _ := primitive.ObjectIDFromHex(filter.Id)

	query := bson.M{"_id": objId, "userId": userID}

	resDel, err := r.devicesCollection.DeleteOne(r.ctx, query)
	if err != nil {
		return false, err
	}
	if resDel.DeletedCount == 0 {
		return false, errors.New("no document with that Id exists")
	}
	return true, nil
}
