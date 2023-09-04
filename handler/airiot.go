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
func (h *AirIotHandler) Ws2Indoor(c *gin.Context) {

	userID, _ := c.Get("UserSub") //

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	client := airs.Client{
		ID:   userID.(string),
		Conn: ws,
	}
	ps := h.airIoTService.Ws2AddClient(client)
	fmt.Println("New client is connected, total: ", ps.Clients)
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Something went wrong", err)
			res := h.airIoTService.Ws2RemoveClient(client)
			fmt.Println("total clients and subscriptions", len(res.Clients), len(res.Subscriptions))
			return
		}
		h.airIoTService.Ws2HandleReceiveMessage(client, msgType, msg)
	}

	//for {
	//	//Read Message from client
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	//If client message is ping will return pong
	//	if string(message) == "ping" {
	//		message = []byte("pong")
	//	}
	//	//Response message to client
	//	err = ws.WriteMessage(mt, message)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Ws2 Indoor" + userID.(string),
	})
}
