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
	normalRequest, _ := http.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(responseRecorder, normalRequest)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "test", responseRecorder.Body.String())

	faviconResponseRecorder := httptest.NewRecorder()
	faviconRequest, _ := http.NewRequest(http.MethodGet, "/favicon.ico", nil)
	router.ServeHTTP(faviconResponseRecorder, faviconRequest)

	assert.Equal(t, http.StatusNoContent, faviconResponseRecorder.Code)
	assert.Equal(t, "", faviconResponseRecorder.Body.String())
}

func setupRouterFavicon() *gin.Engine {
	router := gin.Default()
	router.Use(HandleFaviconRequests())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "test")
	})

	return router
}
