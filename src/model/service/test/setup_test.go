package service_test

import (
	"github.com/matheusvidal21/crud-go/src/model/service"
	"github.com/matheusvidal21/crud-go/src/test/mocks"
	"go.uber.org/mock/gomock"
	"testing"
)

func SetupTest(t *testing.T) (service.UserDomainService, *mocks.MockUserRepository) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	repository := mocks.NewMockUserRepository(ctrl)
	return service.NewUserDomainService(repository), repository
}
