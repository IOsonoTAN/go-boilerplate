package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigurationByKeyHandler(t *testing.T) {
	// key := "campaignStatus"
	endpoint := "/v1/config/:key"
	res := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, req := gin.CreateTestContext(res)
	ctx.Params = []gin.Param{gin.Param{Value: "v"}}

	req.GET(endpoint, GetConfigurationByKeyHandler)
	ctx.Request, _ = http.NewRequest(http.MethodGet, endpoint, nil)
	req.ServeHTTP(res, ctx.Request)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "OK", res.Body.String())
}
