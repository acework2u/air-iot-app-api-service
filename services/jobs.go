package services

import "time"

type JobsAccept struct {
	Timestamp int `json:"timestamp"`
	Execution struct {
		JobID           string `json:"jobId"`
		Status          string `json:"status"`
		QueuedAt        int    `json:"queuedAt"`
		LastUpdatedAt   int    `json:"lastUpdatedAt"`
		VersionNumber   int    `json:"versionNumber"`
		ExecutionNumber int    `json:"executionNumber"`
		JobDocument     struct {
			C string `json:"c"`
		} `json:"jobDocument"`
	} `json:"execution"`
}

type JobsStatus struct {
	Status string `json:"status"`
}
type DeviceReq struct {
	DeviceSn string `uri:"deviceSn" json:"deviceSn"`
}
type JobsUpdateReq struct {
	DeviceSn  string `json:"deviceSn"`
	JobId     string `json:"jobId"`
	JobStatus string `json:"jobStatus"`
}
type JobsReq struct {
	DeviceSn string `json:"deviceSn"`
	UserId   string `json:"userId"`
}

type LisJobsThing struct {
	ExecutionSummaries []struct {
		JobExecutionSummary struct {
			ExecutionNumber int       `json:"ExecutionNumber"`
			LastUpdatedAt   time.Time `json:"LastUpdatedAt"`
			QueuedAt        time.Time `json:"QueuedAt"`
			RetryAttempt    int       `json:"RetryAttempt"`
			StartedAt       any       `json:"StartedAt"`
			Status          string    `json:"Status"`
		} `json:"JobExecutionSummary"`
		JobID string `json:"JobId"`
	} `json:"ExecutionSummaries"`
	NextToken      any `json:"NextToken"`
	ResultMetadata struct {
	} `json:"ResultMetadata"`
}

type JobsService interface {
	JobsThingsHandler(userId string) (interface{}, error)
	CreateJobsThings(sn string) (interface{}, error)
	GetJobsThings(device string) (interface{}, error)
	GetQueJobsThings(device string) (interface{}, error)
	UpdateJobsThings() (interface{}, error)
}
