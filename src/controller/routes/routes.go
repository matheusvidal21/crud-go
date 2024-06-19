package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/controller"
	"github.com/matheusvidal21/crud-go/src/model"
)

func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", model.MiddlewareVerifyToken, controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.MiddlewareVerifyToken, controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", model.MiddlewareVerifyToken, controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.MiddlewareVerifyToken, controller.DeleteUser)
	r.POST("/login", controller.LoginUser)
}
