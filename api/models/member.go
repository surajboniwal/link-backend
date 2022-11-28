package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	OrganisationID primitive.ObjectID `bson:"organisation_id" json:"organisation_id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"user_id" binding:"required"`
}
