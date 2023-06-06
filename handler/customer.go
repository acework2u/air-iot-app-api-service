package handler

import (
	"log"
	"net/http"

	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	cusService services.CustomerService
}

func NewCustomerHandler(cusService services.CustomerService) CustomerHandler {
	return CustomerHandler{cusService}
}

func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {

	res, err := h.cusService.AllCustomers()

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": res,
	})
}

func (h *CustomerHandler) PostCustomer(ctx *gin.Context) {

	var customer *services.CreateCustomerRequest

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//
	response, err := h.cusService.CreateNewCustomer(customer)

	if err != nil {
		utils.ResponseSuccess(ctx, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "successful",
		"message": "Create Customer",
		"data":    response,
	})

}
func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {

	custId := ctx.Param("id")

	_ = custId

	//var customer *interface{}

	var customer *services.UpdateCustomer
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		// utils.StatusBadRequest("Data no match")
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  http.StatusBadRequest,
			"message": err,
		})

		return
	}

	if len(customer.Name) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "this len",
		})
		// utils.StatusBadRequest("Name is required")
		return
	}

	doc, err := h.cusService.UpdateCustomer(custId, customer)

	//doc, err := utils.ToDoc(customer)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Update Success",
		"data":    doc,
	})
}

func (h *CustomerHandler) DelCustomer(ctx *gin.Context) {
	delID := ctx.Param("id")

	if len(delID) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id is required",
		})
		return
	}

	err := h.cusService.DeleteCustomer(delID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": delID,
	})
}
