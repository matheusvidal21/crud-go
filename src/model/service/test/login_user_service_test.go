package service_test

import (
	"github.com/golang-jwt/jwt"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

func TestUserDomainService_LoginUserServices(t *testing.T) {
	service, repository := SetupTest(t)

	t.Run("when_calling_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		userDomainMock := model.NewUserDomain(userDomain.GetEmail(), userDomain.GetPassword(), userDomain.GetName(), userDomain.GetAge())
		userDomainMock.EncryptPassword()

		repository.EXPECT().FindUserByEmailAndPassword(userDomain.GetEmail(), userDomainMock.GetPassword()).
			Return(nil, rest_err.NewInternalServerError("error trying to create user"))

		user, token, err := service.LoginUserServices(userDomain)

		assert.NotNil(t, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
		assert.EqualValues(t, "error trying to create user", err.Message)
	})

	t.Run("when_calling_create_token_returns_error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(gomock.NewController(t))
		userDomainMock.EXPECT().GetEmail().Return("test@test.com")
		userDomainMock.EXPECT().GetPassword().Return("testpass")
		userDomainMock.EXPECT().EncryptPassword()
		userDomainMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("error trying to create token"))

		repository.EXPECT().FindUserByEmailAndPassword("test@test.com", "testpass").
			Return(userDomainMock, nil)

		user, token, err := service.LoginUserServices(userDomainMock)

		assert.NotNil(t, err)
		assert.Empty(t, token)
		assert.Nil(t, user)
		assert.EqualValues(t, "error trying to create token", err.Message)
	})

	t.Run("when_user_and_password_is_valid_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		secret := "test"
		os.Setenv("JWT_SECRET_KEY", secret)
		defer os.Clearenv()

		userDomain := model.NewUserDomain("teste@teste.com", "testpass", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(userDomain.GetEmail(), gomock.Any()).
			Return(userDomain, nil)

		userDomainReturn, token, err := service.LoginUserServices(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), userDomain.GetID())
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		tokenReturned, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}
			return nil, rest_err.NewBadRequestError("invalid token")
		})

		_, ok := tokenReturned.Claims.(jwt.MapClaims)
		if !ok || !tokenReturned.Valid {
			t.FailNow()
			return
		}

	})

}
