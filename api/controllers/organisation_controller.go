package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
	"github.com/surajboniwal/link-backend/api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrganisationController struct {
	organisationRepository repositories.OrganisationRepository
}

func NewOrganisationController(repo repositories.OrganisationRepository) OrganisationController {
	return OrganisationController{
		organisationRepository: repo,
	}
}

func (organisationController OrganisationController) GetOrganisation(ctx *gin.Context) {

	data, err := organisationController.organisationRepository.GetOrganisation()

	if err != nil {
		helpers.ResponseDispatch(ctx, nil, err, http.StatusInternalServerError)
		return
	}

	helpers.ResponseDispatch(ctx, data, nil, http.StatusOK)
}

func (organisationController OrganisationController) GetSingleOrganisation(ctx *gin.Context) {

	id, paramErr := primitive.ObjectIDFromHex(ctx.Param("id"))

	if paramErr != nil {
		helpers.ResponseDispatch(ctx, nil, "Invalid ID", http.StatusBadRequest)
		return
	}

	data, err := organisationController.organisationRepository.GetSingleOrganisation(id)

	if err != nil {
		helpers.ResponseDispatch(ctx, nil, err, http.StatusBadRequest)
		return
	}

	helpers.ResponseDispatch(ctx, data, nil, http.StatusOK)
}
