package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler This a controller that handle about home api
func PingHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
