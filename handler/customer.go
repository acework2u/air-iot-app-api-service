package handler

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserInfo struct {
	Username            string `json:"username" structs:"username"`
	Email               string `json:"email" structs:"email"`
	EmailVerified       bool   `json:"emailVerified" json:"emailVerified"`
	Exp                 string `json:"exp" structs:"exp"`
	Iat                 string `json:"iat" structs:"iat"`
	Iss                 string `json:"iss" structs:"iss"`
	Jti                 string `json:"jti" structs:"jti"`
	OriginJti           string `json:"origin_jti" structs:"origin_jti"`
	PhoneNumber         string `json:"phone_number" structs:"phone_number"`
	PhoneNumberVerified bool   `json:"phone_number_verified" structs:"phone_number_verified"`
	Sub                 string `json:"sub" structs:"sub"`
	TokenUse            string `json:"token_use" structs:"token_use"`
}

type CustomerHandler struct {
	cusService services.CustomerService
}

type User struct {
	Id string `json:"id" uri:"id" binding:"required"`
}

func NewCustomerHandler(cusService services.CustomerService) CustomerHandler {
	return CustomerHandler{cusService}
}
func (h *CustomerHandler) GetCustomer(c *gin.Context) {

	userToken, check := c.Get("UserToken")
	userId, _ := c.Get("UserId")

	_ = userToken

	if check {
		result, err := h.cusService.CustomerById(userId.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "no data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": result,
		})

	}

}

func (h *CustomerHandler) GetCustomerById(c *gin.Context) {

	ReqId := c.Param("id")
	userId, ok := c.Get("UserId")

	if ok {
		if ReqId == userId {
			result, err := h.cusService.CustomerById(userId.(string))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "no data",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": result,
			})

			return

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "id is required",
		})
		return
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
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {

	cusID := c.Query("id")
	userId, ok := c.Get("UserId")

	infoUpdate := services.UpdateInfoRequest{}

	if ok {

		if cusID == userId {

			if err := c.ShouldBindJSON(&infoUpdate); err != nil {

				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": err.Error(),
				})
				return
			}
			infoUpdate.UpdateAt = time.Now()

			updateId := userId.(string)

			resInfo, ok := h.cusService.UpdateCustomer(updateId, &infoUpdate)

			if ok != nil {
				fmt.Println("In handler error")
				fmt.Println(ok.Error())
			}

			_ = resInfo

			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": infoUpdate,
			})

		} else {

			// userId do not match
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Access don't allowed",
			})
			return
		}

	}

	//
	//_ = custId
	//
	////var customer *interface{}
	//
	//var customer *services.UpdateCustomer
	//err := ctx.ShouldBindJSON(&customer)
	//if err != nil {
	//	// utils.StatusBadRequest("Data no match")
	//	ctx.JSON(http.StatusBadGateway, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err,
	//	})
	//
	//	return
	//}
	//
	//if len(customer.Name) == 0 {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": "this len",
	//	})
	//	// utils.StatusBadRequest("Name is required")
	//	return
	//}
	//
	//doc, err := h.cusService.UpdateCustomer(custId, customer)
	//
	////doc, err := utils.ToDoc(customer)
	//
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//ctx.JSON(http.StatusAccepted, gin.H{
	//	"status":  http.StatusAccepted,
	//	"message": "Update Success",
	//	"data":    doc,
	//})
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

//func AcceptAnyValue(arg any) {
//	switch v := arg.(type) {
//	case string:
//		fmt.Printf("String: %s", v)
//	case int:
//		fmt.Printf("Int32: %d", v)
//	case float64:
//		fmt.Printf("float64: %f", v)
//	case map[string]int:
//		fmt.Printf("map[string]int: %+v", v)
//	case map[int]string:
//		fmt.Printf("map[int]string: %+v", v)
//	case map[string]map[any]any:
//		fmt.Printf("map[string]map[any]any: %+v", v)
//	case []int:
//		fmt.Printf("[]int: %+v", v)
//	default:
//		fmt.Printf("Undefined type: %s", reflect.TypeOf(v))
//	}
//
//	fmt.Println()
//}
