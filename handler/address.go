package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
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

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "New Address",
	})
}
