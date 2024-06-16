package repository

import (
	"context"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

var (
	journey_update_user_repository = "update_user_repository"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository", zap.String("journey", journey_update_user_repository))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.database.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	objectId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: value}}
	if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", journey_update_user_repository))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateUser repository executed with success",
		zap.String("userId", userId),
		zap.String("journey", journey_update_user_repository))
	return nil
}
