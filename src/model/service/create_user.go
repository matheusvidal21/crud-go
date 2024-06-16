package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	journey_create_user_services = "create_user_services"
)

func (ud *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser services", zap.String("journey", journey_create_user_services))

	if user, _ := ud.FindUserByEmailServices(userDomain.GetEmail()); user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	user, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", journey_create_user_services))
		return nil, err
	}
	logger.Info("CreateUser service executed successfully",
		zap.String("user_id", user.GetID()),
		zap.String("journey", journey_create_user_services))
	return user, nil
}
