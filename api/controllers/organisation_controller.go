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

func (organisationController OrganisationController) CreateOrganisation(ctx *gin.Context) {

	org := models.Organisation{
		Name:        "Hello",
		Revenue:     "1000",
		Phone:       "9974849404",
		Website:     "https://surajboniwal.dev",
		Location:    "India",
		Designation: "Founder",
		Industry:    "Tech",
	}

	organisationController.organisationRepository.CreateOrganisation(&org)

	helpers.ResponseDispatch(ctx, org, nil, http.StatusOK)
}

func (organisationController OrganisationController) GetOrganisation(ctx *gin.Context) {

	data := organisationController.organisationRepository.GetOrganisation()

	helpers.ResponseDispatch(ctx, data, nil, http.StatusOK)
}
