package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/model/service"
)

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserController struct {
	userService service.UserDomainService
}

func NewUserController(userService service.UserDomainService) UserControllerInterface {
	return &UserController{
		userService: userService,
	}
}
