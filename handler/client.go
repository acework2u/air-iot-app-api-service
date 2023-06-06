package handler

import (
	"fmt"
	"log"
	"net/http"

	service "github.com/acework2u/air-iot-app-api-service/services/clientcoginto"
	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	clientService service.ClientCognito
}

type ErrMsg struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

func NewClientHandler(clientService service.ClientCognito) ClientHandler {
	return ClientHandler{clientService: clientService}
}

func (h *ClientHandler) PostSignUp(ctx *gin.Context) {

	var clientRq *service.ClientSignUp

	err := ctx.ShouldBindJSON(&clientRq)
	if err != nil {

		// if er:=err();er!=nil {

		// }

		// er := err.Error()

		// // s := string(fmt.Sprintf(`$v`, er))
		// data := ErrMsg{}

		// er = map[string]interface{}{er}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	// client SignUp
	// res, err := h.clientService.SignUp(clientRq.Email, clientRq.Password)

	res, err := h.clientService.SignUp(clientRq.Email, clientRq.Password)
	if err != nil {

		log.Println(err)
		return

	}

	log.Println(res)

	// if err != nil {
	// 	log.Println(err)
	// 	//panic(err)

	// 	return
	// }

	// log.Panicln(res)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("emil %v and Pass %v", clientRq.Email, clientRq.Password),
	})
}

func (h *ClientHandler) PostConfirmSignUp(ctx *gin.Context) {

	var user *service.UserConfirm

	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Confirm",
	})
}
