package handler

import (
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username              string `json:"username" structs:"username"`
	Email                 string `json:"email" structs:"email"`
	Email_verified        bool   `json:"email_verifyed" json:"email_verifyed"`
	Exp                   string `json:"exp" structs:"exp"`
	Iat                   string `json:"iat" structs:"iat"`
	Iss                   string `json:"iss" structs:"iss"`
	Jti                   string `json:"jti" structs:"jti"`
	Origin_jti            string `json:"origin_jti" structs:"origin_jti"`
	Phone_number          string `json:"phone_number" structs:"phone_number"`
	Phone_number_verified bool   `json:"phone_number_verified" structs:"phone_number_verified"`
	Sub                   string `json:"sub" structs:"sub"`
	Token_use             string `json:"token_use" structs:"token_use"`
}

type CustomerHandler struct {
	cusService services.CustomerService
}

func NewCustomerHandler(cusService services.CustomerService) CustomerHandler {
	return CustomerHandler{cusService}
}

func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {

	userToken, check := ctx.Get("UserToken")

	userName, _ := ctx.Get("UserId")

	if check {

		res, err := h.cusService.AllCustomers()

		fmt.Println("userName")
		fmt.Println(userName)

		// fmt.Sprintf("format string %T", userToken)

		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})

			return
		}

		_ = res

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": userToken,
		})

	}

}

// CreateCustomer godoc
// @Summary Create Customer
// @Description Save customer data in DB
// @Param customer body services.CreateCustomerRequest true "Create Customer"
// @Produce application/json
// @Tags customers
// @Success 200 {object} response{}
// @Router /customers [get]

func (h *CustomerHandler) PostCustomer(ctx *gin.Context) {

	var customer *services.CreateCustomerRequest

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

func AcceptAnyValue(arg any) {
	switch v := arg.(type) {
	case string:
		fmt.Printf("String: %s", v)
	case int:
		fmt.Printf("Int32: %d", v)
	case float64:
		fmt.Printf("float64: %f", v)
	case map[string]int:
		fmt.Printf("map[string]int: %+v", v)
	case map[int]string:
		fmt.Printf("map[int]string: %+v", v)
	case map[string]map[any]any:
		fmt.Printf("map[string]map[any]any: %+v", v)
	case []int:
		fmt.Printf("[]int: %+v", v)
	default:
		fmt.Printf("Undefined type: %s", reflect.TypeOf(v))
	}

	fmt.Println()
}
