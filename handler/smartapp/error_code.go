package smartapp

import (
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type ErrorCodeHandler struct {
	resp utils.Response
}

func NewErrorCodeHandler() ErrorCodeHandler {
	return ErrorCodeHandler{resp: utils.Response{}}
}

func (h *ErrorCodeHandler) GetErrorCode(c *gin.Context) {

	h.resp.Success(c, "Error Code Handler OK")
}
