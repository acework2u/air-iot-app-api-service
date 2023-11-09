package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type AirThingHandler struct {
	airThingService service.AirThinkService
	resp            utils.Response
}

func NewAirThingHandler(airThingService service.AirThinkService) AirThingHandler {
	return AirThingHandler{airThingService: airThingService, resp: utils.Response{}}
}
func (h *AirThingHandler) GetCerts(c *gin.Context) {

	userToken, _ := c.Get("UserAuthToken")

	res, err := h.airThingService.GetCerts(userToken.(string))

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)
}
func (h *AirThingHandler) Connect(c *gin.Context) {

	TokenID, _ := c.Get("UserAuthToken")

	respCert, err := h.airThingService.ThingConnect(TokenID.(string))

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, respCert)
}

// GetAirs godoc
// @Summary Get Air list
// @Description Get Air list
// @Tags AirThings
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /airs [get]
func (h *AirThingHandler) GetAirs(c *gin.Context) {

	userId, _ := c.Get("UserId")
	if len(userId.(string)) < 0 {
		h.resp.BadRequest(c, "field user id is required")
		return
	}
	resData, err := h.airThingService.GetAirs(userId.(string))

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, resData)

}

// AddAirThings godoc
// @Summary add a new air things
// @Description add a new air things
// @Tags AirThings
// @Security BearerAuth
// @Param AirInfo body service.AirInfo true "Air information"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /airs [post]
func (h *AirThingHandler) AddAir(c *gin.Context) {

	userId, _ := c.Get("UserId")
	airInfo := &service.AirInfo{}
	err := c.ShouldBindJSON(airInfo)
	airInfo.UserId = userId.(string)
	cusErr := utils.NewErrorHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return
	}

	resAir, err := h.airThingService.AddAir(airInfo)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, resAir)
}

// UpdateAirThings godoc
// @Summary Update Air Information
// @Description Update Air Information
// @Tags AirThings
// @Security BearerAuth
// @Param id path int true "Acccount ID"
// @Param AirInfo body service.UpdateAirInfo true "Update air infomation"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /air/{id} [put]
func (h *AirThingHandler) UpdateAir(c *gin.Context) {

	airInfoUpdate := service.UpdateAirInfo{}
	err := c.ShouldBindJSON(&airInfoUpdate)

	cusErr := utils.NewErrorHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}
	id := c.Param("id")
	userId, _ := c.Get("UserId")
	airInfoUpdate.UserId = userId.(string)

	filter := &service.FilterUpdate{
		Id:     id,
		UserId: userId.(string),
	}

	airInfo, err := h.airThingService.UpdateAir(filter, &airInfoUpdate)
	_ = airInfo
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, airInfoUpdate)
}

// AirDelete godoc
// @Summary Delete Air by Id
// @Description Delete Air by Id
// @Tags AirThings
// @Security BearerAuth
// @Param id path int true "Account ID"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /airs/{id} [delete]
func (h *AirThingHandler) DelAir(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("UserId")

	if len(id) < 0 {
		h.resp.BadRequest(c, "device id is required")
		return
	}

	err := h.airThingService.DeleteAir(id, userId.(string))
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, "Delete device "+id+" a successful ")
}
