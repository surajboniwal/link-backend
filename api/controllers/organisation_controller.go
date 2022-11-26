package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/models"
	"github.com/surajboniwal/link-backend/api/repositories"
)

type OrganisationController struct {
	OrganisationRepository repositories.OrganisationRepository
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

	organisationController.OrganisationRepository.CreateOrganisation(&org)

	ctx.JSON(200, org)
}

func (organisationController OrganisationController) GetOrganisation(ctx *gin.Context) {

	data := organisationController.OrganisationRepository.GetOrganisation()

	ctx.JSON(200, gin.H{
		"status":  true,
		"message": data,
	})
}
