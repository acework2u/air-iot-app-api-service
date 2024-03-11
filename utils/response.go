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
	Status  int         `json:"status,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

func (r *Response) Success(c *gin.Context, msg interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": msg,
	})
}
func (r *Response) BadRequest(c *gin.Context, msg interface{}) {
	c.Header("Content-Type", "application/json")
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": msg,
	})
}

func ResponseSuccess(c *gin.Context, msg *ApiResponse) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, msg)
}

func ResponseFailed(c *gin.Context, msg *ApiResponse) {
	c.Header("Content-Type", "application/json")
	c.AbortWithStatusJSON(http.StatusBadRequest, msg)
	//c.JSON(http.StatusBadRequest, msg)
}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}
