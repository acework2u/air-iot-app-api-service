package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AirRepositoryDB struct {
	airCollection *mongo.Collection
	ctx           context.Context
}

func NewAirRepository(ctx context.Context, airCollection *mongo.Collection) AirRepository {
	return &AirRepositoryDB{
		ctx:           ctx,
		airCollection: airCollection,
	}
}

func (r *AirRepositoryDB) RegisterAir(info *AirInfo) (*DBAirInfo, error) {

	airInfo := (*AirInfo)(info)
	now := time.Now()
	airInfo.RegisterDate = now.Local()
	airInfo.UpdatedDate = airInfo.RegisterDate

	check, _ := r.checkDuplicate(airInfo.Serial, airInfo.UserId)
	if check > 0 {
		return nil, errors.New("Your product is a duplicate.")
	}
	res, err := r.airCollection.InsertOne(r.ctx, airInfo)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("Serial No. already exists.")
		}
		return nil, err
	}
	newDevice := &DBAirInfo{}
	query := bson.M{"_id": res.InsertedID}
	if err = r.airCollection.FindOne(r.ctx, query).Decode(newDevice); err != nil {
		return nil, err
	}

	return newDevice, nil

}
func (r *AirRepositoryDB) UpdateAir(info *AirInfo) (*DBAirInfo, error) {
	return nil, nil
}
func (r *AirRepositoryDB) Airs(userId string) ([]*DBAirInfo, error) {

	filter := bson.M{"userId": userId}
	cursor, err := r.airCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	airs := []*DBAirInfo{}
	for cursor.Next(r.ctx) {
		air := &DBAirInfo{}
		err := cursor.Decode(air)
		if err != nil {
			return nil, err
		}
		airs = append(airs, air)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(airs) == 0 {
		return []*DBAirInfo{}, nil
	}
	return airs, nil
}
func (r *AirRepositoryDB) checkDuplicate(serial string, userId string) (int64, error) {

	cursor, err := r.airCollection.CountDocuments(r.ctx, bson.M{"serial": serial, "userId": userId})
	if err != nil {
		return 0, err
	}
	return cursor, nil
}
