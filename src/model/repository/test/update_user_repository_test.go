package repository_test

import (
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/model/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	mtestDb := setupTest(t)

	mtestDb.Run("when_sending_a_valid_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)
		domain := model.NewUserDomain("teste@teste.com", "@#12345", "Teste", 20)

		domain.SetID(primitive.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetID(), domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)
		domain := model.NewUserDomain("teste@teste.com", "@#12345", "Teste", 20)

		domain.SetID(primitive.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetID(), domain)

		assert.NotNil(t, err)
	})
}
