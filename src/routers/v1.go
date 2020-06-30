package routers

import (
	Handlers "github.com/IOsonoTAN/go-boilerplate/src/handlers"
	"github.com/gin-gonic/gin"
)

// RoutesVersionOne is a sub-router for version one
func RoutesVersionOne(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/v1")
	{
		v1.GET("/config/:key", Handlers.GetConfigurationByKeyHandler)
	}
	return router
}
