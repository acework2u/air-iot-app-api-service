package handler

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
}

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
	cusAddress services.AddressService
	res        utils.Response
}

type User struct {
	Id string `json:"id" uri:"id" binding:"required"`
}

func NewCustomerHandler(cusService services.CustomerService) CustomerHandler {
	return CustomerHandler{cusService: cusService, res: utils.Response{}}
}

// GetCustomer  godoc
// @Summary  Get User info
// @Description Return User Information
// @Produce  application/json
// @Tags  Users
// @Success 200 {object} Response{}
// @Router  /my [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {

	userToken, check := c.Get("UserToken")
	userId, _ := c.Get("UserId")

	_ = userToken

	if check {
		result, err := h.cusService.CustomerById(userId.(string))

		webres := &Response{
			Status:  http.StatusBadRequest,
			Message: "No Data",
		}

		if err != nil {

			//response.Status = http.StatusBadRequest
			//response.Message = "No data"
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, webres)
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
func (h *CustomerHandler) PostCustomer(ctx *gin.Context) {

	var customer *services.CreateCustomerRequest

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.cusService.CreateNewCustomer(customer)

	if err != nil {

		msgFailerd := &utils.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}

		utils.ResponseSuccess(ctx, msgFailerd)
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

func (h *CustomerHandler) PostNewAddress(c *gin.Context) {
	userId, _ := c.Get("UserId")
	var addressInfo *services.CustomerAddress
	err := c.ShouldBindJSON(&addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	addressInfo.CustomerId = userId.(string)
	addressInfo.UpdateAt = time.Now()
	// Insert to DB
	resp, err := h.cusService.CustomerNewAddress(addressInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resp,
	})
}
func (h *CustomerHandler) UpdateAddress(c *gin.Context) {
	userId, _ := c.Get("UserId")
	addressInfo := services.CustomerAddress{}
	filter := services.Filter{}

	c.ShouldBindUri(&filter)
	err := c.ShouldBindJSON(&addressInfo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	addressInfo.CustomerId = userId.(string)
	addressInfo.UpdateAt = time.Now()

	cusAddr, err := h.cusService.CustomerUpdateAddress(&filter, &addressInfo)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	fmt.Println(cusAddr)

	h.res.Success(c, cusAddr)
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
