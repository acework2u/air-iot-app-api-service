package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceRouter struct {
	deviceHandler handler.DevicesHandler
}

func NewDeviceRouter(devicesHandler handler.DevicesHandler) DeviceRouter {
	return DeviceRouter{devicesHandler}
}

func (r *DeviceRouter) DeviceRoute(rg *gin.RouterGroup) {

	router := rg.Group("/devices", middleware.CognitoAuthMiddleware())
	router.GET("", r.deviceHandler.GetDevice)
	router.GET("/checkdup", r.deviceHandler.GetCheckDup)
	router.POST("", r.deviceHandler.PostDevice)
	router.PUT("/:id", r.deviceHandler.PutDevice)
}
