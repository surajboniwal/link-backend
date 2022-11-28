package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
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

func (organisationController OrganisationController) GetOrganisation(ctx *gin.Context) {

	data, err := organisationController.organisationRepository.GetOrganisation()

	if err != nil {
		helpers.ResponseDispatch(ctx, nil, err, http.StatusInternalServerError)
		return
	}

	helpers.ResponseDispatch(ctx, data, nil, http.StatusOK)
}
