package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AirThingHandler struct {
	airThingService service.AirThinkService
}

func NewAirThingHandler(airThingService service.AirThinkService) AirThingHandler {
	return AirThingHandler{airThingService}
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

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "air all",
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
