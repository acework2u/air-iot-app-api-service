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
// @Summary User signin
// @Description Authenticates a user and provides authorize API Calls
// @ID Authentication
// @Tags Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Schemes https http
// @Param SignIn body services.SignInRequest true "User and Password "
// @Success 200 {object} utils.ApiResponse{}
// @Success 400 {object} utils.ApiResponse{}
// @Router /auth/signin [post]
func (h *AuthHandler) PostSignIn(c *gin.Context) {
	authInput := &services.SignInRequest{}
	authResponse := &utils.ApiResponse{}

	err := c.ShouldBindJSON(authInput)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	res, err := h.authService.SignIn(authInput.Username, authInput.Password)

	if err != nil {

		errText := utils.ApiResponse{Status: http.StatusBadRequest, Message: "Incorrect username or password."}
		h.resp.BadRequest(c, errText)
		return
	}

	authResponse = &utils.ApiResponse{
		Status:  http.StatusOK,
		Message: res.AuthenticationResult,
	}

	h.resp.Success(c, authResponse)

}

// Authenticate godoc
// @Summary User Sign Up
// @Description User SignUp for use a Air IoT resource
// @Produce application/json
// @Tags Authentication
// @Param SignUp body services.SignUpRequest true "New User information"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /auth/signup [post]
func (h *AuthHandler) PostSignUp(c *gin.Context) {

	var authSignUp *services.SignUpRequest

	err := c.ShouldBindJSON(&authSignUp)

	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	result, ok := h.authService.SignUp(authSignUp.Username, authSignUp.Password, authSignUp.PhoneNo)
	if ok != nil {
		txtErr := strings.Split(ok.Error(), ",")
		h.resp.BadRequest(c, fmt.Sprintf("%s", txtErr[len(txtErr)-1]))
		return
	}
	// Add to Customer
	apiResult := utils.ApiResponse{Status: http.StatusOK, Message: result}
	c.JSON(http.StatusOK, apiResult)
}

// Authenticate godoc
// @Summary User confirm is sign up
// @Description New user confirm a sign up
// @Produce application/json
// @Tags Authentication
// @Param SignUp body services.UserConfirm true "New User confirm information"
// @Success 200 {object} utils.ApiResponse{}
// @Success 400 {object} utils.ApiResponse{}
// @Router /auth/confirm [post]
func (h *AuthHandler) PostConfirm(c *gin.Context) {

	var user *services.UserConfirm

	if err := c.ShouldBindJSON(&user); err != nil {

		h.resp.BadRequest(c, err.Error())
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})

		return
	}

	// Confirm
	result, err := h.authService.UserConfirm(user.User, user.ConfirmationCode)

	if err != nil {

		resErr := "Invalid verification code provided, please try again."
		h.resp.BadRequest(c, resErr)
		return
	}

	apiResponse := utils.ApiResponse{
		Status:  http.StatusOK,
		Message: result,
	}

	h.resp.Success(c, apiResponse)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": result,
	//})
}

// Authenticate godoc
// @Summary Resend confirm code for a new user
// @Description retern Resend confirmation code for a new user
// @Produce application/json
// @Tags Authentication
// @Param ResendConfirmCode body services.ResendConfirmCode true "Resend confirm code"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /auth/resend-confirm-code [post]
func (h *AuthHandler) PostResendConfirmCode(c *gin.Context) {

	var user *services.ResendConfirmCode
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	result, ok := h.authService.ResendConfirmCode(user.Username)

	if ok != nil {
		h.resp.BadRequest(c, ok.Error())
		return
	}
	resultMsg := fmt.Sprintf("ใช้ Confirmation code จากอีเมล %v เพื่อยืนยันตัวตน", *result.CodeDeliveryDetails.Destination)
	h.resp.Success(c, resultMsg)

}

// Authenticate godoc
// @Summary Refresh user token
// @Description refresh new user token
// @Produce json
// @Tags Authentication
// @Security BearerAuth
// @Param refresh_token body string true "refresh a token string"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /auth/refresh-token [post]
func (h *AuthHandler) PostRefreshToken(c *gin.Context) {

	refreshToken := &services.SignInResponse{}
	custErr := utils.NewErrorHandler(c)
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
			h.resp.BadRequest(c, reshErr[1])
			return
		}

		h.resp.BadRequest(c, ok.Error())
		return
	}

	h.resp.Success(c, resToken)
	return
}

// Authenticate godoc
// @Summary Refresh user token
// @Description refresh new user token
// @Produce json
// @Tags Authentication
// @Param resendConfirmCode body services.ResendConfirmCode true "response confirm code"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /auth/forgot-password [post]
func (h *AuthHandler) PostForgotPw(c *gin.Context) {

	userName := &services.ResendConfirmCode{}

	err := c.ShouldBindJSON(userName)
	cusErr := utils.NewErrorHandler(c)
	if err != nil {

		erList := cusErr.MyErr(err)

		for _, fe := range erList {
			textErr := fmt.Sprintf("field is a %s %s", fe.Field, fe.Msg)
			h.resp.BadRequest(c, textErr)

		}

		return
	}

	response, err := h.authService.ForgotPassword(userName.Username)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	sendToEmail := *response.CodeDeliveryDetails.Destination
	txtResponse := fmt.Sprintf("ระบบได้ส่ง Password Reset Code ไปที่ %v : %v กรุณาใช้ Code เพื่อยืนยัน", response.CodeDeliveryDetails.DeliveryMedium, sendToEmail)
	// Response Success
	h.resp.Success(c, txtResponse)

}

// Authenticate godoc
// @Summary Refresh user token
// @Description refresh new user token
// @Produce json
// @Tags Authentication
// @Param ConfirmNewPassword body services.UserConfirmNewPassword true "response confirm new password"
// @Success 200 {object} utils.ApiResponse{}
// @Failure 400 {object} utils.ApiResponse{}
// @Router /auth/confirm-password [post]
func (h *AuthHandler) PostConfirmNewPassword(c *gin.Context) {
	confirmReq := services.UserConfirmNewPassword{}
	customErr := utils.NewErrorHandler(c)
	err := c.ShouldBindJSON(&confirmReq)

	if err != nil {
		customErr.CustomError(err)
		return
	}

	response, err := h.authService.ConfirmNewPassword(&confirmReq)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	// Response Success
	h.resp.Success(c, response)
}

func (h *AuthHandler) PostChangePassword(c *gin.Context) {

	changePasswordReq := services.ChangePasswordReq{}
	err := c.ShouldBindJSON(&changePasswordReq)
	custErr := utils.NewErrorHandler(c)
	if err != nil {
		custErr.CustomError(err)
		return
	}

	resChange, err := h.authService.ChangePassword(&changePasswordReq)
	if err != nil {
		h.resp.BadRequest(c, err.Error())
		return
	}

	// Success Response
	h.resp.Success(c, resChange)

}

func (h *AuthHandler) DelCustomer(c *gin.Context) {
	acsessToken := services.UserDelete{}
	err := c.ShouldBindJSON(&acsessToken)
	utils.NewErrorHandler(c)

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
