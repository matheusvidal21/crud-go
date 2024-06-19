package repository_test

import (
	"fmt"
	"github.com/matheusvidal21/crud-go/src/model/repository"
	"github.com/matheusvidal21/crud-go/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	mtestDb := setupTest(t)

	mtestDb.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@teste.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail(userEntity.Email)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail("test")
		assert.Equal(t, err.Message, fmt.Sprintf("User not found with email: test"))
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserById(t *testing.T) {
	mtestDb := setupTest(t)

	mtestDb.Run("when_sending_a_valid_id_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@teste.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByID("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByID("test")
		assert.EqualValues(t, err.Message, fmt.Sprintf("User not found with id: test"))
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	mtestDb := setupTest(t)

	mtestDb.Run("when_sending_a_valid_email_and_password_returns_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@teste.com",
			Password: "test",
			Name:     "test",
			Age:      50,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, userEntity.Password)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mtestDb.Run("returns_error_when_mongodb_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword("test", "testpass")
		assert.EqualValues(t, err.Message, "Error trying to find user by email and password")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("returns_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(database_name)
		repo := repository.NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword("test", "testpass")
		assert.EqualValues(t, err.Message, fmt.Sprintf("User or password is invalid"))
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
