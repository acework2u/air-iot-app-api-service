package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authHandler handler.AuthHandler
}

func NewAuthRouter(authHandler handler.AuthHandler) AuthController {
	return AuthController{authHandler: authHandler}
}

func (rc *AuthController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/signin", rc.authHandler.PostSignIn)
	router.POST("/signup", rc.authHandler.PostSignUp)
	router.POST("/confirm", rc.authHandler.PostConfirm)
	router.POST("/resend-confirm-code", rc.authHandler.PostResendConfirmCode)

}
