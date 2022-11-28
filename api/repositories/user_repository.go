package repositories

import (
	"context"

	"github.com/surajboniwal/link-backend/api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	CreateUser(*models.User) (interface{}, error)
}

type UserRepositoryImpl struct {
	db *mongo.Database
}

func NewUserRepositoryImpl(db *mongo.Database) UserRepositoryImpl {
	return UserRepositoryImpl{
		db: db,
	}
}

func (userRepository UserRepositoryImpl) CreateUser(user *models.User) (interface{}, error) {
	result, err := userRepository.db.Collection("users").InsertOne(context.Background(), user, options.InsertOne())

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}
