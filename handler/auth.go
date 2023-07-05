package handler

import (
	"fmt"
	"net/http"

	services "github.com/acework2u/air-iot-app-api-service/services/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthenServices
}

func NewAuthHandler(authService services.AuthenServices) AuthHandler {
	return AuthHandler{authService: authService}
}

func (h *AuthHandler) PostSignIn(c *gin.Context) {
	authInput := &services.SignInRequest{}
	err := c.ShouldBindJSON(authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	res, err := h.authService.SignIn(authInput.Username, authInput.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "NotAuthorizedException: Incorrect username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}

func (h *AuthHandler) PostSignUp(c *gin.Context) {

	var authSignUp *services.SignUpRequest

	err := c.ShouldBindJSON(&authSignUp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	result, ok := h.authService.SignUp(authSignUp.Username, authSignUp.Password, authSignUp.PhoneNo)
	if ok != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ok.Error(),
		})

		return
	}

	// Add to Customer

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}

func (h *AuthHandler) PostConfirm(c *gin.Context) {

	var user *services.UserConfirm

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	// Confirm

	result, err := h.authService.UserConfirm(user.User, user.ConfirmationCode)

	if err != nil {

		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}

func (h *AuthHandler) PostResendConfirmCode(c *gin.Context) {

	var user *services.ResendConfirmCode

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	result, ok := h.authService.ResendConfirmCode(user.Username)

	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ok.Error(),
		})

		return
	}

	resultMsg := fmt.Sprintf("ใช้ Confirmation code จากอีเมล %v เพื่อยืนยันตัวตน", *result.CodeDeliveryDetails.Destination)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resultMsg,
	})
}
