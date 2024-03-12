package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockUserRepository serves as a mock for UserRepository
type MockUserRepository struct {
	repository.UserRepository
}

func TestUserService(t *testing.T) {
	tests := []struct {
		name     string
		newUser  *repository.User
		want     int
		wantErr  bool
		errValue string
	}{
		{
			name:     "create user valid",
			newUser:  &repository.User{Name: "John Doe"},
			want:     http.StatusCreated,
			wantErr:  false,
			errValue: "",
		},
		{
			name:     "create user error",
			newUser:  &repository.User{Name: ""},
			want:     http.StatusBadRequest,
			wantErr:  true,
			errValue: "name already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockUserRepository)
			mockService := NewUserService(mockRepo)

			router := gin.Default()
			router.POST("/users", mockService.CreateUser)

			payload := `{"Name":"` + tt.newUser.Name + `"}`

			req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(payload))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.want, resp.Code)
		})
	}
}