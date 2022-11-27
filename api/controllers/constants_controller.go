package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/surajboniwal/link-backend/api/repositories"
)

type ConstantsController struct {
	constantsRepository repositories.ConstantsRepository
}

func NewConstantsController(repo repositories.ConstantsRepository) ConstantsController {
	return ConstantsController{
		constantsRepository: repo,
	}
}

func (constantsController ConstantsController) GetConstants(context *gin.Context) {
	constants := constantsController.constantsRepository.GetConstants()

	context.JSON(
		http.StatusOK,
		gin.H{
			"status": true,
			"data":   constants,
		},
	)
}