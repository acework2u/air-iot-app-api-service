package handler

import (
	"fmt"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type ScheduleHandle struct {
	scheduleService service.ScheduleService
	resp            utils.Response
}

func NewScheduleHandler(scheduleService service.ScheduleService) ScheduleHandle {
	return ScheduleHandle{scheduleService: scheduleService, resp: utils.Response{}}
}

// ScheduleJobs godoc
// @Summary get a schedule job air things
// @Description get a schedule job air things
// @Produce json
// @Tags AirScheduleJobs
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /schedule [get]
func (h *ScheduleHandle) GetScheduleJobs(c *gin.Context) {

	userId, _ := c.Get("UserId")

	log.Println(userId)

	res, err := h.scheduleService.GetSchedules(userId.(string))
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)

}

// ScheduleJobs godoc
// @Summary add a new job for Schedule
// @Description add a new job for Schedule
// @Produce json
// @Tags AirScheduleJobs
// @Security BearerAuth
// @Param jobScheduleReq body service.JobScheduleReq true "air scheule job"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /schedule [post]
func (h *ScheduleHandle) PostScheduleJobs(c *gin.Context) {

	userId, _ := c.Get("UserId")
	jobInfo := &service.JobScheduleReq{}

	err := c.ShouldBindJSON(jobInfo)

	cusErr := utils.NewErrorHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return
	}
	now := time.Now()
	jobInfo.CreatedDate = now.Local()
	jobInfo.UpdatedDate = jobInfo.CreatedDate
	jobInfo.UserId = userId.(string)

	log.Println(jobInfo)
	//res := jobInfo
	res, err := h.scheduleService.NewJobSchedules(userId.(string), jobInfo)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Success
	h.resp.Success(c, res)

}

func (h *ScheduleHandle) UpdateScheduleJobs(c *gin.Context) {

	updateReq := service.UpdateJobSchedule{}
	jobId := c.Param("jobId")

	err := c.ShouldBindJSON(&updateReq)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	jobDb, err := h.scheduleService.UpdateJobInSchedule(jobId, &updateReq)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	updateTxt := fmt.Sprintf("Uptated is a success : %s ,\n %v", jobId, jobDb)

	h.resp.Success(c, updateTxt)
}

func (h *ScheduleHandle) DelScheduleJobs(c *gin.Context) {

	jobId := c.Param("jobId")

	if len(jobId) < 0 {
		h.resp.BadRequest(c, "job id is required")
		return
	}
	err := h.scheduleService.DeleteJobSchedule(jobId)
	if err != nil {
		h.resp.BadRequest(c, "can't delete a job")
		return
	}
	//job := c.Param("jobId")
	// Success
	delText := fmt.Sprintf("delete a schedule job %s to successful", jobId)
	h.resp.Success(c, delText)
}
