package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organisation struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Revenue     primitive.ObjectID `bson:"revenue" json:"revenue"`
	Phone       Phone              `bson:"phone" json:"phone"`
	Website     string             `bson:"website" json:"website"`
	Location    string             `bson:"location" json:"location"`
	Designation string             `bson:"designation" json:"designation"`
	Industry    primitive.ObjectID `bson:"industry" json:"industry"`
}
