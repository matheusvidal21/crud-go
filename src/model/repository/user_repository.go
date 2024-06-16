package repository

import (
	"github.com/matheusvidal21/crud-go/src/configuration/rest_err"
	"github.com/matheusvidal21/crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(userId string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
	FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
}

type userRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database: database,
	}
}
