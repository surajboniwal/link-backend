package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	Register()
	Login()
}

type AuthRepositoryImpl struct {
	db *mongo.Database
}

func NewAuthRepositoryImpl(db *mongo.Database) AuthRepositoryImpl {
	return AuthRepositoryImpl{
		db: db,
	}
}

func (r AuthRepositoryImpl) Register() {

}

func (r AuthRepositoryImpl) Login() {

}
