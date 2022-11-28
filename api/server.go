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

	organisationController := controllers.NewOrganisationController(repositories.NewOrganisationRepositoryImpl(db))
	authController := controllers.NewAuthController(repositories.NewAuthRepositoryImpl(db), organisationController)
	constantsController := controllers.NewConstantsController(repositories.NewConstantsRepositoryImpl(db))

	controllers := routers.Controllers{
		OrganisationController: organisationController,
		AuthController:         authController,
		ConstantsController:    constantsController,
	}

	routers.SetupRouter(router, controllers)

	server.router = router

	return server
}

func (server *Server) StartServer(port int) {
	stringPort := strconv.Itoa(port)
	server.router.Run(":" + stringPort)
}
