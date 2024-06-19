package repository_test

import (
	"github.com/matheusvidal21/crud-go/src/model/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	mtestDb := setupTest(t)

	mtestDb.Run("when_sending_a_valid_userId_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)
		err := repo.DeleteUser("test")

		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)
		err := repo.DeleteUser("test")

		assert.NotNil(t, err)
	})
}
