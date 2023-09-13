package handler

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/utils"
	"net/http"
	"strings"

	services "github.com/acework2u/air-iot-app-api-service/services/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthenServices
	resp        utils.Response
}

func NewAuthHandler(authService services.AuthenServices) AuthHandler {
	return AuthHandler{authService: authService, resp: utils.Response{}}
}

// Authenticate godoc
// @Summary Air IoT User Authentication
// @Description Authenticates a user and provides a JWT to Authorize API Calls
// @ID Authentication
// @Tags Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "User Credentials"
// @Param password formData string true "User Credentials"
// @Success 200 {object} utils.ApiResponse
// @Failure 401 {object} utils.ApiResponse
// @Failure 400 {object} utils.ApiResponse
// @Router /auth/signin [post]
func (h *AuthHandler) PostSignIn(c *gin.Context) {
	authInput := &services.SignInRequest{}
	err := c.ShouldBindJSON(authInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	res, err := h.authService.SignIn(authInput.Username, authInput.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "NotAuthorizedException: Incorrect username or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res.AuthenticationResult,
	})
}

func (h *AuthHandler) PostSignUp(c *gin.Context) {

	var authSignUp *services.SignUpRequest

	err := c.ShouldBindJSON(&authSignUp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	result, ok := h.authService.SignUp(authSignUp.Username, authSignUp.Password, authSignUp.PhoneNo)
	if ok != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ok.Error(),
		})

		return
	}

	// Add to Customer

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}

func (h *AuthHandler) PostConfirm(c *gin.Context) {

	var user *services.UserConfirm

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	// Confirm
	result, err := h.authService.UserConfirm(user.User, user.ConfirmationCode)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid verification code provided, please try again.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": result,
	})
}

func (h *AuthHandler) PostResendConfirmCode(c *gin.Context) {

	var user *services.ResendConfirmCode

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	result, ok := h.authService.ResendConfirmCode(user.Username)

	if ok != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ok.Error(),
		})

		return
	}

	resultMsg := fmt.Sprintf("ใช้ Confirmation code จากอีเมล %v เพื่อยืนยันตัวตน", *result.CodeDeliveryDetails.Destination)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resultMsg,
	})
}

func (h *AuthHandler) PostRefreshToken(c *gin.Context) {

	refreshToken := &services.SignInResponse{}
	custErr := utils.NewCustomHandler(c)
	err := c.ShouldBindJSON(refreshToken)
	if err != nil {
		custErr.CustomError(err)
		return
	}
	refToken := *refreshToken.RefreshToken
	resToken, ok := h.authService.RefreshToken(refToken)
	if ok != nil {

		reshErr := strings.Split(ok.Error(), ",")
		reshErr = strings.Split(reshErr[3], ":")
		if len(reshErr[1]) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": reshErr[1],
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": ok.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resToken,
	})

	return
}

func (h *AuthHandler) PostForgotPw(c *gin.Context) {

	userName := services.ResendConfirmCode{}
	c.ShouldBindJSON(&userName)

	if len(userName.Username) < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User is required",
		})
		return
	}

	response, err := h.authService.ForgotPassword(userName.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	sendToEmail := *response.CodeDeliveryDetails.Destination
	txtResponse := fmt.Sprintf("ระบบได้ส่ง Password Reset Code ไปที่ %v : %v กรุณาใช้ Code เพื่อยืนยัน", response.CodeDeliveryDetails.DeliveryMedium, sendToEmail)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": txtResponse,
	})
}

func (h *AuthHandler) PostConfirmNewPassword(c *gin.Context) {

	confirmReq := services.UserConfirmNewPassword{}
	customErr := utils.NewCustomHandler(c)
	err := c.ShouldBindJSON(&confirmReq)

	if err != nil {
		customErr.CustomError(err)
		return
	}

	response, err := h.authService.ConfirmNewPassword(&confirmReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": response,
	})
}

func (h *AuthHandler) PostChangePassword(c *gin.Context) {

	changePasswordReq := services.ChangePasswordReq{}
	err := c.ShouldBindJSON(&changePasswordReq)
	custErr := utils.NewCustomHandler(c)
	if err != nil {
		custErr.CustomError(err)
		return
	}

	resChange, err := h.authService.ChangePassword(&changePasswordReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": resChange,
	})
}

func (h *AuthHandler) DelCustomer(c *gin.Context) {
	acsessToken := services.UserDelete{}
	err := c.ShouldBindJSON(&acsessToken)
	utils.NewCustomHandler(c)

	if err != nil {
		txtErr := strings.Split(err.Error(), ":")
		if len(txtErr) > 0 {
			h.resp.BadRequest(c, txtErr[len(txtErr)-1])
			return
		}
		h.resp.BadRequest(c, err.Error())
		return
	}

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}
	h.resp.Success(c, "Delete your account is a complete")
	//h.resp.Success(c, "Delete Customer")
}
