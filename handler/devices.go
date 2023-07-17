package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type DevicesHandler struct {
	deviceService service.DevicesService
}

func NewDeviceHandler(deviceService service.DevicesService) DevicesHandler {

	return DevicesHandler{deviceService: deviceService}
}

func (h *DevicesHandler) GetDevice(c *gin.Context) {

	var deviceInfo *service.Device

	if err := c.ShouldBindJSON(&deviceInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	mesRes := &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: deviceInfo,
	}

	utils.ResponseSuccess(c, mesRes)
}

func (h *DevicesHandler) PostDevice(c *gin.Context) {

	var deviceInfo *service.Device

	userId, _ := c.Get("UserId")

	if err := c.ShouldBindJSON(&deviceInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//
	deviceInfo.UserId = userId.(string)
	deviceInfo.SerialNo = strings.ToUpper(deviceInfo.SerialNo)
	resInfo, err := h.deviceService.NewDevice(deviceInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	mesRes := &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: resInfo,
	}

	utils.ResponseSuccess(c, mesRes)

}
