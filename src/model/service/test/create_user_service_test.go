package service_test

import (
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	service, repository := SetupTest(t)

	t.Run("when_user_already_exists_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "Email is already registered in another account")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("error trying to create user"))

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomain.GetID(), user.GetID())
		assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
		assert.EqualValues(t, userDomain.GetName(), user.GetName())
		assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())
	})

}
