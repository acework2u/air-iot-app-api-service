package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrHandler struct {
	ctx *gin.Context
}

func NewCustomerHandler(ctx *gin.Context) ErrHandler {
	return ErrHandler{ctx: ctx}
}

func (c *ErrHandler) StatusBadRequest(msg string) {
	c.ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": msg,
	})

}
