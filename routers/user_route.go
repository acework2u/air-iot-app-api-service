package routers

import (
	users "github.com/acework2u/air-iot-app-api-service/services/user"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController users.UserService
}

func NewUserRouteController(userController users.UserService) UserRouteController {
	return UserRouteController{userController: userController}
}

func (r *UserRouteController) UserRoute(rg *gin.RouterGroup, userService users.UserService) {
	router := rg.Group("/users")

	//router.POST("/", r.userController.CreateUser)
	router.GET("/", r.userController.FindAll)
}
