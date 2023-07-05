package handler

import (
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	userStr services.UserService
}

func NewUserHandler(userStr services.UserService) UserHandler {
	return UserHandler{userStr: userStr}
}

func (h *UserHandler) UserRoute(rg *gin.RouterGroup) {

	cus, err := h.userStr.GetUsers()

	if err != nil {
		log.Println(err)
		return

	}

	customers := &cus

	router := rg.Group("/users")
	router.GET("/", func(ctx *gin.Context) {
		log.Println(&cus)
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": customers})
	})

	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": "OK"})

}
