package handler

import (
	"net/http"

	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/gin-gonic/gin"
)

type ThingsHandler struct {
	thingsService services.ThinksService
}

func NewThingsHandler(thingsService services.ThinksService) ThingsHandler {
	return ThingsHandler{thingsService}
}

func (h *ThingsHandler) ConnectThing(c *gin.Context) {

	res, err := h.thingsService.GetCerds()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}

func (h *ThingsHandler) ThingsCert(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Working ThingCert",
	})

	//var userLogin *services.UserReq
	//
	//err := c.ShouldBindJSON(&userLogin)
	//
	//file, err := c.FormFile("image")
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//fmt.Println(file)

	//h.thingsService.UploadFile(file)

	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//resP, err := h.thingsService.GetUserCert(userLogin)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//_ = resP
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": resP,
	//})
}

func (h *ThingsHandler) UploadFile(c *gin.Context) {

	file, err := c.FormFile("image")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	res, err := h.thingsService.UploadToS3(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}
