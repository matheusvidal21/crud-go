package model

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "create_user_domain"))
	ud.EncryptPassword()
	return nil
}
