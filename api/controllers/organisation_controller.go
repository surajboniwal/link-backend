package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
	"github.com/surajboniwal/link-backend/api/models"
	"github.com/surajboniwal/link-backend/api/repositories"
)

type OrganisationController struct {
	organisationRepository repositories.OrganisationRepository
}

func NewOrganisationController(repo repositories.OrganisationRepository) OrganisationController {
	return OrganisationController{
		organisationRepository: repo,
	}
}

func (organisationController OrganisationController) CreateOrganisation(org models.Organisation) (interface{}, error) {

	id, err := organisationController.organisationRepository.CreateOrganisation(&org)

	if err != nil {
		return nil, err
	}

	return id, nil

}

func (organisationController OrganisationController) GetOrganisation(ctx *gin.Context) {

	data, err := organisationController.organisationRepository.GetOrganisation()

	if err != nil {
		helpers.ResponseDispatch(ctx, nil, err, http.StatusInternalServerError)
		return
	}

	helpers.ResponseDispatch(ctx, data, nil, http.StatusOK)
}
