package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	return nil, nil
}

func (r *ScheduleRepositoryDB) NewJob(userId string, job *ScheduleJob) (*ScheduleJobDB, error) {
	now := time.Now()
	jobInfo := (*ScheduleJob)(job)
	_ = jobInfo

	jobInfo.CreatedDate = now.Local()
	jobInfo.UpdatedDate = jobInfo.CreatedDate

	return nil, nil
}
func (r *ScheduleRepositoryDB) UpdateJob(jobId string) (*ScheduleJobDB, error) {
	return nil, nil
}
func (r *ScheduleRepositoryDB) DeleteJob(jobId string) (*ScheduleJobDB, error) {
	return nil, nil
}
