package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get",
	})
}

func (h *ProductHandler) PostProduct(c *gin.Context) {

	productReq := &service.ProductNew{}
	err := c.ShouldBindJSON(productReq)
	cusErr := utils.NewCustomerHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}

	res, err := h.productService.CreateProduct(productReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}

func (h *ProductHandler) DelProduct(c *gin.Context) {

	serial := c.Param("id")

	if len(serial) > 0 {
		err := h.productService.DeleteProduct(serial)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "delete successful",
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": "serial no is wrong",
	})

}
