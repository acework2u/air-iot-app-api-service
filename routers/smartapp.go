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

type AcCompressorRouter struct {
	compHandler smartapp.CompressorHandler
}

func NewAcCompressorRouter(handler smartapp.CompressorHandler) AcCompressorRouter {
	return AcCompressorRouter{compHandler: handler}
}

func (rt *AcCompressorRouter) AcCompressorRoute(rg *gin.RouterGroup) {
	router := rg.Group("/check-compressor")
	router.GET("", rt.compHandler.GetCheckCompressors)
	router.GET("/:sn", rt.compHandler.GetCheckCompressor)

}
