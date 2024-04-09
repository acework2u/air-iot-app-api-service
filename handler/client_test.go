package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/middleware"
	service "github.com/acework2u/air-iot-app-api-service/services/clientcoginto"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var BaseUrl = "http://localhost:8080/api/v1"

func TestClientHandler_PostSignUp(t *testing.T) {

	client := http.Client{}
	apiUrl := fmt.Sprintf("%v/auth/signup", BaseUrl)
	resp, err := client.Post(apiUrl, "Application/json", nil)
	if err != nil {

	}
	_ = resp

}
func Router() *gin.Engine {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/signup")
	publicRoutes.POST("/signin")

	protectedRoutes := router.Group("/v1/api")
	protectedRoutes.Use(middleware.CognitoAuthMiddleware())

	return router

}

func makeRequest(method, url string, body interface{}, isAuthenticatedRequest bool) *httptest.ResponseRecorder {
	reqBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if isAuthenticatedRequest {
		request.Header.Add("Authorization", "Bearer "+bearerToken())
	}
	writer := httptest.NewRecorder()
	Router().ServeHTTP(writer, request)
	return writer
}

func bearerToken() string {
	user := service.SignUpRequest{}
	apiUrl := fmt.Sprintf("%v/auth/signin", BaseUrl)
	writer := makeRequest("POST", apiUrl, user, false)
	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)
	//json.Unmarshal(writer.Body.Bytes(), &response)
	return response["jwt"]
}
