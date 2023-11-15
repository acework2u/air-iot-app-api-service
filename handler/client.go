package handler

import (
	"fmt"
	service "github.com/acework2u/air-iot-app-api-service/services/clientcoginto"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ClientHandler struct {
	clientService service.ClientCognito
	resp          utils.Response
}

type ErrMsg struct {
	Key   string `json:"key"`
	Error string `json:"error"`
}

func NewClientHandler(clientService service.ClientCognito) ClientHandler {
	return ClientHandler{clientService: clientService, resp: utils.Response{}}
}

func (h *ClientHandler) PostSignUp(c *gin.Context) {

	// var clientRq *service.ClientSignUp
	var clientRq *service.SignUpRequest
	// var oe *smithy.OperationError
	// var oe *ErrMsg

	if err := c.ShouldBindJSON(&clientRq); err != nil {

		//fmt.Println(err)
		//var result map[string]interface{}
		//
		//fmt.Println(reflect.TypeOf(err))
		////fmt.Println(reflect.TypeOf(err.Error()))
		//fmt.Println(result)
		//fmt.Println(err.Error())
		//fmt.Println(reflect.TypeOf(err.Error()))

		// if errors.As(err, &oe) {

		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  http.StatusBadRequest,
		// 		"message": fmt.Sprintf("fail to call service: %s, Operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap()),
		// 	})

		// }
		h.resp.BadRequest(c, err.Error())
		return
	}

	// client SignUp
	// res, err := h.clientService.SignUp(clientRq.Email, clientRq.Password)
	res, err := h.clientService.SignUp(clientRq.Email, clientRq.Password)
	if err != nil {
		h.resp.BadRequest(c, c.Error(err))
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusNotFound,
		//	"message": c.Error(err),
		//})
		//log.Println(err)
		return

	}

	log.Println(res)
	h.resp.Success(c, fmt.Sprintf("emil %v and Pass %v", clientRq.Email, clientRq.Password))
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": fmt.Sprintf("emil %v and Pass %v", clientRq.Email, clientRq.Password),
	//})
}

func (h *ClientHandler) PostConfirmSignUp(c *gin.Context) {

	var user *service.UserConfirm

	if err := c.ShouldBindJSON(&user); err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	result, ok := h.clientService.ConfirmSignUp(user.User, user.ConfirmationCode)

	if ok != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": ok.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}

func (h *ClientHandler) PostSignIn(c *gin.Context) {

	var user *service.ClientSignUp

	if err := c.ShouldBindJSON(&user); err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	result, resOut, ok := h.clientService.SignIn(user.Email, user.Password)

	if ok != nil {
		h.resp.BadRequest(c, ok.Error())
		return
	}
	//Success
	h.resp.Success(c, result)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
		"token":   resOut.AuthenticationResult.IdToken,
	})

}
