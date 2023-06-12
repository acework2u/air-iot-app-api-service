package handler

import (
	"net/http"

	services "github.com/acework2u/air-iot-app-api-service/services/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}
