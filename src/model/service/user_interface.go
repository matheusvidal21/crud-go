package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/model/repository"
)

type userDomainService struct {
	repository repository.UserRepository
}

func NewUserDomainService(repository repository.UserRepository) UserDomainService {
	return &userDomainService{
		repository: repository,
	}
}

type UserDomainService interface {
	CreateUserServices(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserServices(string) *rest_err.RestErr
}
