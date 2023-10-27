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
			Id:        jobs.Id,
			JobId:     jobs.JobId,
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
		JobId:       int(0),
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
		Id:        job.Id,
		SerialNo:  job.SerialNo,
		JobId:     job.JobId,
		Command:   jobs,
		Mode:      job.Mode,
		Duration:  job.Duration,
		StartDate: job.StartDate,
		EndDate:   job.EndDate,
		Status:    job.Status,
	}

	return resJob, nil
}
func (s *scheduleService) UpdateJobInSchedule(jobId string, jobInfo *UpdateJobSchedule) (*JobDbSchedule, error) {
	jobRes := &JobDbSchedule{}

	updateInfo := &repository.ScheduleJobUpdate{}

	job, err := s.scheduleRepo.UpdateJob(jobId, updateInfo)

	if err != nil {
		return nil, err
	}

	jobRes = &JobDbSchedule{
		SerialNo: job.SerialNo,
		Duration: job.Duration,
	}

	return jobRes, nil
}
func (s *scheduleService) DeleteJobSchedule(jobId string) error {
	err := s.scheduleRepo.DeleteJob(jobId)
	if err != nil {
		return err
	}
	return nil
}
func (s *scheduleService) UpdateJobId(jobId string, entryId int) error {

	info := &repository.ScheduleJobInfo{JobId: entryId}
	err := s.scheduleRepo.UpdateScheduleId(jobId, info)
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

			time.Sleep(2 * time.Second)

		}

	}

}
func (s *scheduleService) CornJob() {

	bkc, _ := time.LoadLocation("Asia/Bangkok")
	cr := cron.New(cron.WithLocation(bkc))

	cr.Start()
	defer cr.Stop()
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
		numberJob := len(jobWork)

		//

		for _, job := range jobWork {
			myJobs := AirJob{}
			myJobs.SerialNo = job.SerialNo
			myJobs.Command = job.Command

			if !job.Status {
				jobId := cron.EntryID(int(job.JobId))
				cr.Remove(jobId)
				fmt.Println("Job id Remove is : ", jobId)

			}

			if job.Status {

				crn := cr.Entries()

				fmt.Println("crn = ", crn)
				idJob := job.Id.Hex()
				if job.JobId == 0 {

					fmt.Printf("ID %s \n", idJob)
					fmt.Println("cron ID", job.JobId)

					schId, err := cr.AddFunc(getTimeCornJob(job.Duration), func() {
						s.job(myJobs)
					})
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Printf("%s", schId)
					err = s.UpdateJobId(idJob, int(schId))
					if err != nil {
						fmt.Println("Error in UpdateJObId")
						fmt.Println(err.Error())
					}

				} else {
					// Have a job in db
					m := cr.Entries()
					if len(m) >= 0 && len(m) < numberJob {
						fmt.Println("lese then 0")

						schId, err := cr.AddFunc(getTimeCornJob(job.Duration), func() {
							s.job(myJobs)
						})
						if err != nil {
							fmt.Println(err.Error())
						}
						fmt.Printf("%s", schId)
						err = s.UpdateJobId(idJob, int(schId))
						if err != nil {
							fmt.Println("Error in UpdateJObId")
							fmt.Println(err.Error())
						}

					}

					//fmt.Println(crn)

				}

				//err = s.UpdateJobId(Id, int(schId))
				//if err != nil {
				//	fmt.Println("Error in UpdateJObId")
				//	fmt.Println(err.Error())
				//}
				//
				//fmt.Println(myJobs)
			}

			m := cr.Entries()

			fmt.Println("This a job Entry is : ", len(m))

			for _, g := range m {
				fmt.Printf("%v", g.ID)
			}

		}

	}

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
			Id:        item.Id,
			JobId:     item.JobId,
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
