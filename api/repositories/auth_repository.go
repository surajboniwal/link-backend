package repositories

import "go.mongodb.org/mongo-driver/mongo"

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

type RegisterParams struct {
	FirstName        string `json:"first_name" bson:"first_name"`
	LastName         string `json:"last_name" bson:"last_name"`
	Email            string `json:"email" bson:"email"`
	OrganisationName string `json:"organisation_name" bson:"organisation_name"`
}
