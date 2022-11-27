package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

func SetupConstantsRouter(router *gin.Engine, constantsController controllers.ConstantsController) {
	r := router.Group("/constants")
	r.GET("/", constantsController.GetConstants)
}
