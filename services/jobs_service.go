package services

import (
	"context"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type jobsService struct {
	Ctx context.Context
}

func NewJobsService() JobsService {
	return &jobsService{}
}

func (s *jobsService) JobsThingsHandler(deviceSn string) (interface{}, error) {

	deviceSn = "2300F15050017"
	//deviceSn = "2300F15050023"
	clientMqtt, err := NewAwsMqttConnect(deviceSn)
	//clientMqtt, err := NewAwsMqttConnect(userId)

	if err != nil {
		panic(err)
	}
	//defer client.Disconnect()
	//
	//jobTopic := "$aws/things/2300F15050017/jobs/notify"
	//jobTopic := "$aws/things/2300F15050017/jobs/notify-next"
	fmt.Sprintf("$aws/things/%v/jobs/notify-next", deviceSn)
	jobTopic := fmt.Sprintf("$aws/things/%v/jobs/notify-next", deviceSn)

	fmt.Println("Work Topic")
	fmt.Println(jobTopic)

	//jobTopic := "$aws/things/2300F15050017/shadow/name/air-users/update/accepted"
	shadowsVal := &JobsAccept{}
	result := make(chan *JobsAccept)

	go func(result chan<- *JobsAccept) {

		shadVal := &JobsAccept{}
		fmt.Println("Work in go")
		err := clientMqtt.SubscribeWithHandler(jobTopic, 1, func(client MQTT.Client, message MQTT.Message) {
			//msgPayload := fmt.Sprintf("%v", string(message.Payload()))

			//fmt.Println(message)
			//fmt.Sprintf("%v", string(message.Payload()))
			fmt.Println("Work in job")
			//fmt.Sprintf("%v", string(message.Payload()))

			ok := json.Unmarshal(message.Payload(), &shadVal)
			if ok != nil {
				fmt.Println("err")
				fmt.Println(err)
			}
			fmt.Println("shadowsVal")
			fmt.Sprintf("%v", shadVal)
			fmt.Println(shadVal)
			result <- shadVal
			if len(shadVal.Execution.JobID) > 0 {
				upStatus := JobsStatus{}
				upStatus.Status = "IN_PROGRESS"
				data, _ := json.Marshal(upStatus)
				jobsPubTopic := fmt.Sprintf("$aws/things/%v/jobs/%v/update", deviceSn, shadVal.Execution.JobID)

				ok := clientMqtt.Publish(jobsPubTopic, data, 0)
				if ok != nil {
					fmt.Println("PUB FAILED")
				}

				fmt.Println("PUB SUCCEEDED")

				upStatus.Status = "SUCCEEDED"
				data, _ = json.Marshal(upStatus)
				ok = clientMqtt.Publish(jobsPubTopic, data, 0)
				if ok != nil {
					fmt.Println("PUB FAILED")
				}

			}

			//SUCCEEDED

			//FAILED

		})
		if err != nil {
			fmt.Println(err)
		}
	}(result)

	//time.Sleep(4 * time.Second)

	fmt.Println("Out Scope")

	fmt.Println("PUB SUCCEEDED")
	shadowsVal = <-result
	fmt.Println(<-result)
	fmt.Println(shadowsVal)

	fmt.Println("is Connected")

	return shadowsVal, nil
}

func (s *jobsService) CreateJobs(deviceSn string) (interface{}, error) {

	return nil, nil
}
