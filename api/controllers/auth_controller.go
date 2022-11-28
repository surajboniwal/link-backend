package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
	"github.com/surajboniwal/link-backend/api/models"
	"github.com/surajboniwal/link-backend/api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	authRepository         repositories.AuthRepository
	organisationRepository repositories.OrganisationRepository
	userRepository         repositories.UserRepository
}

func NewAuthController(repo repositories.AuthRepository, organisationRepository repositories.OrganisationRepository, userRepository repositories.UserRepository) AuthController {
	return AuthController{
		authRepository:         repo,
		organisationRepository: organisationRepository,
		userRepository:         userRepository,
	}
}

func (authController AuthController) Register(context *gin.Context) {

	var params RegisterParams

	if err := context.BindJSON(&params); err != nil {
		helpers.ResponseDispatch(context, nil, err.Error(), http.StatusBadRequest)
		return
	}

	organisation := models.Organisation{
		ID:          primitive.NewObjectID(),
		Name:        params.OrganisationName,
		Revenue:     params.Revenue,
		Phone:       params.Phone,
		Website:     params.Website,
		Location:    params.Location,
		Designation: params.Designation,
		Industry:    params.Industry,
	}

	//Create an organisation
	_, err := authController.organisationRepository.CreateOrganisation(&organisation)

	if err != nil {
		helpers.ResponseDispatch(context, nil, err, http.StatusInternalServerError)
		return
	}

	//Create an user

	// _, userError := authController.userRepository.CreateUser()

	//Create an member

	helpers.ResponseDispatch(context, params, nil, http.StatusOK)
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
	Phone            models.Phone       `json:"phone" binding:"required"`
	Website          string             `json:"website" binding:"required"`
	Location         string             `json:"location" binding:"required"`
	Designation      string             `json:"designation" binding:"required"`
	Industry         primitive.ObjectID `json:"industry_option" binding:"required"`
}
