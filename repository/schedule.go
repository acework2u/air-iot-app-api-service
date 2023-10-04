package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ScheduleJob struct {
	SerialNo    string    `bson:"serialNo" json:"serialNo"`
	UserId      string    `bson:"userId" json:"userId"`
	Command     []AirCmd  `bson:"command" json:"command"`
	Mode        string    `bson:"mode" json:"mode"`
	Duration    []string  `bson:"duration" json:"duration"`
	StartDate   time.Time `bson:"startDate" json:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	Status      bool      `bson:"status" json:"status"`
	CreatedDate time.Time `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedDate time.Time `bson:"updatedDate,omitempty" json:"updatedDate"`
}
type ScheduleJobDB struct {
	Id          primitive.ObjectID `bson:"_id" json:"id"`
	SerialNo    string             `bson:"serialNo" json:"serialNo"`
	UserId      string             `bson:"userId" json:"userId"`
	Command     []AirCmd           `bson:"command" json:"command"`
	Mode        string             `bson:"mode" json:"mode"`
	Duration    []string           `bson:"duration" json:"duration"`
	StartDate   time.Time          `bson:"startDate" json:"startDate"`
	EndDate     time.Time          `json:"endDate" bson:"endDate"`
	Status      bool               `bson:"status" json:"status"`
	CreatedDate time.Time          `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedDate time.Time          `bson:"updatedDate,omitempty" json:"updatedDate"`
}
type ScheduleJobUpdate struct {
	SerialNo    string    `bson:"serialNo" json:"serialNo"`
	UserId      string    `bson:"userId" json:"userId"`
	Command     []string  `bson:"command" json:"command"`
	Mode        string    `bson:"mode" json:"mode"`
	Duration    []string  `bson:"duration" json:"duration"`
	StartDate   time.Time `bson:"startDate" json:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	Status      bool      `bson:"status" json:"status"`
	UpdatedDate time.Time `bson:"updatedDate,omitempty" json:"updatedDate"`
}

type JobCommand struct {
	Cmd   string `bson:"cmd" json:"cmd"`
	Value string `bson:"value" json:"value"`
}

type AirCmd struct {
	Cmd   string `json:"cmd"`
	Value string `json:"value"`
}

type DurationJob struct {
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `json:"endDate" bson:"endDate"`
}

type ScheduleRepository interface {
	ListJob(UserId string) ([]*ScheduleJobDB, error)
	NewJob(userId string, job *ScheduleJob) (*ScheduleJobDB, error)
	UpdateJob(jobId string, update *ScheduleJobUpdate) (*ScheduleJobDB, error)
	DeleteJob(jobId string) error
	JobsSchedule() ([]*ScheduleJobDB, error)
}
