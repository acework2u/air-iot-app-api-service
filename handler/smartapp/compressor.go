package smartapp

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services/smartapp"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type CompressorHandler struct {
	bomService smartapp.BomService
	resp       utils.Response
}

func NewCompressorHandler(bomService smartapp.BomService) CompressorHandler {
	return CompressorHandler{bomService: bomService, resp: utils.Response{}}
}

func (h *CompressorHandler) GetCheckCompressor(c *gin.Context) {

	indoorSn := c.Param("sn")
	compInfo, err := h.bomService.CheckCompressor(indoorSn)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	h.resp.Success(c, compInfo)

}
func (h *CompressorHandler) GetCheckCompressors(c *gin.Context) {
	compInfo, err := h.bomService.CompressorList()
	if err != nil {
		fmt.Println(err)
		h.resp.BadRequest(c, err)
		return
	}

	h.resp.Success(c, compInfo)
}
