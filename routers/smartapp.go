package routers

import (
	"github.com/acework2u/air-iot-app-api-service/handler/smartapp"
	"github.com/gin-gonic/gin"
)

type AcErrorCodeRouter struct {
	errorCodeHandler smartapp.ErrorCodeHandler
}

func NewAcErrorCodeRouter(errorCodeHandler smartapp.ErrorCodeHandler) AcErrorCodeRouter {
	return AcErrorCodeRouter{errorCodeHandler: errorCodeHandler}
}

func (rt *AcErrorCodeRouter) ErrorCodeRoute(rg *gin.RouterGroup) {
	router := rg.Group("/error-code")
	router.GET("", rt.errorCodeHandler.GetErrorCode)
	router.GET("/:code", rt.errorCodeHandler.GetErrorByCode)

}
