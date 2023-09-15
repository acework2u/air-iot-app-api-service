package handler

import (
	"fmt"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
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

	_ = res

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}
func (h *AirThingHandler) Connect(c *gin.Context) {

	TokenID, _ := c.Get("UserAuthToken")

	respCert, _ := h.airThingService.ThingConnect(TokenID.(string))

	_ = respCert
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": respCert,
	})
}
func (h *AirThingHandler) GetAirs(c *gin.Context) {

	userId, _ := c.Get("UserId")
	if len(userId.(string)) < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "no data",
		})
		return
	}

	resData, err := h.airThingService.GetAirs(userId.(string))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resData,
	})
}
func (h *AirThingHandler) AddAir(c *gin.Context) {

	userId, _ := c.Get("UserId")
	airInfo := &service.AirInfo{}
	err := c.ShouldBindJSON(airInfo)
	airInfo.UserId = userId.(string)
	cusErr := utils.NewCustomHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}

	resAir, err := h.airThingService.AddAir(airInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resAir,
	})
}

func (h *AirThingHandler) UpdateAir(c *gin.Context) {
	airInfoUpdate := service.UpdateAirInfo{}
	err := c.ShouldBindJSON(&airInfoUpdate)
	cusErr := utils.NewCustomHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}
	id := c.Param("id")
	userId, _ := c.Get("UserId")
	airInfoUpdate.UserId = userId.(string)
	fmt.Println(userId)
	_ = userId

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
