package routers

import (
	Handlers "github.com/IOsonoTAN/go-boilerplate/src/handlers"
	"github.com/gin-gonic/gin"
)

// RoutesV1 is a sub-router for version one
func RoutesV1(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/v1")
	{
		v1.GET("/config/:key", Handlers.GetConfigurationByKeyHandler)
		v1.POST("/config", Handlers.PostNewConfigurationHandler)
	}
	return router
}
