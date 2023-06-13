package handler

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

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
	// var oe *smithy.OperationError
	// var oe *ErrMsg

	if err := ctx.ShouldBindJSON(&clientRq); err != nil {

		fmt.Println(err)
		var result map[string]interface{}

		fmt.Println(reflect.TypeOf(err))
		//fmt.Println(reflect.TypeOf(err.Error()))
		fmt.Println(result)
		fmt.Println(err.Error())
		fmt.Println(reflect.TypeOf(err.Error()))

		// if errors.As(err, &oe) {

		// 	ctx.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  http.StatusBadRequest,
		// 		"message": fmt.Sprintf("fail to call service: %s, Operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap()),
		// 	})

		// }

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

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusNotFound,
			"message": ctx.Error(err),
		})
		log.Println(err)
		return

	}

	log.Println(res)

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

		return
	}

	//

	result, ok := h.clientService.ConfirmeSignUp(user.User, user.ConfirmationCode)

	if ok != nil {

		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": ok.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}

func (h *ClientHandler) PostSignIn(ctx *gin.Context) {

	var user *service.ClientSignUp

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}
	result, resOut, ok := h.clientService.SignIn(user.Email, user.Password)

	// fmt.Println(result)
	// fmt.Println(ok)

	if ok != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"status":  http.StatusNoContent,
			"message": ok.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
		"token":   resOut.AuthenticationResult.IdToken,
	})

}
