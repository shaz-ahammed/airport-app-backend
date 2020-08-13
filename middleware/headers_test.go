package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddSecurityHeadersNoTls(t *testing.T) {
	router := setupRouterSecurityHeaders(false)

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "test", responseRecorder.Body.String())

	assert.Equal(t, "DENY", responseRecorder.Header().Get("X-Frame-Options"))
	assert.Equal(t, "nosniff", responseRecorder.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "no-referrer", responseRecorder.Header().Get("Referrer-Policy"))
	assert.Equal(t, "default-src 'none';", responseRecorder.Header().Get("Content-Security-Policy"))
}

func TestAddSecurityHeadersWithTlsEnabled(t *testing.T) {
	router := setupRouterSecurityHeaders(true)

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "test", responseRecorder.Body.String())

	assert.Equal(t, "DENY", responseRecorder.Header().Get("X-Frame-Options"))
	assert.Equal(t, "nosniff", responseRecorder.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "no-referrer", responseRecorder.Header().Get("Referrer-Policy"))
	assert.Equal(t, "default-src 'none'; upgrade-insecure-requests;", responseRecorder.Header().Get("Content-Security-Policy"))
	assert.Equal(t, "max-age=94608000 ;includeSubDomains; preload", responseRecorder.Header().Get("Strict-Transport-Security"))
}

func setupRouterSecurityHeaders(shouldEnableTls bool) *gin.Engine {
	router := gin.Default()
	router.Use(AddSecurityHeaders(shouldEnableTls))

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "test")
	})

	return router
}
