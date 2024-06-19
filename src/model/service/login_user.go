package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	journey_login_user_service = "login_user_service"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", journey_login_user_service))

	userDomain.EncryptPassword()
	user, err := ud.FindUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		logger.Error("Error trying to find user by email and password", err, zap.String("journey", journey_login_user_service))
		return nil, "", err
	}

	if user == nil {
		logger.Error("User not found", err, zap.String("journey", journey_login_user_service))
		return nil, "", rest_err.NewNotFoundError("User not found")
	}

	token, err := user.GenerateToken()
	if err != nil {
		logger.Error("Error trying to generate token", err, zap.String("journey", journey_login_user_service))
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully",
		zap.String("user_id", user.GetID()),
		zap.String("journey", journey_login_user_service))
	return user, token, nil
}
