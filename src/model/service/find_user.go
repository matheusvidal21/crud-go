package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	journey_find_user_by_id_services    = "find_user_by_id_services"
	journey_find_user_by_email_services = "find_user_by_email_services"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services", zap.String("journey", journey_find_user_by_id_services))
	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", journey_find_user_by_email_services))
	return ud.repository.FindUserByEmail(email)
}
