package handler

import (
	"fmt"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type ScheduleHandle struct {
	scheduleService service.ScheduleService
	res             utils.Response
}

func NewScheduleHandler(scheduleService service.ScheduleService) ScheduleHandle {
	return ScheduleHandle{scheduleService: scheduleService, res: utils.Response{}}
}

func (h *ScheduleHandle) GetScheduleJobs(c *gin.Context) {

	userId, _ := c.Get("UserId")

	fmt.Println(userId)

	res, err := h.scheduleService.GetSchedules(userId.(string))
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	h.res.Success(c, res)

}

func (h *ScheduleHandle) PostScheduleJobs(c *gin.Context) {

	userId, _ := c.Get("UserId")
	jobInfo := &service.JobScheduleReq{}

	fmt.Println(userId)

	err := c.ShouldBindJSON(jobInfo)

	cusErr := utils.NewCustomHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return
	}
	now := time.Now()
	jobInfo.CreatedDate = now.Local()
	jobInfo.UpdatedDate = jobInfo.CreatedDate
	jobInfo.UserId = userId.(string)

	fmt.Println(jobInfo)
	//res := jobInfo
	res, err := h.scheduleService.NewJobSchedules(userId.(string), jobInfo)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	h.res.Success(c, res)

}

func (h *ScheduleHandle) UpdateScheduleJobs(c *gin.Context) {

	updateReq := service.UpdateJobSchedule{}
	jobId := c.Param("jobId")

	err := c.ShouldBindJSON(&updateReq)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	jobDb, err := h.scheduleService.UpdateJobInSchedule(jobId, &updateReq)

	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	updateTxt := fmt.Sprintf("Uptated is a success : %s ,\n %v", jobId, jobDb)

	h.res.Success(c, updateTxt)
}

func (h *ScheduleHandle) DelScheduleJobs(c *gin.Context) {

	jobId := c.Param("jobId")

	if len(jobId) < 0 {
		h.res.BadRequest(c, "job id is required")
		return
	}
	err := h.scheduleService.DeleteJobSchedule(jobId)
	if err != nil {
		h.res.BadRequest(c, "can't delete a job")
		return
	}
	//job := c.Param("jobId")
	delText := fmt.Sprintf("delete a schedule job %s to successful", jobId)
	h.res.Success(c, delText)
}
