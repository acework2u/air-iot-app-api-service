package smartapp

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services/smartapp"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ErrorCodeHandler struct {
	acErrService smartapp.AcErrorService
	resp         utils.Response
}

func NewErrorCodeHandler(acErrService smartapp.AcErrorService) ErrorCodeHandler {
	return ErrorCodeHandler{resp: utils.Response{}, acErrService: acErrService}
}

func (h *ErrorCodeHandler) GetErrorCode(c *gin.Context) {

	res, err := h.acErrService.GetErrors()
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)
}

func (h *ErrorCodeHandler) GetErrorByCode(c *gin.Context) {

	errCode, _ := strconv.Atoi(c.Param("code"))

	fmt.Println("Error Code", errCode)
	res, err := h.acErrService.GetErrorByCode(errCode)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, res)
}
