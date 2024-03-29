package controllers

import (
"net/http"
"net/http/httptest"
"testing"
"github.com/gin-gonic/gin"
"github.com/stretchr/testify/assert"
)

func TestHandleHealthController(t *testing.T) {
  router := gin.Default()
  router.GET("/health", HandleHealth)
  req, err := http.NewRequest("GET", "/health", nil)
  if err != nil {
    t.Fatal(err)
  }
  rr := httptest.NewRecorder()
  router.ServeHTTP(rr, req)
  assert.Equal(t,http.StatusOK,rr.Code)
}
