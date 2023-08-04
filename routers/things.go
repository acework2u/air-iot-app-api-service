package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	"github.com/gin-gonic/gin"
)

type ThingController struct {
	thingsHandler handler.ThingsHandler
}

func NewThingsRouter(thingsHandler handler.ThingsHandler) ThingController {
	return ThingController{thingsHandler}
}

func (rc *ThingController) ThingsRoute(rg *gin.RouterGroup) {
	router := rg.Group("/iot", middleware.CognitoAuthMiddleware())

	router.GET("/user/certs", rc.thingsHandler.UserCert)
	router.POST("/upload", rc.thingsHandler.UploadFile)
	router.GET("/thing/register", rc.thingsHandler.ThingsRegister)
	router.GET("/thing/connected", rc.thingsHandler.ThingConnect)
	router.GET("/thing/cert", rc.thingsHandler.ThingsCert)
	router.GET("/thing/cmd", rc.thingsHandler.CmdThing)
	router.GET("/thing/shadows", rc.thingsHandler.Shadows)

	//router.GET("/things", rc.thingsHandler.ConnectThing)
	//router.POST("/things", rc.thingsHandler.ConnectThing)
}
