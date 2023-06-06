package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/gin-gonic/gin"
)

type ClientController struct {
	clientHandler handler.ClientHandler
}

func NewClientRouter(clientHandler handler.ClientHandler) ClientController {
	return ClientController{clientHandler}
}

func (rc *ClientController) ClientRoute(rg *gin.RouterGroup) {
	router := rg.Group("/client")

	router.POST("/signup", rc.clientHandler.PostSignUp)
	router.POST("/signup/confirm", rc.clientHandler.PostConfirmSignUp)

}
