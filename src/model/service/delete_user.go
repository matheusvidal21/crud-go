package service

import (
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
)

var (
	journey_delete_user_services = "delete_user_services"
)

func (*userDomainService) DeleteUserServices(userId string) *rest_err.RestErr {
	return nil
}
