package handler

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/services"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"github.com/gin-gonic/gin"
	"log"
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
// @Tags  Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router  /my [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {

	_, check := c.Get("UserToken")
	userId, _ := c.Get("UserId")

	if check {
		result, err := h.cusService.CustomerById(userId.(string))
		if err != nil {
			h.res.BadRequest(c, "no data")
			return
		}
		h.res.Success(c, result)

	}

}

// GetCustomerById godoc
// @Summary get user information by id
// @Description get user information by id
// @Tags Users
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      int  true  "Account ID"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router  /my/{id} [get]
func (h *CustomerHandler) GetCustomerById(c *gin.Context) {

	ReqId := c.Param("id")
	userId, ok := c.Get("UserId")

	if ok {
		if ReqId == userId {
			result, err := h.cusService.CustomerById(userId.(string))
			if err != nil {
				h.res.BadRequest(c, "no data")
				return
			}
			h.res.Success(c, result)
			return

		}
		h.res.BadRequest(c, "field user id is required")
		return
	}
}

// NewCustomerById godoc
// @Summary new user information
// @Description New user information
// @Tags Users
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      int  true  "Account ID"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router  /my [post]

func (h *CustomerHandler) PostCustomer(c *gin.Context) {

	var customer *services.CreateCustomerRequest

	if err := c.ShouldBindJSON(&customer); err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	response, err := h.cusService.CreateNewCustomer(customer)

	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	h.res.Success(c, response)
}

// UpdateCustomer godoc
// @Summary update user information
// @Description post method user information update
// @Tags Users
// @Security BearerAuth
// @Param userInfo body services.UpdateInfoRequest true "User information"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /my/info [post]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {

	cusID := c.Query("id")
	userId, ok := c.Get("UserId")
	infoUpdate := &services.UpdateInfoRequest{}
	if ok {
		if cusID == userId {
			// if err := c.ShouldBindJSON(infoUpdate); err != nil {
			//h.res.BadRequest(c, err.Error())
			// return
			// }

			err := c.ShouldBindJSON(infoUpdate)

			cusErr := utils.NewErrorHandler(c)
			if err != nil {
				cusErr.CustomError(err)
				return
			}

			infoUpdate.UpdateAt = time.Now()
			updateId := userId.(string)
			_, ok := h.cusService.UpdateCustomer(updateId, infoUpdate)

			if ok != nil {
				h.res.BadRequest(c, ok.Error())
				return
			}
			h.res.Success(c, infoUpdate)

		} else {
			// userId do not match
			h.res.BadRequest(c, "Access don't allowed")
			return
		}

	}

}

// DelCustomer godoc
// @Summary Delete user by id
// @Description Delete user by id
// @Tags Users
// @Security BearerAuth
// @Param        id   path      int  true  "Account ID"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /my/{id} [delete]
func (h *CustomerHandler) DelCustomer(c *gin.Context) {
	delID := c.Param("id")

	if len(delID) == 0 {
		h.res.BadRequest(c, "field id is required")
		return
	}
	err := h.cusService.DeleteCustomer(delID)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	txtDel := fmt.Sprintf("delete is id %v successful", delID)
	h.res.Success(c, txtDel)

}

// PostNewAddress godoc
// @Summary PostNewAddress user by id
// @Description PostNewAddress user by id
// @Tags Users
// @Security BearerAuth
// @Param CustomerAddress body services.CustomerAddress true "Customer Address"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /my/address [post]
func (h *CustomerHandler) PostNewAddress(c *gin.Context) {
	userId, _ := c.Get("UserId")
	var addressInfo *services.CustomerAddress
	err := c.ShouldBindJSON(&addressInfo)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	addressInfo.CustomerId = userId.(string)
	addressInfo.UpdateAt = time.Now()
	// Insert to DB
	resp, err := h.cusService.CustomerNewAddress(addressInfo)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	h.res.Success(c, resp)
}

// UpdateAddress godoc
// @Summary Update user address
// @Description Update
// @Tags Users
// @Produce json
// @Param        id   path      int  true  "Account ID"
// @Param address body services.CustomerAddress true "Address information"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /my/address/:id [put]
func (h *CustomerHandler) UpdateAddress(c *gin.Context) {

	userId, _ := c.Get("UserId")
	addressInfo := services.CustomerAddress{}
	filter := services.Filter{}

	c.ShouldBindUri(&filter)
	err := c.ShouldBindJSON(&addressInfo)

	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}

	addressInfo.CustomerId = userId.(string)
	addressInfo.UpdateAt = time.Now()
	cusAddr, err := h.cusService.CustomerUpdateAddress(&filter, &addressInfo)
	if err != nil {
		h.res.BadRequest(c, err.Error())
		return
	}
	log.Println(cusAddr)
	//fmt.Println(cusAddr)
	h.res.Success(c, cusAddr)
}
