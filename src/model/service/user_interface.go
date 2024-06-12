package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
)

type userDomainService struct{}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type UserDomainService interface {
	CreateUser(domainInterface model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
