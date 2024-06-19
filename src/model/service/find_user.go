package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.uber.org/zap"
)

var (
	journey_find_user_by_id_services                 = "find_user_by_id_services"
	journey_find_user_by_email_services              = "find_user_by_email_services"
	journey_find_user_by_email_and_password_services = "find_user_by_email_and_password_services"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID services", zap.String("journey", journey_find_user_by_id_services))
	user, err := ud.repository.FindUserByID(id)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", journey_find_user_by_id_services))
		return nil, err
	}
	logger.Info("FindUserByID services executed with success",
		zap.String("user_id", id),
		zap.String("journey", journey_find_user_by_id_services))
	return user, nil
}

func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services", zap.String("journey", journey_find_user_by_email_services))
	user, err := ud.repository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", journey_find_user_by_email_services))
		return nil, err
	}
	logger.Info("FindUserByEmail services executed with success",
		zap.String("email", email),
		zap.String("journey", journey_find_user_by_email_services))
	return user, nil
}

func (ud *userDomainService) FindUserByEmailAndPasswordServices(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword services", zap.String("journey", journey_find_user_by_email_and_password_services))
	user, err := ud.repository.FindUserByEmailAndPassword(email, password)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", journey_find_user_by_email_and_password_services))
		return nil, err
	}
	logger.Info("FindUserByEmailAndPassword services executed with success",
		zap.String("email", email),
		zap.String("journey", journey_find_user_by_email_and_password_services))
	return user, nil
}
