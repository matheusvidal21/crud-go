package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	journey_update_user_services = "update_user_services"
)

func (ud *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser services", zap.String("journey", journey_update_user_services))

	err := ud.repository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", journey_update_user_services))
		return err
	}
	logger.Info("UpdateUser service executed successfully",
		zap.String("user_id", userId),
		zap.String("journey", journey_update_user_services))
	return nil
}
