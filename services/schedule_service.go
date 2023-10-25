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
func (s *scheduleService) DeleteJobSchedule(jobId string) error {
	err := s.scheduleRepo.DeleteJob(jobId)
	if err != nil {
		return err
	}
	return nil
}

func (s *scheduleService) job(ac AirJob) {

	if len(ac.Command) > 0 {

		for _, item := range ac.Command {
			airJob := utils.NewAirCmd(ac.SerialNo, item.Cmd, item.Value)
			ok := airJob.Action()

			if ok != nil {
				return
			}
			acPayload := airJob.GetPayload()

			if len(acPayload) > 10 {
				go s.thingServ.PubUpdateShadows(ac.SerialNo, acPayload)

			}

			//fmt.Println("SN: ", ac.SerialNo, "ac Command ", item.Cmd, " ", item.Value)
			//fmt.Println(acPayload)

			//fmt.Printf(" serial %v : payload : %v", ac.SerialNo, acPayload)

			//fmt.Printf(" serial %v : payload : %v", ac.SerialNo, acPayload)
			//go s.thingServ.PubUpdateShadows(ac.SerialNo, acPayload)
			//go func() {
			//
			//	fmt.Printf(" serial %v : payload : %v", ac.SerialNo, acPayload)
			//
			//	if len(acPayload) > 10 {
			//		_, err := s.thingServ.PubUpdateShadows(ac.SerialNo, acPayload)
			//
			//		if err != nil {
			//			fmt.Println("Err In Job")
			//			fmt.Println("SN: ", ac.SerialNo, "ac Command ", item.Cmd, " ", item.Value)
			//
			//			fmt.Println(err.Error())
			//			return
			//			//return
			//		}
			//	}
			//
			//}()

			time.Sleep(2 * time.Second)

		}

	}

}
func (s *scheduleService) job2(ac AirJob) {

}
func (s *scheduleService) CornJob() {

	bkc, _ := time.LoadLocation("Asia/Bangkok")
	cr := cron.New(cron.WithLocation(bkc))
	_ = cr
	defer cr.Stop()
	cr.Start()

	for {

		time.Sleep(time.Minute)
		//time.Sleep(4 * time.Second)

		fmt.Println("Air Payload")

		jobWork, err := s.WorkList()

		if err != nil {
			fmt.Println("No Data")
			return
		}
		//
		fmt.Println(len(jobWork))
		//
		myJobs := AirJob{}
		for _, job := range jobWork {

			myJobs.SerialNo = job.SerialNo
			myJobs.Command = job.Command

			fmt.Println(myJobs)

			//s.job(*myJobs)
			//cr.AddFunc()

			//fmt.Println("myJobs ")
			//fmt.Println(myJobs)

			//cr.AddFunc(getTimeCornJob(job.Duration), func() {
			//	s.job(*myJobs)
			//})
			//s.job(*myJobs)
			//
			//fmt.Printf("%v , %s \n", job.SerialNo, getTimeCornJob(job.Duration))
			//fmt.Println("Command", job.Command)

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
func getTimeCornJob(duration []string) string {

	strDuration := strings.Join(duration[:], " ")

	return strDuration
}
