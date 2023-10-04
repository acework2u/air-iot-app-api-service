package repository

import (
	"context"
	"errors"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ScheduleRepositoryDB struct {
	scheduleCollection *mongo.Collection
	ctx                context.Context
}

func NewScheduleRepository(ctx context.Context, scheduleCollection *mongo.Collection) ScheduleRepository {

	return &ScheduleRepositoryDB{
		ctx:                ctx,
		scheduleCollection: scheduleCollection,
	}
}

func (r *ScheduleRepositoryDB) ListJob(UserId string) ([]*ScheduleJobDB, error) {
	filter := bson.M{"userId": UserId}
	cursor, err := r.scheduleCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	jobs := []*ScheduleJobDB{}

	for cursor.Next(r.ctx) {
		job := &ScheduleJobDB{}
		err := cursor.Decode(job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)

	}
	if err := cursor.Err(); err != nil {
		if len(jobs) == 0 {
			return nil, err
		}
	}
	return jobs, err

}
func (r *ScheduleRepositoryDB) NewJob(userId string, job *ScheduleJob) (*ScheduleJobDB, error) {
	now := time.Now()
	jobInfo := (*ScheduleJob)(job)
	_ = jobInfo

	jobInfo.CreatedDate = now.Local()
	jobInfo.UpdatedDate = jobInfo.CreatedDate

	// Insert
	res, err := r.scheduleCollection.InsertOne(r.ctx, jobInfo)
	if err != nil {
		return nil, err
	}
	newJob := &ScheduleJobDB{}
	query := bson.M{"_id": res.InsertedID}
	if err = r.scheduleCollection.FindOne(r.ctx, query).Decode(newJob); err != nil {
		return nil, err
	}

	return newJob, nil
}
func (r *ScheduleRepositoryDB) UpdateJob(jobId string, updateInfo *ScheduleJobUpdate) (*ScheduleJobDB, error) {

	objId, _ := primitive.ObjectIDFromHex(jobId)
	doc, err := utils.ToDoc(updateInfo)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{{Key: "$set", Value: doc}}

	res := r.scheduleCollection.FindOneAndUpdate(r.ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	jobInfo := &ScheduleJobDB{}
	if err := res.Decode(jobInfo); err != nil {
		return nil, errors.New("no job with that Id exists")
	}
	return jobInfo, err

	return nil, nil
}
func (r *ScheduleRepositoryDB) DeleteJob(jobId string) error {
	id, _ := primitive.ObjectIDFromHex(jobId)
	query := bson.M{"_id": id}
	resDel, err := r.scheduleCollection.DeleteOne(r.ctx, query)
	if err != nil {
		return err
	}
	if resDel.DeletedCount == 0 {
		return errors.New("no job with that Id exists")
	}
	return nil

}
func (r *ScheduleRepositoryDB) JobsSchedule() ([]*ScheduleJobDB, error) {
	filter := bson.M{"status": true}
	cursor, err := r.scheduleCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)
	jobs := []*ScheduleJobDB{}

	for cursor.Next(r.ctx) {
		job := &ScheduleJobDB{}
		err := cursor.Decode(job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)

	}
	if err := cursor.Err(); err != nil {
		if len(jobs) == 0 {
			return nil, err
		}
	}
	return jobs, err

}
