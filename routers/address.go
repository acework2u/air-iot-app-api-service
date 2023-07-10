package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/gin-gonic/gin"
)

type AddressController struct {
	addrHandler handler.AddressHandler
}

func NewAddressRouter(addrHandler handler.AddressHandler) AddressController {
	return AddressController{addrHandler}
}

func (ra *AddressController) AddressRoute(rg *gin.RouterGroup) {

	router := rg.Group("/address")
	router.GET("", ra.addrHandler.GetAddress)
	router.POST("/", ra.addrHandler.PostNewAddress)

}
