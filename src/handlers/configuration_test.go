package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type GetConfigurationByKeyParams struct {
	Key string `uri:"key" binding:"required"`
}

func GenerateHandlerRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/v1/config/:key", GetConfigurationByKeyHandler)

	return router
}

func TestSuccessGetConfigurationByKeyHandler(t *testing.T) {
	router := GenerateHandlerRouter()
	configKey := "amPointCostValue"

	req, err := http.NewRequest(http.MethodGet, "/v1/config/"+configKey, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)
}

func TestFailedGetConfigurationByKeyHandler(t *testing.T) {
	router := GenerateHandlerRouter()
	configKey := "testKey"

	req, err := http.NewRequest(http.MethodGet, "/v1/config/"+configKey, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 404)
}
