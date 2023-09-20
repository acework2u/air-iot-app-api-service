package services

import (
	"context"
	"fmt"
)

type jobsService struct {
	Ctx context.Context
}

func NewJobsService() JobsService {
	return &jobsService{}
}

func (s *jobsService) JobsThingsHandler(userId string) (interface{}, error) {
	client, err := NewAwsMqttConnect(userId)

	if err != nil {
		panic(err)
	}
	defer client.Disconnect()
	fmt.Println("is Connected")

	dataResponse := "is Connected"
	return dataResponse, nil
}
