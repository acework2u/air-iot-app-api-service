package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ScheduleJob struct {
	SerialNo    string    `bson:"serialNo" json:"serialNo"`
	UserId      string    `bson:"userId" json:"userId"`
	Command     []string  `bson:"command" json:"command"`
	Mode        string    `bson:"mode" json:"mode"`
	Duration    []string  `bson:"duration" json:"duration"`
	StartDate   time.Time `bson:"startDate" json:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	CreatedDate time.Time `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedDate time.Time `bson:"updatedDate,omitempty" json:"updatedDate"`
}
type ScheduleJobDB struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	SerialNo    string             `bson:"serialNo" json:"serialNo"`
	UserId      string             `bson:"userId" json:"userId"`
	Command     []string           `bson:"command" json:"command"`
	Mode        string             `bson:"mode" json:"mode"`
	Duration    []string           `bson:"duration" json:"duration"`
	StartDate   time.Time          `bson:"startDate" json:"startDate"`
	EndDate     time.Time          `json:"endDate" bson:"endDate"`
	CreatedDate time.Time          `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate,omitempty" json:"updatedDate"`
}

type JobCommand struct {
	Cmd string `bson:"cmd" json:"cmd"`
	Val string `bson:"val" json:"val"`
}

type DurationJob struct {
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `json:"endDate" bson:"endDate"`
}

type ScheduleRepository interface {
	ListJob(UserId string) ([]*ScheduleJobDB, error)
	NewJob(userId string, job *ScheduleJob) (*ScheduleJobDB, error)
	UpdateJob(jobId string) (*ScheduleJobDB, error)
	DeleteJob(jobId string) (*ScheduleJobDB, error)
}
