package service

import (
	"fmt"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "create_user_domain"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
