package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
	"github.com/surajboniwal/link-backend/api/repositories"
)

type UserController struct {
	userRepository repositories.UserRepository
}

func NewUserController(userRepository repositories.UserRepository) UserController {
	return UserController{
		userRepository: userRepository,
	}
}

func (userController UserController) CreateUser(context *gin.Context) {
	helpers.ResponseDispatch(context, "User created", nil, http.StatusCreated)
}
