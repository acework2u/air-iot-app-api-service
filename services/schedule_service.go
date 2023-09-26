package services

import (
	"errors"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"time"
)

type scheduleService struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleService(scheduleRepo repository.ScheduleRepository) ScheduleService {
	return &scheduleService{scheduleRepo}
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
		job := &JobDbSchedule{
			Id:        jobs.Id.String(),
			SerialNo:  jobs.SerialNo,
			Command:   jobs.Command,
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
func (s *scheduleService) NewJobSchedules(userId string, jobInfo *JobSchedule) (*JobDbSchedule, error) {

	dataInfo := &repository.ScheduleJob{
		SerialNo:  jobInfo.SerialNo,
		UserId:    jobInfo.UserId,
		Command:   jobInfo.Command,
		Mode:      jobInfo.Mode,
		Duration:  jobInfo.Duration,
		StartDate: jobInfo.StartDate,
		EndDate:   jobInfo.EndDate,
	}

	job, err := s.scheduleRepo.NewJob(userId, dataInfo)

	if err != nil {
		return nil, err
	}

	resJob := &JobDbSchedule{
		SerialNo:  job.SerialNo,
		Command:   job.Command,
		Mode:      job.Mode,
		Duration:  job.Duration,
		StartDate: job.StartDate,
		EndDate:   job.EndDate,
		Status:    job.Status,
	}

	return resJob, nil
}

func (s *scheduleService) CornJob() {

	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Schedule Job")
	}
}
