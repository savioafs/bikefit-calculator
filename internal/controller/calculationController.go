package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/bikefit-calculator/internal/calculation"
	"github.com/savioafs/bikefit-calculator/internal/entity"
)

func Calculate(ctx *gin.Context) {
	var measuresInput entity.Measures

	err := ctx.BindJSON(&measuresInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid data to calculate",
		})
		return
	}

	result := calculation.Resultcalculation(measuresInput)

	ctx.JSON(http.StatusOK, result)
}
