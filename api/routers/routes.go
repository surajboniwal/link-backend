package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

type Controllers struct {
	organisationController controllers.OrganisationController
	authController         controllers.AuthController
}

func NewControllers(organisationController controllers.OrganisationController, authController controllers.AuthController) Controllers {
	return Controllers{
		organisationController: organisationController,
		authController:         authController,
	}
}

func SetupRouter(router *gin.Engine, controllers Controllers) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	})

	SetupOrganisationRouter(router, controllers.organisationController)
	SetupAuthRouter(router, controllers.authController)

}
