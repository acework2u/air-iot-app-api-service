package handler

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math"
	"net/http"
	"time"
)

type ThingsHandler struct {
	thingsService services.ThinksService
}

type airCmdReq struct {
	SerialNo string `json:"serialNo" validate:"required" binding:"required"`
	Cmd      string `json:"cmd" validate:"required" binding:"required"`
	Value    string `json:"value" validate:"required" binding:"required"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewThingsHandler(thingsService services.ThinksService) ThingsHandler {
	return ThingsHandler{thingsService}
}

func (h *ThingsHandler) ConnectThing(c *gin.Context) {

	res, err := h.thingsService.GetCerts()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
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

func (h *ThingsHandler) UserCert(c *gin.Context) {

	//tokenId, _ := c.Get("UserToken")
	tokenId, _ := c.Get("UserAuthToken")

	resCert, err := h.thingsService.ThingsCert(tokenId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resCert,
	})
}

func (h *ThingsHandler) ThingsCert(c *gin.Context) {

	idToken, _ := c.Get("UserToken")

	fmt.Println(idToken)

	airMode, _ := hex.DecodeString("1")

	fmt.Println("airMode")
	fmt.Println(airMode)

	airPayload := make([]byte, 40)
	copy(airPayload[0:], string(1))
	copy(airPayload[1:], string(3))
	copy(airPayload[2:], string(60))
	copy(airPayload[3:], string(120))
	copy(airPayload[4:], string(1))
	copy(airPayload[14:], string(1))

	fmt.Println(airPayload)
	fmt.Println(airPayload)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": airPayload,
	})
}

func (h *ThingsHandler) UploadFile(c *gin.Context) {

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	res, err := h.thingsService.UploadToS3(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

func (h *ThingsHandler) ThingsRegister(c *gin.Context) {

	userSub, _ := c.Get("UserSub")
	userID, _ := c.Get("UserId")

	ress, err := h.thingsService.ThingRegister(userID.(string))

	if err != nil {
		fmt.Println("Err")
		fmt.Println(err.Error())
	}

	_ = ress

	fmt.Println(userSub)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": ress,
	})
}

func (h *ThingsHandler) ThingConnect(c *gin.Context) {

	userID, _ := c.Get("UserId")

	resp, err := h.thingsService.ThingsConnected(userID.(string), "")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": &resp,
	})
}

func (h *ThingsHandler) CmdThing(c *gin.Context) {

	var userCmd *airCmdReq

	if err := c.ShouldBindJSON(&userCmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	airUser := utils.NewAirCmd(userCmd.SerialNo, userCmd.Cmd, userCmd.Value)

	ok := airUser.Action()
	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "command is wrong",
		})
		return
	}

	res := airUser.GetPayload()

	// Normal command
	//output, err := h.thingsService.ThingsConnected(res, userCmd.SerialNo)

	// Shadows Ac Control Command

	userID, _ := c.Get("UserSub")
	shadows, err := h.thingsService.ThinksShadows(userID.(string), res)

	_ = shadows

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": shadows,
	})

}

func (h *ThingsHandler) PostShadows(c *gin.Context) {

	userCmd := &airCmdReq{}

	err := c.ShouldBindJSON(userCmd)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}

	airUser := utils.NewAirCmd(userCmd.SerialNo, userCmd.Cmd, userCmd.Value)

	ok := airUser.Action()
	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "command is wrong",
		})
		return
	}

	shadowPayload := airUser.GetPayload()

	res, err := h.thingsService.PubUpdateShadows(userCmd.SerialNo, shadowPayload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})

}

func (h *ThingsHandler) Shadows(c *gin.Context) {
	//userID, _ := c.Get("UserId")
	userID, _ := c.Get("UserSub")
	res := ""
	resShadows, _ := h.thingsService.ThinksShadows(userID.(string), res)

	shadows, err := utils.GetClaimsFromToken(resShadows.State.Reported.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
	}
	reg1000 := shadows["data"].(map[string]interface{})["reg1000"].(string)
	acVal := utils.NewGetAcVal(reg1000)
	ac1000 := acVal.Ac1000()

	c.JSON(http.StatusOK, gin.H{
		"message": ac1000,
	})

}

func (h *ThingsHandler) WsShadows(c *gin.Context) {

	//thingName := "2300F15050017"
	shadowName := "air-users"
	acCmdReq := &airCmdReq{}

	err := c.ShouldBindJSON(acCmdReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}
	thinkName := fmt.Sprintf("%v", acCmdReq.SerialNo)
	data, err := h.thingsService.PubGetShadows(thinkName, shadowName)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"Message": data,
	})

}

func (h *ThingsHandler) PostCmdShadows(c *gin.Context) {

	//thingName := "2300F15050017"
	shadowName := "air-users"
	acCmdReq := &airCmdReq{}

	err := c.ShouldBindJSON(acCmdReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})
		return
	}
	thinkName := fmt.Sprintf("%v", acCmdReq.SerialNo)
	data, err := h.thingsService.PubGetShadows(thinkName, shadowName)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"Message": data,
	})

}

func (h *ThingsHandler) WsIoT(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	cogintoId := "646c33ba0e5800006e000abd"
	client, err := h.thingsService.NewAwsMqttConnect(cogintoId)
	shadowsAcceptTopic := "$aws/things/2300F15050023/shadow/name/air-users/update/accepted"
	shadowRejectTopic := "$aws/things/2300F15050023/shadow/name/air-users/update/rejected"
	shadowsUpdateDocTopic := "$aws/things/2300F15050023/shadow/name/air-users/update/documents"

	revMsg := make(chan *services.IndoorInfo)

	for {
		//Read Message from client
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
			//break
		}
		//If client message is ping will return pong
		//if string(message) == "ping" {
		//	message = []byte("pong")
		//}
		fmt.Println("Ws Working...")
		err = client.SubscribeWithHandler(shadowsAcceptTopic, 0, func(client MQTT.Client, message MQTT.Message) {
			//msgPayload := fmt.Sprintf(`%v`, string(message.Payload()))

			//fmt.Println("In update accepted")
			//fmt.Println(msgPayload)
			shadowDoc := &utils.ShadowAcceptStrut{}
			json.Unmarshal(message.Payload(), shadowDoc)

			//resData := message.Payload()
			//resOutData = &resData
			//Response message to client
			ws.WriteMessage(mt, message.Payload())

		})
		err = client.SubscribeWithHandler(shadowsUpdateDocTopic, 0, func(client MQTT.Client, message MQTT.Message) {
			msgPayload := fmt.Sprintf(`%v`, string(message.Payload()))
			_ = msgPayload
			//fmt.Println("In update Doc")
			//fmt.Println(msgPayload)
			//resData := message.Payload()
			//resOutData = &resData
			//Response message to client
			ws.WriteMessage(mt, message.Payload())

		})
		err = client.SubscribeWithHandler(shadowRejectTopic, 0, func(client MQTT.Client, message MQTT.Message) {
			msgPayload := fmt.Sprintf(`%v`, string(message.Payload()))
			fmt.Println("In Update Reject")
			fmt.Println(msgPayload)
			//resData := message.Payload()
			//resOutData = &resData
			//Response message to client
			ws.WriteMessage(mt, message.Payload())

		})

		go func(msg chan<- *services.IndoorInfo) {
			data, err := h.thingsService.PubGetShadows("2300F15050023", "")

			if err != nil {
				return
			}
			revMsg <- data
			time.Sleep(4 * time.Second)
		}(revMsg)

		if err != nil {
			return
		}

		message, _ = json.Marshal(<-revMsg)
		//Response message to client
		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			//break
			return
		}
	}

}

func decimalToBinary(num int) {
	var binary []int

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}
	if len(binary) == 0 {
		fmt.Printf("%d\n", 0)
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			fmt.Printf("%d", binary[i])
		}
		fmt.Println()
	}
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
