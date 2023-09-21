package services

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

type JobsReq struct {
	DeviceSn string `json:"deviceSn"`
	UserId   string `json:"userId"`
}

type JobsService interface {
	JobsThingsHandler(userId string) (interface{}, error)
	CreateJobs(sn string) (interface{}, error)
}
