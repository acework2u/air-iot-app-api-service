package handler

import (
	"encoding/hex"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type ThingsHandler struct {
	thingsService services.ThinksService
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

	resp, err := h.thingsService.ThingsConnected(userID.(string))

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

	cmd := "on"

	dataFrame := utils.AirPower(cmd)
	dataFrame2 := utils.AirPower2(cmd)
	secretKey := "SaijoDenkiSmartIOT"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"serialNumber": "2300F15050017",
		"data": map[string]string{
			"cmd": fmt.Sprintf("%x", dataFrame2),
		},
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println("Jwt Error")
		fmt.Println(err)
	}

	fmt.Println(dataFrame)
	fmt.Println(dataFrame2)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": tokenString,
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
