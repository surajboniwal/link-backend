package repositories

import (
	"context"
	"fmt"

	"github.com/surajboniwal/link-backend/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemberRepository interface {
	CreateMember(member *models.Member) (interface{}, error)
	GetMembers(id primitive.ObjectID) ([]models.Member, error)
}

type MemberRepositoryImpl struct {
	db *mongo.Database
}

func NewMemberRepositoryImpl(db *mongo.Database) MemberRepositoryImpl {
	return MemberRepositoryImpl{
		db: db,
	}
}

func (repo MemberRepositoryImpl) CreateMember(member *models.Member) (interface{}, error) {
	result, err := repo.db.Collection("members").InsertOne(context.Background(), member, options.InsertOne())

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (repo MemberRepositoryImpl) GetMembers(id primitive.ObjectID) ([]models.Member, error) {
	var data []models.Member

	fmt.Println(id)

	cursor, err := repo.db.Collection("members").Find(context.Background(), bson.D{{"organisation_id", id}}, options.Find())

	if err != nil {
		return data, err
	}

	cursor.All(context.Background(), &data)

	return data, nil
}
