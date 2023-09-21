package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type JobsController struct {
	jobsHandler handler.JobsHandler
}

func NewJobsController(jobsHandler handler.JobsHandler) JobsController {
	return JobsController{jobsHandler: jobsHandler}
}

func (r *JobsController) JobsRoute(rg *gin.RouterGroup) {

	router := rg.Group("air-jobs", middleware.CognitoAuthMiddleware())
	router.GET("", r.jobsHandler.GetJobsDevice)
	router.POST("", r.jobsHandler.PostCreateJobs)
	router.PUT("", r.jobsHandler.PostCreateJobs)

	router.GET("/:deviceSn", r.jobsHandler.GetJobsShadowsDevice)
	router.GET("/q/:deviceSn", r.jobsHandler.GetJobsQDevice)

}
