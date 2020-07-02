package handlers

import (
	"net/http"

	MongoDB "github.com/IOsonoTAN/go-boilerplate/src/models/mongodb"
	"github.com/gin-gonic/gin"
)

// GetConfigurationByKeyParams a struct that bind for uri
type GetConfigurationByKeyParams struct {
	Key string `uri:"key" binding:"required"`
}

// GetConfigurationByKeyHandler a controller of configuration that get value from database
func GetConfigurationByKeyHandler(ctx *gin.Context) {
	var params GetConfigurationByKeyParams
	ctx.BindUri(&params) // same as ctx.Param("key")

	result, err := MongoDB.GetConfigurationByKey(params.Key)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "The key is not found"}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

// NewConfiguration a struct for handle of request body as json
type NewConfiguration struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// PostNewConfigurationHandler a function to create a new configuration into database
func PostNewConfigurationHandler(ctx *gin.Context) {
	var newConfiguration NewConfiguration
	ctx.BindJSON(&newConfiguration)

	result, err := MongoDB.InsertNewConfiguration(
		newConfiguration.Key,
		newConfiguration.Value,
		newConfiguration.Type,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "Something went wrong, Please try again."}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
