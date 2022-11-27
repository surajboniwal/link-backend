package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/repositories"
)

type AuthController struct {
	authRepository repositories.AuthRepository
}

func NewAuthController(repo repositories.AuthRepository) AuthController {
	return AuthController{
		authRepository: repo,
	}
}

func (authController AuthController) Register(context *gin.Context) {
	context.JSON(
		http.StatusCreated,
		gin.H{
			"status": true,
		},
	)
}

func (authController AuthController) Login(context *gin.Context) {
	context.JSON(
		http.StatusCreated,
		gin.H{
			"status": true,
		},
	)
}
