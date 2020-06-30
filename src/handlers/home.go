package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler This a controller that handle about home api
func HomeHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
