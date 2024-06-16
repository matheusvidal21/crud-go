package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

var (
	journey_delete_user_services = "delete_user_services"
)

func (ud *userDomainService) DeleteUserServices(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser services", zap.String("journey", journey_delete_user_services))
	err := ud.repository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", journey_delete_user_services))
		return err
	}
	logger.Info("DeleteUser service executed successfully",
		zap.String("user_id", userId),
		zap.String("journey", journey_delete_user_services))
	return nil
}
