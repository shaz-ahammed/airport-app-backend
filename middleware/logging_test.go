package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestZerologConsoleRequestLogging(t *testing.T) {
	router := setupRouterLogging()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/test", nil)

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "test", responseRecorder.Body.String())

	undefinedResponseRecorder := httptest.NewRecorder()
	undefinedRequest := httptest.NewRequest("GET", "/non-defined", nil)

	router.ServeHTTP(undefinedResponseRecorder, undefinedRequest)

	assert.Equal(t, http.StatusNotFound, undefinedResponseRecorder.Code)
}

func setupRouterLogging() *gin.Engine {
	router := gin.Default()
	router.Use(ZerologConsoleRequestLogging())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "test")
	})
	return router
}
