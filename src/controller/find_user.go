package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

var (
	journey_find_user_by_id    = "find_user_by_id_controller"
	journey_find_user_by_email = "find_user_by_email_controller"
)

func (uc *UserController) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID controller", zap.String("journey", journey_find_user_by_id))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info(fmt.Sprintf("User autheticaded: %#v", user))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", journey_find_user_by_id))
		errorMessage := rest_err.NewBadRequestError("UserID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.userService.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call FindUserByIDServices", err, zap.String("journey", journey_find_user_by_id))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUserByID controller executed with success",
		zap.String("userId", userId),
		zap.String("journey", journey_find_user_by_id),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *UserController) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller", zap.String("journey", journey_find_user_by_email))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info(fmt.Sprintf("User autheticaded: %#v", user))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", journey_find_user_by_email))
		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.userService.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call FindUserByEmailServices", err, zap.String("journey", journey_find_user_by_email))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUserByEmail controller executed with success",
		zap.String("userEmail", userEmail),
		zap.String("journey", journey_find_user_by_email),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
