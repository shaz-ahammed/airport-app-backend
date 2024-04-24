package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTraceSpanTags(t *testing.T) {
	router := setupRouterTrace()

	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.MethodGet, request.Method)
	assert.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)
	assert.Equal(t, "test", responseRecorder.Body.String())
}

func setupRouterTrace() *gin.Engine {
	router := gin.Default()
	router.Use(TraceSpanTags())

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "test")
	})
	return router
}
