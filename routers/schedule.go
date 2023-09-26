package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type ScheduleRouter struct {
	scheduleHandler handler.ScheduleHandle
}

func NewScheduleRouter(scheduleHandler handler.ScheduleHandle) ScheduleRouter {
	return ScheduleRouter{scheduleHandler: scheduleHandler}
}
func (r *ScheduleRouter) ScheduleRoute(rg *gin.RouterGroup) {
	router := rg.Group("schedule", middleware.CognitoAuthMiddleware())
	router.GET("", r.scheduleHandler.GetScheduleJobs)
	router.POST("", r.scheduleHandler.PostScheduleJobs)
}
