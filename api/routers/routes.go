package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

type Controllers struct {
	OrganisationController controllers.OrganisationController
}

func SetupRouter(router *gin.Engine, controllers Controllers) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": true,
		})
	})

	SetupOrganisationRouter(router, controllers.OrganisationController)

}
