package repository

import (
	"context"
	"fmt"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"github.com/matheusvidal21/crud-go/src/model/repository/entity"
	"github.com/matheusvidal21/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"os"
)

var (
	journey_find_user_by_email_repository              = "find_user_by_email_repository"
	journey_find_user_by_id_repository                 = "find_user_by_id_repository"
	journey_find_user_by_email_and_password_repository = "find_user_by_email_and_password_repository"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository", zap.String("journey", journey_find_user_by_email_repository))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.database.Collection(collection_name)

	entity := &entity.UserEntity{}

	filter := bson.D{{
		Key:   "email",
		Value: email,
	}}
	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", journey_find_user_by_email_repository))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", journey_find_user_by_email_repository))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", journey_find_user_by_email_repository),
		zap.String("email", email),
		zap.String("user_id", entity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*entity), nil
}

func (ur *userRepository) FindUserByID(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID repository", zap.String("journey", journey_find_user_by_id_repository))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.database.Collection(collection_name)

	entity := &entity.UserEntity{}

	objectID, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{"_id", objectID}}
	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with id: %s", userId)
			logger.Error(errorMessage, err, zap.String("journey", journey_find_user_by_id_repository))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", journey_find_user_by_id_repository))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", journey_find_user_by_id_repository),
		zap.String("user_id", entity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*entity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailAndPassword repository", zap.String("journey", journey_find_user_by_email_and_password_repository))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.database.Collection(collection_name)

	entity := &entity.UserEntity{}

	filter := bson.D{{"email", email}, {"password", password}}
	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password is invalid"
			logger.Error(errorMessage, err, zap.String("journey", journey_find_user_by_email_and_password_repository))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, err, zap.String("journey", journey_find_user_by_email_and_password_repository))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully",
		zap.String("journey", journey_find_user_by_email_and_password_repository),
		zap.String("email", email),
		zap.String("userId", entity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*entity), nil
}
