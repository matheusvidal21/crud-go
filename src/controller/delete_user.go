package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

var (
	journey_delete_user_controller = "delete_user_controller"
)

func (uc *UserController) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("journey", journey_delete_user_controller))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate user id", err, zap.String("journey", journey_delete_user_controller))
		errorMessage := rest_err.NewBadRequestError("Invalid userId, must be a valid hex")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := uc.userService.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error trying to call DeleteUser service", err, zap.String("journey", journey_delete_user_controller))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully",
		zap.String("user_id", userId),
		zap.String("journey", journey_delete_user_controller))
	c.Status(http.StatusOK)
}
