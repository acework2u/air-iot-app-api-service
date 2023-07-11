package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AddressHandler struct {
	addrService service.AddressService
}

func NewAddressHandler(addrService service.AddressService) AddressHandler {
	return AddressHandler{addrService}
}

func (h *AddressHandler) GetAddress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "address list",
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
