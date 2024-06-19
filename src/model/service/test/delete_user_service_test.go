package service_test

import (
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUserDomainService_DeleteUserServices(t *testing.T) {
	service, repository := SetupTest(t)

	t.Run("when_sending_a_valid_userId_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUserServices(id)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_valid_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("error trying to delete user"))

		err := service.DeleteUserServices(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to delete user")
	})

}
