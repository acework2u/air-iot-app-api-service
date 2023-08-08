package handler

import (
	"encoding/hex"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strings"
)

type ThingsHandler struct {
	thingsService services.ThinksService
}

type airCmdReq struct {
	SerialNo string `json:"serialNo" validate:"required" binding:"required"`
	Cmd      string `json:"cmd" validate:"required" binding:"required"`
	Value    string `json:"value" validate:"required" binding:"required"`
}

func NewThingsHandler(thingsService services.ThinksService) ThingsHandler {
	return ThingsHandler{thingsService}
}

func (h *ThingsHandler) ConnectThing(c *gin.Context) {

	res, err := h.thingsService.GetCerds()

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

	switch strings.ToLower(userCmd.Cmd) {
	case "power":
	case "temp":
	case "mode":
	case "fan":
	case "swing":

	}

	if userCmd.Value == "on" {
		userCmd.Value = "1"
	}
	if userCmd.Value == "off" {
		userCmd.Value = "0"
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

func (h *ThingsHandler) Shadows(c *gin.Context) {
	//userID, _ := c.Get("UserId")
	userID, _ := c.Get("UserSub")

	_ = userID

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
