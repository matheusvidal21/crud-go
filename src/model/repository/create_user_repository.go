package repository

import (
	"context"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

var (
	journey_create_user_repository = "create_user_repository"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("journey", journey_create_user_repository))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.database.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to insert user", err, zap.String("journey", journey_create_user_repository))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("CreateUser repository executed with success",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", journey_create_user_repository))
	return converter.ConvertEntityToDomain(*value), nil
}
