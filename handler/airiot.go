package handler

import (
	"fmt"
	airs "github.com/acework2u/air-iot-app-api-service/services/airiot"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AirIotHandler struct {
	airIoTService airs.AirIoTService
}

func NewAirIoTHandler(airIoTService airs.AirIoTService) AirIotHandler {

	return AirIotHandler{airIoTService: airIoTService}
}
func (h *AirIotHandler) GetIndoor(c *gin.Context) {

	req := &airs.AirRq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	shadowsName := "air-users"
	serial := fmt.Sprintf("%s", req.Serial)

	indoor, err := h.airIoTService.GetIndoorVal(serial, shadowsName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": indoor,
	})
}
func (h *AirIotHandler) GetShadowsDoc(c *gin.Context) {

	req := &airs.AirRq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	shadowsName := "air-users"
	serial := fmt.Sprintf("%s", req.Serial)

	indoor, err := h.airIoTService.GetShadowsDocument(serial, shadowsName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(indoor)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": indoor,
	})
}
func (h *AirIotHandler) CheckDefault(c *gin.Context) {

	indoor, err := h.airIoTService.CheckAwsDefault()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(indoor)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": indoor,
	})
}
func (h *AirIotHandler) CheckProduction(c *gin.Context) {
	indoor, err := h.airIoTService.CheckAwsDefault()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(indoor)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": indoor,
	})

}
func (h *AirIotHandler) CheckAWS(c *gin.Context) {
	indoor, err := h.airIoTService.CheckAws()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(indoor)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": indoor,
	})

}
