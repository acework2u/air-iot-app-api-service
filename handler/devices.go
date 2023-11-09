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
	resp          utils.Response
}

func NewDeviceHandler(deviceService service.DevicesService) DevicesHandler {

	return DevicesHandler{deviceService: deviceService, resp: utils.Response{}}
}

// GetDevice godoc
// @Summary Get Device list
// @Description Get device list
// @Tags ThingsDevice
// @Security BearerAuth
// @Param DeviceRequest body service.DeviceRequest true "Get Device list"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /devices [get]
func (h *DevicesHandler) GetDevice(c *gin.Context) {

	userId, _ := c.Get("UserId")

	deviceReq := &service.DeviceRequest{
		UserId: userId.(string),
	}

	deviceResponse, err := h.deviceService.ListDevice(deviceReq)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, deviceResponse)

}

// PostDevice godoc
// @Summary New a Things Device
// @Description user is a new things device
// @Tags ThingsDevice
// @Security BearerAuth
// @Param DeviceInfo body service.Device true "a device information"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /devices [post]
func (h *DevicesHandler) PostDevice(c *gin.Context) {

	var deviceInfo *service.Device

	userId, _ := c.Get("UserId")

	if err := c.ShouldBindJSON(&deviceInfo); err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	deviceInfo.UserId = userId.(string)
	deviceInfo.SerialNo = strings.ToUpper(deviceInfo.SerialNo)
	resInfo, err := h.deviceService.NewDevice(deviceInfo)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	mesRes := &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: resInfo,
	}
	h.resp.Success(c, mesRes)

}

// PutDevice godoc
// @Summary Update my things device
// @Description Update my things device
// @Tags ThingsDevice
// @Security BearerAuth
// @Param DeviceInfo body service.ReqUpdateDevice true "Update things device"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /devices/:id [put]
func (h *DevicesHandler) PutDevice(c *gin.Context) {

	var deviceInfo *service.ReqUpdateDevice
	err := c.ShouldBindJSON(&deviceInfo)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	//userId, _ := c.Get("UserId")
	deviceId := c.Param("id")
	resDevice, err := h.deviceService.UpdateDevice(deviceId, deviceInfo)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, resDevice)
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

// DelDevice godoc
// @Summary Delete a things device
// @Description Delete a things device
// @Tags ThingsDevice
// @Security BearerAuth
// @Param id path string true "Device id"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /devices/:id [delete]
func (h *DevicesHandler) DelDevice(c *gin.Context) {

	deviceID := c.Param("id")
	userID, _ := c.Get("UserId")
	filter := &service.DeviceFilter{
		UserId: userID.(string),
		Id:     deviceID,
	}
	_, err := h.deviceService.DeleteDevice(filter)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, "delete success success")
}
