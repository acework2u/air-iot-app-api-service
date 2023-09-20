package handler

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type JobsHandler struct {
	thingsService services.ThinksService
	res           utils.Response
	jobsService   services.JobsService
}

func NewJobsHandler(jobsService services.JobsService, thingsService services.ThinksService) JobsHandler {

	return JobsHandler{jobsService: jobsService, thingsService: thingsService, res: utils.Response{}}
}

func (h *JobsHandler) GetJobsDevice(c *gin.Context) {

	userId, _ := c.Get("UserId")

	job, err := h.jobsService.JobsThingsHandler(userId.(string))

	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	msg := fmt.Sprintf("%s", job)

	h.res.Success(c, msg)
}
