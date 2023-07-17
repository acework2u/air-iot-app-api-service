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

type ApiResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

func ResponseSuccess(c *gin.Context, msg *ApiResponse) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, msg)
}

func ResponseFailed(c *gin.Context, msg *ApiResponse) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusBadRequest, msg)
}
