package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
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
	res, err := h.scheduleService.GetSchedules(userId.(string))
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	h.res.Success(c, res)

}

func (h *ScheduleHandle) PostScheduleJobs(c *gin.Context) {

	userId, _ := c.Get("UserId")
	jobInfo := &service.JobSchedule{}

	err := c.ShouldBindJSON(jobInfo)
	cusErr := utils.NewCustomHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return
	}
	res, err := h.scheduleService.NewJobSchedules(userId.(string), jobInfo)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	h.res.Success(c, res)

}
