package handler

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	thingsService services.ThinksService
	resp          utils.Response
	jobsService   services.JobsService
}

func NewJobsHandler(jobsService services.JobsService, thingsService services.ThinksService) JobsHandler {

	return JobsHandler{jobsService: jobsService, thingsService: thingsService, resp: utils.Response{}}
}

func (h *JobsHandler) GetJobsDevice(c *gin.Context) {

	userId, _ := c.Get("UserId")

	job, err := h.jobsService.JobsThingsHandler(userId.(string))

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	msg := fmt.Sprintf("%s", job)

	h.resp.Success(c, msg)
}
func (h *JobsHandler) GetJobsShadowsDevice(c *gin.Context) {

	//userId, _ := c.Get("UserId")

	deviceSn := services.DeviceReq{}

	c.ShouldBindUri(&deviceSn)

	//fmt.Println(deviceSn.DeviceSn)
	//fmt.Println("Device SN")

	job, err := h.jobsService.GetJobsThings(deviceSn.DeviceSn)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, job)
}
func (h *JobsHandler) GetJobsQDevice(c *gin.Context) {

	//userId, _ := c.Get("UserId")

	deviceSn := services.DeviceReq{}

	_ = c.ShouldBindUri(&deviceSn)

	//fmt.Println(deviceSn.DeviceSn)
	//fmt.Println("Device SN")	//fmt.Println(deviceSn.DeviceSn)
	//fmt.Println("Device SN")

	job, err := h.jobsService.GetQueJobsThings(deviceSn.DeviceSn)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, job)
}
func (h *JobsHandler) PostCreateJobs(c *gin.Context) {

	userId, _ := c.Get("UserId")

	jobInput := services.CreateJobsReq{}

	err := c.ShouldBindJSON(&jobInput)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	jobInput.JobId = fmt.Sprintf("%v_%v", jobInput.DeviceSn, userId)

	h.resp.Success(c, jobInput)
}
func (h *JobsHandler) PutUpdateJobs(c *gin.Context) {

	h.resp.Success(c, "Jobs ")
}
