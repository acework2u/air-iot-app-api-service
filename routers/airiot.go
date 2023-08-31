package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type AirIoTRouter struct {
	airIoTHandler handler.AirIotHandler
}

func NewAirIoTRouter(airIoTHandler handler.AirIotHandler) AirIoTRouter {
	return AirIoTRouter{airIoTHandler}
}
func (r *AirIoTRouter) AirIoTRoute(rg *gin.RouterGroup) {

	router := rg.Group("air-iot", middleware.CognitoAuthMiddleware())
	router.GET("/indoors", r.airIoTHandler.GetIndoor)
	router.GET("/indoor", r.airIoTHandler.GetShadowsDoc)

	router.GET("/chk-d", r.airIoTHandler.CheckDefault)
	router.GET("/chk-p", r.airIoTHandler.CheckProduction)
	router.GET("/chk-a", r.airIoTHandler.CheckAWS)

	router.GET("/ws", r.airIoTHandler.WsIoT)

	route2 := rg.Group("ws-air", middleware.CognitoAuthMiddleware())
	route2.GET("", r.airIoTHandler.MqSubShadows)

}
