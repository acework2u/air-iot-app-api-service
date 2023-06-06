package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	cusHandler handler.CustomerHandler
}

func NewCustomerRouter(cusHandler handler.CustomerHandler) CustomerController {
	return CustomerController{cusHandler}

}

func (rc *CustomerController) CustRoute(rg *gin.RouterGroup) {
	router := rg.Group("/customers")
	router.GET("/", rc.cusHandler.GetCustomer)
	router.POST("/", rc.cusHandler.PostCustomer)
	router.PUT("/:id", rc.cusHandler.UpdateCustomer)
	router.DELETE("/:id", rc.cusHandler.DelCustomer)
}
