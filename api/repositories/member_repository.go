package repositories

import (
	"context"

	"github.com/surajboniwal/link-backend/api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemberRepository interface {
	CreateMember(member *models.Member) (interface{}, error)
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
