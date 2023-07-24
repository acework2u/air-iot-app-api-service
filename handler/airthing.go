package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
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
