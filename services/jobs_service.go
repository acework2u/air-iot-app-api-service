package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type jobsService struct {
	Ctx       context.Context
	airCfg    *AirThingConfig
	Cfg       *aws.Config
	IotClient *iot.Client
}

func NewJobsService(airCfg *AirThingConfig) JobsService {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(airCfg.Region), config.WithSharedConfigProfile("default"))

	if err != nil {
		panic(err)
	}

	iotClient := iot.NewFromConfig(cfg)
	return &jobsService{
		Cfg:       &cfg,
		IotClient: iotClient,
		Ctx:       context.TODO()}
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

			fmt.Println("Work in job")

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
func (s *jobsService) CreateJobsThings(deviceSn string) (interface{}, error) {

	jobInput := &iot.DescribeJobInput{
		JobId: aws.String("airConfig-v2"),
	}
	destOut, err := s.IotClient.DescribeJob(s.Ctx, jobInput)
	if err != nil {
		return nil, err
	}

	return destOut, nil
}
func (s *jobsService) GetJobsThings(deviceSn string) (interface{}, error) {

	jobInput := &iot.ListJobExecutionsForThingInput{ThingName: aws.String(deviceSn)}
	jobOut, err := s.IotClient.ListJobExecutionsForThing(s.Ctx, jobInput)
	if err != nil {
		return nil, err
	}

	return jobOut, nil
}
func (s *jobsService) GetQueJobsThings(device string) (interface{}, error) {
	jobId := "airConfig-v2"

	jobInput := &iot.DescribeJobExecutionInput{
		JobId:     aws.String(jobId),
		ThingName: aws.String(device),
	}
	//DescribeJobExecution
	jobsOut, err := s.IotClient.DescribeJobExecution(s.Ctx, jobInput)
	if err != nil {
		return nil, err
	}

	return jobsOut, nil
}
func (s *jobsService) UpdateJobsThings() (interface{}, error) {

	//s.IotClient.UpdateJob(s.Ctx, &iot.UpdateJobExecution{})
	//
	return nil, nil
}
