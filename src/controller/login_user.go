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
	journey_login_user_controller = "login_user_controller"
)

func (uc *UserController) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", journey_login_user_controller))

	var userRequest request.UserLoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", journey_login_user_controller))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
	domainResult, err := uc.userService.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error trying to call LoginUser service", err, zap.String("journey", journey_login_user_controller))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("LoginUser controller executed successfully",
		zap.String("user_id", domainResult.GetID()),
		zap.String("journey", journey_login_user_controller))
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
