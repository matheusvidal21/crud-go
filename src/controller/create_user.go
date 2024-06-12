package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/validation"
	"github.com/matheusvidal21/crud-go/src/controller/model/request"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/view"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *UserController) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "create_user"))
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "create_user"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := uc.userService.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "create_user"))
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}
