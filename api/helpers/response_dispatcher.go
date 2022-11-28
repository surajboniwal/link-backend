package helpers

import (
	"github.com/gin-gonic/gin"
)

func ResponseDispatch(context *gin.Context, data interface{}, err interface{}, statusCode int) {
	context.JSON(
		statusCode,
		gin.H{
			"data":  data,
			"error": err,
		},
	)
}
