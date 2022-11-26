package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/surajboniwal/link-backend/api"
	"github.com/surajboniwal/link-backend/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Database *mongo.Database
	Router   *gin.Engine
}

func main() {

	_ = godotenv.Load()

	db := database.Connect()

	server := api.NewServer(db)

	server.StartServer(8000)

}
