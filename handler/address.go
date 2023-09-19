package handler

import (
	"fmt"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AddressHandler struct {
	addrService service.AddressService
	res         utils.Response
}

func NewAddressHandler(addrService service.AddressService) AddressHandler {
	return AddressHandler{addrService: addrService, res: utils.Response{}}
}

func (h *AddressHandler) GetAddress(c *gin.Context) {

	userID, _ := c.Get("UserId")
	if len(userID.(string)) < 1 {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong!",
		})
		return
	}

	resData, err := h.addrService.AllAddress(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "sorry, i can't this service",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resData,
	})
}

func (h *AddressHandler) PostNewAddress(c *gin.Context) {

	var addressInfo *service.CustomerAddress

	//userToken, check := c.Get("UserToken")
	userId, _ := c.Get("UserId")

	if len(userId.(string)) > 0 {

		err := c.ShouldBindJSON(&addressInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		addressInfo.CustomerId = userId.(string)
		addressInfo.UpdateAt = time.Now()

		resAddress, err := h.addrService.NewAddress(addressInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": resAddress,
		})

	}

}

func (h *AddressHandler) UpdateAddress(c *gin.Context) {

	h.res.Success(c, "update my address")
}

func (h *AddressHandler) DeleteAddress(c *gin.Context) {

	addId := c.Param("id")

	err := h.addrService.DelAddress(addId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Delete Address ID %v Successful", addId),
	})
}
