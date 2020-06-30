package handlers

import (
	"net/http"

	MongoDB "github.com/IOsonoTAN/go-boilerplate/src/models/mongodb"
	"github.com/gin-gonic/gin"
)

// GetConfigurationByKeyHandler This a controller that handle about home api
func GetConfigurationByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")

	result, err := MongoDB.GetConfigurationByKey(key)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "The key is not found"}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
