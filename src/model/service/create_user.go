package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "create_user_domain"))
	userDomain.EncryptPassword()

	user, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "create_user_domain"))
	}
	logger.Info("CreateUser service executed successfully",
		zap.String("user_id", user.GetID()),
		zap.String("journey", "create_user_domain"))
	return user, nil
}
