package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
	"github.com/surajboniwal/link-backend/api/repositories"
	"github.com/surajboniwal/link-backend/api/routers"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	db     *mongo.Database
	router *gin.Engine
}

func NewServer(db *mongo.Database) *Server {

	server := &Server{db: db}
	router := gin.New()

	controllers := routers.NewControllers(controllers.NewOrganisationController(repositories.NewOrganisationRepositoryImpl(db)))

	routers.SetupRouter(router, controllers)

	server.router = router

	return server
}

func (server *Server) StartServer(port int) {
	stringPort := strconv.Itoa(port)
	server.router.Run(":" + stringPort)
}
