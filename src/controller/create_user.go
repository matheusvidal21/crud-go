package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/validation"
	"github.com/matheusvidal21/crud-go/src/controller/model/request"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "create_user"))
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "create_user"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "create_user"))
	c.String(http.StatusCreated, "")
}
