package handlers

import (
	"fmt"
	"net/http"

	MongoDB "github.com/IOsonoTAN/go-boilerplate/src/models/mongodb"
	"github.com/gin-gonic/gin"
)

// GetConfigurationByKeyHandler a controller of configuration that get value from database
func GetConfigurationByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	fmt.Println(key)

	result, err := MongoDB.GetConfigurationByKey(key)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "The key is not found"}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

// NewConfiguration a struct for handle of request body as json
type NewConfiguration struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
	Type  string `bson:"type"`
}

// PostNewConfigurationHandler a function to create a new configuration into database
func PostNewConfigurationHandler(ctx *gin.Context) {
	var newConfiguration NewConfiguration
	ctx.BindJSON(&newConfiguration)

	result, err := MongoDB.InsertNewConfiguration(newConfiguration.Key, newConfiguration.Value, newConfiguration.Type)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "Something went wrong, Please try again."}})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
