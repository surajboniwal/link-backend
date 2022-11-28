package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

func SetupOrganisationRouter(router *gin.Engine, organisationController controllers.OrganisationController) {
	r := router.Group("/organisation")

	r.GET("/", organisationController.GetOrganisation)
}
