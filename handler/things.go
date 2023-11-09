package handler

import (
	"encoding/hex"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math"
	"net/http"
)

type ThingsHandler struct {
	thingsService services.ThinksService
	resp          utils.Response
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
	return ThingsHandler{thingsService: thingsService, resp: utils.Response{}}
}

func (h *ThingsHandler) ConnectThing(c *gin.Context) {

	res, err := h.thingsService.GetCerts()
	if err != nil {
		h.resp.BadRequest(c, err)
		return
	}
	h.resp.Success(c, res)
}

func (h *ThingsHandler) UserCert(c *gin.Context) {

	//tokenId, _ := c.Get("UserToken")
	tokenId, _ := c.Get("UserAuthToken")

	resCert, err := h.thingsService.ThingsCert(tokenId.(string))
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, resCert)
}

func (h *ThingsHandler) ThingsPayload(c *gin.Context) {

	airMode, _ := hex.DecodeString("1")

	fmt.Println("airMode 1 hex to string")
	fmt.Println(airMode)

	airPayload := make([]byte, 40)
	copy(airPayload[0:], string(1))
	copy(airPayload[1:], string(3))
	copy(airPayload[2:], string(60))
	copy(airPayload[3:], string(120))
	copy(airPayload[4:], string(1))
	copy(airPayload[14:], string(1))

	fmt.Println(airPayload)
	h.resp.Success(c, airPayload)

}

func (h *ThingsHandler) UploadFile(c *gin.Context) {

	file, err := c.FormFile("image")

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	res, err := h.thingsService.UploadToS3(file)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)
}

func (h *ThingsHandler) ThingsRegister(c *gin.Context) {

	//userSub, _ := c.Get("UserSub")
	userID, _ := c.Get("UserId")

	RegisResponse, err := h.thingsService.ThingRegister(userID.(string))

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, RegisResponse)

}

func (h *ThingsHandler) ThingConnect(c *gin.Context) {

	userID, _ := c.Get("UserId")

	resp, err := h.thingsService.ThingsConnected(userID.(string), "")

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, resp)
}

func (h *ThingsHandler) CmdThing(c *gin.Context) {

	var userCmd *airCmdReq

	if err := c.ShouldBindJSON(&userCmd); err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	airUser := utils.NewAirCmd(userCmd.SerialNo, userCmd.Cmd, userCmd.Value)

	ok := airUser.Action()
	if ok != nil {
		h.resp.BadRequest(c, "command's wrong")
		return
	}
	res := airUser.GetPayload()

	userID, _ := c.Get("UserSub")
	shadows, err := h.thingsService.ThinksShadows(userID.(string), res)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Show shadows
	h.resp.Success(c, shadows)
}

// PostShadows godoc
// @Summary Air things shadows command
// @Description Air things shadows command
// @Tags AirThingsCommand
// @Security BearerAuth
// @Param AirCommandReq body airCmdReq true "Air cmd request"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /thing/shadows [post]
func (h *ThingsHandler) PostShadows(c *gin.Context) {

	userCmd := &airCmdReq{}
	err := c.ShouldBindJSON(userCmd)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	airUser := utils.NewAirCmd(userCmd.SerialNo, userCmd.Cmd, userCmd.Value)

	ok := airUser.Action()
	if ok != nil {
		h.resp.BadRequest(c, "command's wrong")
		return
	}

	shadowPayload := airUser.GetPayload()

	res, err := h.thingsService.PubUpdateShadows(userCmd.SerialNo, shadowPayload)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)

}

func (h *ThingsHandler) Shadows(c *gin.Context) {
	//userID, _ := c.Get("UserId")
	userID, _ := c.Get("UserSub")
	res := ""
	resShadows, _ := h.thingsService.ThinksShadows(userID.(string), res)

	shadows, err := utils.GetClaimsFromToken(resShadows.State.Reported.Message)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	reg1000 := shadows["data"].(map[string]interface{})["reg1000"].(string)
	acVal := utils.NewGetAcVal(reg1000)
	ac1000 := acVal.Ac1000()

	// Success
	h.resp.Success(c, ac1000)

}

func (h *ThingsHandler) WsShadows(c *gin.Context) {

	//thingName := "2300F15050017"
	shadowName := "air-users"
	acCmdReq := &airCmdReq{}

	err := c.ShouldBindJSON(acCmdReq)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	thinkName := fmt.Sprintf("%v", acCmdReq.SerialNo)
	data, err := h.thingsService.PubGetShadows(thinkName, shadowName)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, data)

}

func (h *ThingsHandler) PostCmdShadows(c *gin.Context) {

	shadowName := "air-users"
	acCmdReq := &airCmdReq{}
	filterErr := utils.NewErrorHandler(c)
	err := c.ShouldBindJSON(acCmdReq)
	if err != nil {
		filterErr.CustomError(err)
		return
	}
	thinkName := fmt.Sprintf("%v", acCmdReq.SerialNo)
	data, err := h.thingsService.PubGetShadows(thinkName, shadowName)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, data)
}

func (h *ThingsHandler) WsIoT(c *gin.Context) {
	userId, _ := c.Get("UserId")
	h.resp.Success(c, userId)
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
