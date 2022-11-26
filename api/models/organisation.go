package models

type Organisation struct {
	Name        string `json:"name" bson:"name"`
	Revenue     string `json:"revenue" bson:"revenue"`
	Phone       string `json:"phone" bson:"phone"`
	Website     string `json:"website" bson:"website"`
	Location    string `json:"location" bson:"location"`
	Designation string `json:"designation" bson:"designation"`
	Industry    string `json:"industry" bson:"industry"`
}
