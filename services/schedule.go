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
	Status      bool      `bson:"status" json:"status"`
	CreatedDate time.Time `bson:"createdDate,omitempty" json:"createdDate"`
	UpdatedDate time.Time `bson:"updatedDate,omitempty" json:"updatedDate"`
}
type JobScheduleReq struct {
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
	Command   []AirCmd  `bson:"command" json:"command"`
	Mode      string    `bson:"mode" json:"mode"`
	Duration  []string  `bson:"duration" json:"duration"`
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `json:"endDate" bson:"endDate"`
	Status    bool      `bson:"status" json:"status"`
}

type JobWork struct {
	SerialNo  string    `bson:"serialNo" json:"serialNo"`
	Command   []AirCmd  `bson:"command" json:"command"`
	Mode      string    `bson:"mode" json:"mode"`
	Duration  []string  `bson:"duration" json:"duration"`
	Status    bool      `bson:"status" json:"status"`
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `json:"endDate" bson:"endDate"`
}

type AirCmd struct {
	Cmd   string `json:"cmd"`
	Value string `json:"value"`
}

type AirJob struct {
	SerialNo string   `bson:"serialNo" json:"serialNo"`
	Command  []AirCmd `bson:"command" json:"command"`
}

type ScheduleService interface {
	GetSchedules(userId string) ([]*JobDbSchedule, error)
	NewJobSchedules(userId string, jobInfo *JobScheduleReq) (*JobDbSchedule, error)
	CornJob()
}
