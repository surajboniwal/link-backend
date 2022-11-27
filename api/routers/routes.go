package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

type Controllers struct {
	OrganisationController controllers.OrganisationController
	AuthController         controllers.AuthController
	ConstantsController    controllers.ConstantsController
}

func SetupRouter(router *gin.Engine, controllers Controllers) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	})

	SetupOrganisationRouter(router, controllers.OrganisationController)
	SetupAuthRouter(router, controllers.AuthController)
	SetupConstantsRouter(router, controllers.ConstantsController)

}
