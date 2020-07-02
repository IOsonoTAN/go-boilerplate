package routers

import (
	Handlers "github.com/IOsonoTAN/go-boilerplate/src/handlers"
	"github.com/gin-gonic/gin"
)

// Routes is a function for create a router in root path
func Routes(router *gin.Engine) *gin.Engine {
	router.GET("/ping", Handlers.PingHandler)

	return router
}
