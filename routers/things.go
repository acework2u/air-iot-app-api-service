package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/gin-gonic/gin"
)

type ThingController struct {
	thingsHandler handler.ThingsHandler
}

func NewThingsRouter(thingsHandler handler.ThingsHandler) ThingController {
	return ThingController{thingsHandler}
}

func (rc *ThingController) ThingsRoute(rg *gin.RouterGroup) {
	router := rg.Group("/iot")

	router.GET("/things", rc.thingsHandler.ConnectThing)
	router.POST("/things", rc.thingsHandler.ConnectThing)
}
