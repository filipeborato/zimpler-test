package controller

import (
	"net/http"
	"zimpler-test/entity"
	"zimpler-test/service"

	"github.com/gin-gonic/gin"
)

func CandyStore(ctx *gin.Context) entity.TopRate {
	topRate, err := service.GoQueryCandy()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Request Error", "error": err})
		return nil
	}

	return topRate
}
