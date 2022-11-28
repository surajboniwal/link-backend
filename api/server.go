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
	router := gin.Default()

	controllers := setupDI(db)

	routers.SetupRouter(router, controllers)

	server.router = router

	return server
}

func setupDI(db *mongo.Database) routers.Controllers {

	constantsRepository := repositories.NewConstantsRepositoryImpl(db)
	organisationRepository := repositories.NewOrganisationRepositoryImpl(db)
	userRepository := repositories.NewUserRepositoryImpl(db)
	authRepository := repositories.NewAuthRepositoryImpl(db)

	constantsController := controllers.NewConstantsController(constantsRepository)
	organisationController := controllers.NewOrganisationController(organisationRepository)
	userController := controllers.NewUserController(userRepository)
	authController := controllers.NewAuthController(authRepository, organisationRepository, userRepository)

	return routers.Controllers{
		OrganisationController: organisationController,
		AuthController:         authController,
		ConstantsController:    constantsController,
		UserController:         userController,
	}

}

func (server *Server) StartServer(port int) {
	stringPort := strconv.Itoa(port)
	server.router.Run(":" + stringPort)
}
