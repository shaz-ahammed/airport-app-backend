package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleFaviconRequests(t *testing.T) {
	router := setupRouterFavicon()

	responseRecorder := httptest.NewRecorder()
	normalRequest, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(responseRecorder, normalRequest)

	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, "test", responseRecorder.Body.String())

	faviconResponseRecorder := httptest.NewRecorder()
	faviconRequest, _ := http.NewRequest("GET", "/favicon.ico", nil)
	router.ServeHTTP(faviconResponseRecorder, faviconRequest)

	assert.Equal(t, 204, faviconResponseRecorder.Code)
	assert.Equal(t, "", faviconResponseRecorder.Body.String())
}

func setupRouterFavicon() *gin.Engine {
	router := gin.Default()
	router.Use(HandleFaviconRequests())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "test")
	})

	return router
}