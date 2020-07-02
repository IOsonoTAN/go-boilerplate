package main

import (
	"log"
	"os"

	Routers "github.com/IOsonoTAN/go-boilerplate/src/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	Routers.Routes(router)
	Routers.RoutesV1(router)

	return router
}

func main() {
	dotEnvError := godotenv.Load()
	if dotEnvError != nil {
		log.Fatal("can't loading .env file")
	}

	appEnv := os.Getenv("ENV")
	if appEnv == "production" {
		gin.SetMode("release")
	}

	appListenPort := os.Getenv("PORT")
	if len(appListenPort) < 1 {
		appListenPort = "3000"
	}

	router := setupRouter()

	router.Run(":" + appListenPort)
}
