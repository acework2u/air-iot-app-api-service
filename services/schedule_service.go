package services

import (
	"errors"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/robfig/cron/v3"
	"strings"
	"time"
)

type scheduleService struct {
	scheduleRepo repository.ScheduleRepository
	thingServ    ThinksService
	airCfg       *AirThingConfig
}
type airCmdReq struct {
	SerialNo string `json:"serialNo" validate:"required" binding:"required"`
	Cmd      string `json:"cmd" validate:"required" binding:"required"`
	Value    string `json:"value" validate:"required" binding:"required"`
}

func NewScheduleService(scheduleRepo repository.ScheduleRepository, cfg *AirThingConfig) ScheduleService {
	thingServ := NewThingClient(cfg.Region, cfg.UserPoolId, cfg.CognitoClientId)
	return &scheduleService{scheduleRepo: scheduleRepo, thingServ: thingServ}
}

func (s *scheduleService) GetSchedules(userId string) ([]*JobDbSchedule, error) {

	if len(userId) < 0 {
		return nil, errors.New("no data")
	}

	res, err := s.scheduleRepo.ListJob(userId)
	if err != nil {
		return nil, err
	}
	jobList := []*JobDbSchedule{}

	for _, jobs := range res {

		acCmd := []AirCmd{}

		for _, command := range jobs.Command {
			cmd := &AirCmd{
				Cmd:   command.Cmd,
				Value: command.Value,
			}
			acCmd = append(acCmd, *cmd)
		}

		job := &JobDbSchedule{
			Id:        jobs.Id.String(),
			SerialNo:  jobs.SerialNo,
			Command:   acCmd,
			Mode:      jobs.Mode,
			Duration:  jobs.Duration,
			StartDate: jobs.StartDate,
			EndDate:   jobs.EndDate,
			Status:    jobs.Status,
		}
		jobList = append(jobList, job)
	}

	return jobList, nil
}
func (s *scheduleService) NewJobSchedules(userId string, jobInfo *JobScheduleReq) (*JobDbSchedule, error) {

	//dataInfo := &repository.ScheduleJob{
	//	SerialNo:  jobInfo.SerialNo,
	//	UserId:    userId,
	//	Command:   jobInfo.Command,
	//	Mode:      jobInfo.Mode,
	//	Duration:  jobInfo.Duration,
	//	StartDate: jobInfo.StartDate,
	//	EndDate:   jobInfo.EndDate,
	//}

	acCmd := []repository.AirCmd{}
	for _, item := range jobInfo.Command {
		cmd := &repository.AirCmd{
			Cmd:   item.Cmd,
			Value: item.Value,
		}
		acCmd = append(acCmd, *cmd)
	}

	dataInfo := &repository.ScheduleJob{
		SerialNo:    jobInfo.SerialNo,
		UserId:      jobInfo.UserId,
		Command:     acCmd,
		Mode:        jobInfo.Mode,
		Duration:    jobInfo.Duration,
		Status:      jobInfo.Status,
		StartDate:   jobInfo.StartDate,
		EndDate:     jobInfo.EndDate,
		CreatedDate: jobInfo.CreatedDate,
		UpdatedDate: jobInfo.UpdatedDate,
	}

	job, err := s.scheduleRepo.NewJob(userId, dataInfo)

	if err != nil {
		return nil, err
	}

	jobs := []AirCmd{}
	for _, items := range job.Command {
		job := &AirCmd{
			Cmd:   items.Cmd,
			Value: items.Value,
		}
		jobs = append(jobs, *job)
	}

	resJob := &JobDbSchedule{
		SerialNo:  job.SerialNo,
		Command:   jobs,
		Mode:      job.Mode,
		Duration:  job.Duration,
		StartDate: job.StartDate,
		EndDate:   job.EndDate,
		Status:    job.Status,
	}

	return resJob, nil
}
func (s *scheduleService) Job(ac airCmdReq) {

	airJob := utils.NewAirCmd(ac.SerialNo, ac.Cmd, ac.Value)
	ok := airJob.Action()
	if ok != nil {
		return
	}
	acPayload := airJob.GetPayload()
	_ = acPayload

}
func (s *scheduleService) CornJob() {

	bkc, _ := time.LoadLocation("Asia/Bangkok")
	cr := cron.New(cron.WithLocation(bkc))
	_ = cr

	defer cr.Stop()
	cr.Start()

	// Demo device : 2300F15050023
	//userCommand := &airCmdReq{
	//	Cmd:      "temp",
	//	Value:    "19",
	//	SerialNo: "2300F15050023",
	//}
	//_ = userCommand
	//
	//airJob := utils.NewAirCmd(userCommand.SerialNo, userCommand.Cmd, userCommand.Value)
	//ok := airJob.Action()
	//if ok != nil {
	//	return
	//}
	//airJobPayload := airJob.GetPayload()
	//_ = airJobPayload

	for {
		time.Sleep(time.Minute)

		fmt.Println("Air Payload")
		//fmt.Println(airJobPayload)
		jobWork, err := s.WorkList()
		if err != nil {
			fmt.Println("No Data")
			return
		}
		//
		fmt.Println(len(jobWork))
		//
		for _, job := range jobWork {

			//fmt.Printf("Job No, %v \n", job.SerialNo, strings.Join(job.Duration, " "))
			fmt.Printf("%v , %s \n", job.SerialNo, strings.Join(job.Duration[:], " "))
			fmt.Println("Command", job.Command)
			//airJob := utils.NewAirCmd(userCommand.SerialNo, userCommand.Cmd, userCommand.Value)
			//ok := airJob.Action()
			//airJobPayload := airJob.GetPayload()
			fmt.Println("Serial No.", job.SerialNo)
			//fmt.Println(airJobPayload)

			//if ok != nil {
			//	return
			//}

		}

	}

	/*
		cr := cron.New(cron.WithLocation(bkc))
		fmt.Println("on Top Out Corn")

		cr.AddFunc("* * * * *", func() {
			fmt.Println(" Out CornJOb Task", now.Local())
		})
		cr.Start()
	*/

	//for {
	//	time.Sleep(time.Minute)
	//	now := time.Now()
	//	tez := fmt.Sprintf("Schedule at : %s", now.Local())
	//	fmt.Println(tez)
	//
	//	cr.AddFunc("*/2 * * * *", func() {
	//		fmt.Println(" In CornJOb Task in 2 minute", now.Local())
	//	})
	//
	//	cr.AddFunc("*/4 * * * *", func() {
	//		fmt.Println(" In CornJOb Task in 4 minute", now.Local())
	//	})
	//
	//	defer cr.Stop()
	//
	//}

	//inspect(cr.Entries())
}
func (s *scheduleService) WorkList() ([]*JobWork, error) {

	workList := []*JobWork{}

	jobs, err := s.scheduleRepo.JobsSchedule()
	if err != nil {
		return nil, err
	}

	for _, item := range jobs {

		acCmd := []AirCmd{}
		for _, i := range item.Command {
			cmd := &AirCmd{
				Cmd:   i.Cmd,
				Value: i.Value,
			}
			acCmd = append(acCmd, *cmd)
		}

		job := &JobWork{
			SerialNo:  item.SerialNo,
			Command:   acCmd,
			Mode:      item.Mode,
			Duration:  item.Duration,
			Status:    item.Status,
			StartDate: item.StartDate,
			EndDate:   item.EndDate,
		}
		workList = append(workList, job)
	}

	return workList, err

}
