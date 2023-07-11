package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	cusHandler handler.CustomerHandler
}

func NewCustomerRouter(cusHandler handler.CustomerHandler) CustomerController {
	return CustomerController{cusHandler}

}

func (rc *CustomerController) CustomerRoute(rg *gin.RouterGroup) {
	router := rg.Group("/my", middleware.CognitoAuthMiddleware())
	router.GET("", rc.cusHandler.GetCustomer)
	router.GET("/:id", rc.cusHandler.GetCustomerById)
	//router.POST("/", rc.cusHandler.PostCustomer)
	router.POST("/info", rc.cusHandler.UpdateCustomer)
	router.POST("/address", rc.cusHandler.PostNewAddress)
	//router.DELETE("/:id", rc.cusHandler.DelCustomer)
}
