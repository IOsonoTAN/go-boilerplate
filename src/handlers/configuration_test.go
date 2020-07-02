package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func GenerateRouterAndRequest(key string) (*gin.Engine, *http.Request) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/v1/config/:key", GetConfigurationByKeyHandler)

	req, err := http.NewRequest(http.MethodGet, "/v1/config/"+key, nil)
	if err != nil {
		fmt.Println(err)
	}

	return router, req
}

func TestSuccessGetConfigurationByKeyHandler(t *testing.T) {
	router, req := GenerateRouterAndRequest("amPointCostValue")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)
}

func TestFailedGetConfigurationByKeyHandler(t *testing.T) {
	router, req := GenerateRouterAndRequest("testKey")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 404)
}
