package smartapp

import (
	"errors"
	"github.com/acework2u/air-iot-app-api-service/services/smartapp"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DiagnosticBoardHandler struct {
	resp              utils.Response
	diagnosticService smartapp.DiagnosticService
}

func NewDiagnosticBoardHandler(diagService smartapp.DiagnosticService) DiagnosticBoardHandler {
	return DiagnosticBoardHandler{diagnosticService: diagService, resp: utils.Response{}}
}

func (h *DiagnosticBoardHandler) GetCheckBoard(c *gin.Context) {

	btu := c.Param("btu")
	compCode := c.Param("comeId")

	btuCheck, _ := strconv.ParseInt(btu, 10, 0)

	req := &smartapp.DiagnosticFilter{Btu: btuCheck, CompId: compCode}

	res, err := h.diagnosticService.CheckDiagnosticBoard(req)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, res)
}
func (h *DiagnosticBoardHandler) GetDiagnosticBoards(c *gin.Context) {
	res, err := h.diagnosticService.DiagnosticBoards()
	if err != nil {
		h.resp.BadRequest(c, errors.New("ไม่พบข้อมูล"))
		return
	}

	h.resp.Success(c, res)
}

func (h *DiagnosticBoardHandler) PostCheckBoard(c *gin.Context) {
	h.resp.Success(c, "OK Diagnostic Board is testing...")
}
