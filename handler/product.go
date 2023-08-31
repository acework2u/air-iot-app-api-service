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

	serialNo := c.Param("id")

	if len(serialNo) < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "serial lte 10  is required ",
		})
		return
	}

	productInfo, err := h.productService.GetProduct(serialNo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "no documents in result",
		})

	}

	cusErr := utils.NewCustomHandler(c)
	if err != nil {

		cusErr.CustomError(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": productInfo,
	})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {

	products, err := h.productService.GetProducts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": products,
	})
}

func (h *ProductHandler) PostProduct(c *gin.Context) {

	productReq := &service.ProductNew{}
	err := c.ShouldBindJSON(productReq)
	cusErr := utils.NewCustomHandler(c)
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

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

	product := &service.ProductInfo{}
	serialNo := c.Param("id")
	err := c.ShouldBindJSON(product)
	cusErr := utils.NewCustomHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}
	// Update
	productResponse, err := h.productService.UpdateProduct(serialNo, product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": productResponse,
	})
}

func (h *ProductHandler) UpdateEWarranty(c *gin.Context) {

	cusErr := utils.NewCustomHandler(c)
	productWarranty := &service.ProductWarranty{}
	err := c.ShouldBindJSON(productWarranty)

	if err != nil {
		cusErr.CustomError(err)
		return
	}
	serialNo := productWarranty.SerialNo

	regwarranty, err := h.productService.UpdateEWarranty(serialNo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "no product in result",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": regwarranty.EWarranty,
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
