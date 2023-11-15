package handler

import (
	"fmt"
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type AddressHandler struct {
	addrService service.AddressService
	resp        utils.Response
}

func NewAddressHandler(addrService service.AddressService) AddressHandler {
	return AddressHandler{addrService: addrService, resp: utils.Response{}}
}

func (h *AddressHandler) GetAddress(c *gin.Context) {

	userID, _ := c.Get("UserId")
	if len(userID.(string)) < 1 {
		h.resp.BadRequest(c, "something wrong")
		return
	}

	resData, err := h.addrService.AllAddress(userID.(string))
	if err != nil {
		h.resp.BadRequest(c, "sorry, i can't this service")
		return
	}
	//
	h.resp.Success(c, resData)
}

func (h *AddressHandler) PostNewAddress(c *gin.Context) {

	var addressInfo *service.CustomerAddress

	//userToken, check := c.Get("UserToken")
	userId, _ := c.Get("UserId")

	if len(userId.(string)) > 0 {

		err := c.ShouldBindJSON(&addressInfo)
		if err != nil {
			h.resp.BadRequest(c, err.Error())
			return
		}
		addressInfo.CustomerId = userId.(string)
		addressInfo.UpdateAt = time.Now()

		resAddress, err := h.addrService.NewAddress(addressInfo)
		if err != nil {
			h.resp.BadRequest(c, err.Error())
			return
		}

		h.resp.Success(c, resAddress)

	}

}

func (h *AddressHandler) UpdateAddress(c *gin.Context) {

	h.resp.Success(c, "update my address")
}

func (h *AddressHandler) DeleteAddress(c *gin.Context) {

	addId := c.Param("id")

	err := h.addrService.DelAddress(addId)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Success
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": fmt.Sprintf("Delete Address ID %v Successful", addId),
	//})
	h.resp.Success(c, fmt.Sprintf("Delete Address ID %v Successful", addId))

}
