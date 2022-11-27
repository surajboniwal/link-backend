package repositories

import (
	"context"
	"log"

	"github.com/surajboniwal/link-backend/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConstantsRepository interface {
	GetConstants() Constants
}

type ConstantsRepositoryImpl struct {
	db *mongo.Database
}

func NewConstantsRepositoryImpl(db *mongo.Database) ConstantsRepositoryImpl {
	return ConstantsRepositoryImpl{
		db: db,
	}
}

func (constantsRepository ConstantsRepositoryImpl) GetConstants() Constants {
	var constants Constants

	industryCursor, err := constantsRepository.db.Collection("industry_options").Find(context.Background(), bson.M{}, options.Find())
	revenueCursor, err2 := constantsRepository.db.Collection("revenue_options").Find(context.Background(), bson.M{}, options.Find())

	if err != nil || err2 != nil {
		log.Fatal("Unable to fetch industries")
	}

	industryCursor.All(context.TODO(), &constants.IndustryOptions)
	revenueCursor.All(context.TODO(), &constants.RevenueOptions)

	return constants
}

type Constants struct {
	IndustryOptions []models.IndustryOptions `json:"industry_options"`
	RevenueOptions  []models.RevenueOptions  `json:"revenue_options"`
}
