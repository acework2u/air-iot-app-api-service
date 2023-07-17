package handler

import (
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type DevicesHandler struct {
	//deviceService service.DeviceService
}

func NewDeviceHandler() DevicesHandler {

	return DevicesHandler{}
}

func (h *DevicesHandler) GetDevice(c *gin.Context) {

	utils.ResponseSuccess(c, "OK")
}
