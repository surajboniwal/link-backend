package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/controllers"
)

func SetupAuthRouter(router *gin.Engine, authController controllers.AuthController) {
	r := router.Group("/auth")
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
}
