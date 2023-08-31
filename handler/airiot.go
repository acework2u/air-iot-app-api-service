package handler

import (
	"encoding/json"
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
func (h *AirIotHandler) WsIoT(c *gin.Context) {
	//conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	return
	//}
	//defer conn.Close()

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
			//break
		}
		fmt.Println("Ws Working...")
		fmt.Println(message)
		msg := fmt.Sprintf("WIOT OK")
		ok, _ := json.Marshal(msg)
		ws.WriteMessage(mt, ok)
	}

}
func (h *AirIotHandler) MqSubShadows(c *gin.Context) {
	airSerial := &airs.AirRq{}
	userID, _ := c.Get("UserSub")

	err := c.ShouldBindJSON(airSerial)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	serial := fmt.Sprintf("%s", airSerial.Serial)
	clientId := fmt.Sprintf("%v", userID)

	res, err := h.airIoTService.ShadowsAir(clientId, serial)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}
