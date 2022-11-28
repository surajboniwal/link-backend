package repositories

import (
	"context"

	"github.com/surajboniwal/link-backend/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrganisationRepository interface {
	CreateOrganisation(*models.Organisation) (interface{}, error)
	GetOrganisation() ([]models.Organisation, error)
}

type OrganisationRepositoryImpl struct {
	db *mongo.Database
}

func NewOrganisationRepositoryImpl(db *mongo.Database) OrganisationRepositoryImpl {
	return OrganisationRepositoryImpl{
		db: db,
	}
}

func (repo OrganisationRepositoryImpl) CreateOrganisation(data *models.Organisation) (interface{}, error) {
	result, err := repo.db.Collection("organisations").InsertOne(context.Background(), data)

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (repo OrganisationRepositoryImpl) GetOrganisation() ([]models.Organisation, error) {
	var data []models.Organisation

	findOptions := options.Find()

	cursor, err := repo.db.Collection("organisations").Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	if err := cursor.All(context.Background(), &data); err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	return data, nil
}
