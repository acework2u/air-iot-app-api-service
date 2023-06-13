package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

func ResponseSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"message": msg,
	})
}

func ResponseFail(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": msg,
	})
}
