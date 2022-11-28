package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

func SetupUserRouter(router *gin.Engine, userController controllers.UserController) {
	r := router.Group("/user")

	r.POST("/", userController.CreateUser)
}
