package repository

import (
	"context"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"os"
)

var (
	journey_delete_user_repository = "delete_user_repository"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUser repository", zap.String("journey", journey_delete_user_repository))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.database.Collection(collection_name)

	objectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: objectId}}
	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		logger.Error("Error trying to delete user", err, zap.String("journey", journey_delete_user_repository))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("DeleteUser repository executed with success",
		zap.String("userId", userId),
		zap.String("journey", journey_delete_user_repository))
	return nil
}
