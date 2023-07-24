package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type AirThingRouter struct {
	airThingHandler handler.AirThingHandler
}

func NewAirThingRouter(airThingHandler handler.AirThingHandler) AirThingRouter {
	return AirThingRouter{airThingHandler}
}

func (r *AirThingRouter) AirThingRoute(rg *gin.RouterGroup) {

	router := rg.Group("airthings", middleware.CognitoAuthMiddleware())
	router.GET("", r.airThingHandler.GetCerts)
	router.GET("/connect", r.airThingHandler.Connect)

}
