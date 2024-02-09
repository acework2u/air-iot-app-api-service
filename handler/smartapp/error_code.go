package smartapp

import (
	"github.com/acework2u/air-iot-app-api-service/services/smartapp"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type ErrorCodeHandler struct {
	//acErrService smartapp.AcErrorService
	acErrService smartapp.ErrorCodeService
	resp         utils.Response
}

func NewErrorCodeHandler(acErrService smartapp.ErrorCodeService) ErrorCodeHandler {
	return ErrorCodeHandler{resp: utils.Response{}, acErrService: acErrService}
}

func (h *ErrorCodeHandler) GetErrorCode(c *gin.Context) {

	res, err := h.acErrService.ErrorCodeList()
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)
}

func (h *ErrorCodeHandler) GetErrorByCode(c *gin.Context) {
	errCode := c.Param("code")
	res, err := h.acErrService.ErrorByCode(errCode)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)
}
