package handler

import (
	"encoding/json"
	"fmt"
	airs "github.com/acework2u/air-iot-app-api-service/services/airiot"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type AirIotHandler struct {
	airIoTService airs.AirIoTService
	resp          utils.Response
}

func NewAirIoTHandler(airIoTService airs.AirIoTService) AirIotHandler {

	return AirIotHandler{airIoTService: airIoTService, resp: utils.Response{}}
}
func (h *AirIotHandler) GetIndoor(c *gin.Context) {

	req := &airs.AirRq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	shadowsName := "air-users"
	serial := fmt.Sprintf("%s", req.Serial)

	indoor, err := h.airIoTService.GetIndoorVal(serial, shadowsName)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, indoor)
}
func (h *AirIotHandler) GetShadowsDoc(c *gin.Context) {

	req := &airs.AirRq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	shadowsName := "air-users"
	serial := fmt.Sprintf("%s", req.Serial)

	indoor, err := h.airIoTService.GetShadowsDocument(serial, shadowsName)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	//fmt.Println(indoor)
	h.resp.Success(c, indoor)

}
func (h *AirIotHandler) CheckDefault(c *gin.Context) {

	indoor, err := h.airIoTService.CheckAwsDefault()
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	//fmt.Println(indoor)
	h.resp.Success(c, indoor)
}
func (h *AirIotHandler) CheckProduction(c *gin.Context) {
	indoor, err := h.airIoTService.CheckAwsDefault()
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	//fmt.Println(indoor)
	h.resp.Success(c, indoor)

}
func (h *AirIotHandler) CheckAWS(c *gin.Context) {
	indoor, err := h.airIoTService.CheckAws()
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	//fmt.Println(indoor)
	h.resp.Success(c, indoor)

}
func (h *AirIotHandler) WsIoT(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		//fmt.Println(err)
		log.Println(err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		_ = message
		if err != nil {
			//fmt.Println(err)
			log.Println(err)
			return
		}
		log.Println("Ws Working...")
		//msg := fmt.Sprintf("WIOT OK")
		ok, _ := json.Marshal(message)
		ws.WriteMessage(mt, ok)
	}

}
func (h *AirIotHandler) MqSubShadows(c *gin.Context) {
	airSerial := &airs.AirRq{}
	userID, _ := c.Get("UserSub")

	err := c.ShouldBindJSON(airSerial)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	serial := fmt.Sprintf("%s", airSerial.Serial)
	clientId := fmt.Sprintf("%v", userID)

	res, err := h.airIoTService.ShadowsAir(clientId, serial)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Success
	h.resp.Success(c, res)
}
func (h *AirIotHandler) Ws2Indoor(c *gin.Context) {

	userID, _ := c.Get("UserSub") //

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		//fmt.Println(err)
		log.Println(err)
		return
	}
	defer ws.Close()

	client := airs.Member{
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

	dataId := fmt.Sprintf("Ws indoor %s", userID.(string))
	h.resp.Success(c, dataId)

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

}
func (h *AirIotHandler) AcIndoor(c *gin.Context) {

	userID, _ := c.Get("UserSub")

	//res, err := h.airIoTService.Airlist(userID.(string))
	err := h.airIoTService.CheckMyAc(userID.(string), "23F01000006")
	if !err {
		h.resp.BadRequest(c, err)
		return
	}
	// is True
	h.resp.Success(c, err)
}
