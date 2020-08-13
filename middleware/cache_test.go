package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDisableCache(t *testing.T) {
	router := setupRouterCache()

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, "test", responseRecorder.Body.String())

	assert.Equal(t, "0", responseRecorder.Header().Get("Expires"))
	assert.Equal(t, "no-cache", responseRecorder.Header().Get("Pragma"))
	assert.Equal(t, "no-store", responseRecorder.Header().Get("Cache-Control"))
}

func setupRouterCache() *gin.Engine {
	router := gin.Default()
	router.Use(DisableCache())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "test")
	})

	return router
}
