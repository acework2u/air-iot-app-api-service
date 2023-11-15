package handler

import (
	service "github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
	resp           utils.Response
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return ProductHandler{productService: productService, resp: utils.Response{}}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

	serialNo := c.Param("id")

	if len(serialNo) < 0 {
		h.resp.BadRequest(c, "serial lte 10  is required")
		return
	}

	productInfo, err := h.productService.GetProduct(serialNo)

	if err != nil {
		h.resp.BadRequest(c, "no document in result")
		return
	}

	cusErr := utils.NewErrorHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return
	}
	// Success
	h.resp.Success(c, productInfo)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {

	products, err := h.productService.GetProducts()
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Success
	h.resp.Success(c, products)
}

func (h *ProductHandler) PostProduct(c *gin.Context) {

	productReq := &service.ProductNew{}
	err := c.ShouldBindJSON(productReq)
	cusErr := utils.NewErrorHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}

	res, err := h.productService.CreateProduct(productReq)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Success
	h.resp.Success(c, res)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {

	product := &service.ProductInfo{}
	serialNo := c.Param("id")
	err := c.ShouldBindJSON(product)
	cusErr := utils.NewErrorHandler(c)
	if err != nil {
		cusErr.CustomError(err)
		return

	}
	// Update
	productResponse, err := h.productService.UpdateProduct(serialNo, product)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	// Success
	h.resp.Success(c, productResponse)
}

func (h *ProductHandler) UpdateEWarranty(c *gin.Context) {

	cusErr := utils.NewErrorHandler(c)
	productWarranty := &service.ProductWarranty{}
	err := c.ShouldBindJSON(productWarranty)

	if err != nil {
		cusErr.CustomError(err)
		return
	}
	serialNo := productWarranty.SerialNo

	regWarranty, err := h.productService.UpdateEWarranty(serialNo)

	if err != nil {
		h.resp.BadRequest(c, "no product in result")
		return
	}
	// Success
	h.resp.Success(c, regWarranty.EWarranty)
}

func (h *ProductHandler) DelProduct(c *gin.Context) {

	serial := c.Param("id")

	if len(serial) > 0 {
		err := h.productService.DeleteProduct(serial)
		if err != nil {
			h.resp.BadRequest(c, err.Error())
			return
		}
		//Success
		h.resp.Success(c, "delete successful")
		return
	}
	h.resp.BadRequest(c, "serial number id wrong")
}
