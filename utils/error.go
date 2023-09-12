package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrHandler struct {
	ctx *gin.Context
}
type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func NewCustomHandler(ctx *gin.Context) ErrHandler {
	return ErrHandler{ctx: ctx}
}

func (c *ErrHandler) CustomError(err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), getErrorMsg(fe)}
		}
		c.ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})

	}
}

func (c *ErrHandler) StatusBadRequest(msg string) {
	c.ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": msg,
	})

}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "numeric":
		return "Invalid numeric"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Should be less than " + fe.Param()
	}

	return "Unknown error"
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "numeric":
		return "Invalid numeric"
	}

	return ""
}
