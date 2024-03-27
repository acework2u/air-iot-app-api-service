package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ErrHandler struct {
	ctx *gin.Context
}
type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		//log, handler, etc
		fmt.Println(err)
	}
	c.JSON(http.StatusInternalServerError, "")
}

func NewErrorHandler(ctx *gin.Context) ErrHandler {
	return ErrHandler{ctx: ctx}
}

func (c *ErrHandler) MyErr(err error) []*ApiError {

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]*ApiError, len(ve))
		for i, fe := range ve {
			out[i] = &ApiError{fe.Field(), getErrorMsg(fe)}
			//c.ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			//return
		}
		return out

	}
	return nil

}

func (c *ErrHandler) CustomError(err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), getErrorMsg(fe)}
		}
		c.ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "errors": out})
		return

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
		return "this field is required"
	case "email":
		return "invalid email" + fe.Param()
	case "numeric":
		return "invalid numeric"
	case "lte":
		return "should be less than " + fe.Param()
	case "gte":
		return "should be greater than " + fe.Param()
	case "min":
		return "should be less than " + fe.Param()
	case "number":
		return fmt.Sprintf("Invalid %v", fe.Field())

	}

	return "Unknown error"
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email"
	case "numeric":
		return "invalid numeric"
	}

	return ""
}
