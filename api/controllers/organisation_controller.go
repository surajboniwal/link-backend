package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/helpers"
	"github.com/surajboniwal/link-backend/api/models"
	"github.com/surajboniwal/link-backend/api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrganisationController struct {
	organisationRepository repositories.OrganisationRepository
	memberRepository       repositories.MemberRepository
}

func NewOrganisationController(repo repositories.OrganisationRepository, memberRepository repositories.MemberRepository) OrganisationController {
	return OrganisationController{
		organisationRepository: repo,
		memberRepository:       memberRepository,
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

func (controller OrganisationController) CreateMember(context *gin.Context) {

	var params models.Member

	if err := context.BindJSON(&params); err != nil {
		helpers.ResponseDispatch(context, nil, err.Error(), http.StatusBadRequest)
		return
	}

	orgId := context.Param("id")

	orgIdObject, orgIdErr := primitive.ObjectIDFromHex(orgId)

	fmt.Print(orgIdObject)

	if orgIdErr != nil {
		helpers.ResponseDispatch(context, nil, "Invalid ID", http.StatusBadRequest)
		return
	}

	params.ID = primitive.NewObjectID()
	params.OrganisationID = orgIdObject

	id, err := controller.memberRepository.CreateMember(&params)

	if err != nil {
		helpers.ResponseDispatch(context, nil, err, http.StatusBadRequest)
		return
	}

	helpers.ResponseDispatch(context, id, nil, http.StatusCreated)
}
