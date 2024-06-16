package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/configuration/validation"
	"github.com/matheusvidal21/crud-go/src/controller/model/request"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

var (
	journey_update_user_controller = "update_user_controller"
)

func (uc *UserController) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", journey_update_user_controller))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate user id", err, zap.String("journey", journey_update_user_controller))
		errorMessage := rest_err.NewBadRequestError("Invalid userId, must be a valid hex")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var userRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", journey_update_user_controller))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)
	err := uc.userService.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error trying to call UpdateUser service", err, zap.String("journey", journey_update_user_controller))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully",
		zap.String("user_id", userId),
		zap.String("journey", journey_update_user_controller))
	c.Status(http.StatusOK)
}
