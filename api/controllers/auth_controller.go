package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
	"github.com/surajboniwal/link-backend/api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	var params RegisterParams

	if err := context.BindJSON(&params); err != nil {
		helpers.ResponseDispatch(context, nil, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseDispatch(context, params, nil, http.StatusBadRequest)
}

func (authController AuthController) Login(context *gin.Context) {
	context.JSON(
		http.StatusCreated,
		gin.H{
			"status": true,
		},
	)
}

type RegisterParams struct {
	FirstName        string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName         string             `json:"last_name" bson:"last_name" binding:"required"`
	Email            string             `json:"email" bson:"email" binding:"required"`
	OrganisationName string             `json:"organisation_name" bson:"organisation_name" binding:"required"`
	Revenue          primitive.ObjectID `json:"revenue_option" binding:"required"`
	Phone            RegisterParamPhone `json:"phone" binding:"required"`
	Website          string             `json:"website" binding:"required"`
	Location         string             `json:"location" binding:"required"`
	Designation      string             `json:"designation" binding:"required"`
	Industry         primitive.ObjectID `json:"industry_option" binding:"required"`
}

type RegisterParamPhone struct {
	CountryCode string `json:"country_code" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
