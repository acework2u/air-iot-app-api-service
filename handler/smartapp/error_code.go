package smartapp

import (
	"github.com/acework2u/air-iot-app-api-service/services/smartapp"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type ErrorCodeHandler struct {
	acErrService smartapp.AcErrorService
	resp         utils.Response
}

func NewErrorCodeHandler(acErrService smartapp.AcErrorService) ErrorCodeHandler {
	return ErrorCodeHandler{resp: utils.Response{}, acErrService: acErrService}
}

func (h *ErrorCodeHandler) GetErrorCode(c *gin.Context) {
	id := 2
	res, err := h.acErrService.GetErrorByCode(id)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)

	//h.resp.Success(c, "Error Code Handler OK")
}

func (h *ErrorCodeHandler) GetErrorByCode(c *gin.Context) {

}
