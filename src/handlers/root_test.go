package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	res := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, req := gin.CreateTestContext(res)

	req.GET("/ping", PingHandler)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/ping", nil)
	req.ServeHTTP(res, ctx.Request)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "OK", res.Body.String())
}
