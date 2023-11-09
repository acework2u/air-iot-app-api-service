package handler

import (
	"fmt"
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

// GetDevice godoc
// @Summary Get Device list
func (h *DevicesHandler) GetDevice(c *gin.Context) {

	userId, _ := c.Get("UserId")

	deviceReq := &service.DeviceRequest{
		UserId: userId.(string),
	}

	deviceResponse, err := h.deviceService.ListDevice(deviceReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": deviceResponse,
	})

	//mesRes := &utils.ApiResponse{
	//	Status:  http.StatusOK,
	//	Message: deviceResponse,
	//}
	//utils.ResponseSuccess(c, mesRes)
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
	deviceInfo.UserId = userId.(string)
	deviceInfo.SerialNo = strings.ToUpper(deviceInfo.SerialNo)
	resInfo, err := h.deviceService.NewDevice(deviceInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	mesRes := &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: resInfo,
	}

	utils.ResponseSuccess(c, mesRes)

}

func (h *DevicesHandler) PutDevice(c *gin.Context) {
	var deviceInfo *service.ReqUpdateDevice
	err := c.ShouldBindJSON(&deviceInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	userId, _ := c.Get("UserId")

	deviceId := c.Param("id")

	_ = userId

	resDevice, err := h.deviceService.UpdateDevice(deviceId, deviceInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resDevice,
	})
}

func (h *DevicesHandler) GetCheckDup(c *gin.Context) {

	userId, _ := c.Get("UserId")
	serialNo := strings.ToUpper("33f3qwi0008222920")
	dup := h.deviceService.CheckDup(userId.(string), serialNo)

	textVal := "dup < = 0"

	if dup > 0 {
		textVal = "dup > 0"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Check Dup %v", textVal),
	})
}

func (h *DevicesHandler) DelDevice(c *gin.Context) {

	deviceID := c.Param("id")
	userID, _ := c.Get("UserId")
	filter := &service.DeviceFilter{
		UserId: userID.(string),
		Id:     deviceID,
	}
	_, err := h.deviceService.DeleteDevice(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Delete Completed",
	})
}
