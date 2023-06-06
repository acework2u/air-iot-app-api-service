package user

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/acework2u/air-iot-app-api-service/repository"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userService repository.UserRepository
}

func NewUserService(userService repository.UserRepository) UserService {
	return UserService{userService: userService}
}

func (s *UserService) CreateUser(ctx *gin.Context) {
	var user *repository.User

	// _ = user

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())

		return
	}

	fmt.Println(user)

	// newUser, err := s.userService.FindPosts(user)
	newUser, err := s.userService.CreateUser(user)

	if err != nil {
		if strings.Contains(err.Error(), "name already exits") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})

}

func (s *UserService) FindAll(ctx *gin.Context) {

	var user *repository.User

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": user,
		"data":   "Find User",
	})

}
