package services

import (
	"time"
)

type JobSchedule struct {
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

type JobUpdateSchedule struct {
	SerialNo    string    `bson:"serialNo" json:"serialNo" validate:"required" binding:"required"`
	UserId      string    `bson:"userId" json:"userId" validate:"required" binding:"required"`
	Command     []string  `bson:"command" json:"command" validate:"required" binding:"required"`
	Mode        string    `bson:"mode" json:"mode" validate:"required" binding:"required"`
	Duration    []string  `bson:"duration" json:"duration" validate:"required" binding:"required"`
	StartDate   time.Time `bson:"startDate" json:"startDate"`
	EndDate     time.Time `json:"endDate" bson:"endDate"`
	UpdatedDate time.Time `bson:"updatedDate,omitempty" json:"updatedDate,omitempty"`
}

type JobDbSchedule struct {
	Id        string    `json:"id"`
	SerialNo  string    `bson:"serialNo" json:"serialNo"`
	Command   []string  `bson:"command" json:"command"`
	Mode      string    `bson:"mode" json:"mode"`
	Duration  []string  `bson:"duration" json:"duration"`
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `json:"endDate" bson:"endDate"`
	Status    bool      `bson:"status" json:"status"`
}

type ScheduleService interface {
	GetSchedules(userId string) ([]*JobDbSchedule, error)
	NewJobSchedules(userId string, jobInfo *JobSchedule) (*JobDbSchedule, error)
}
