package controllers

import (
	mocker "airport-app-backend/mocks/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandleHealthController(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockService := mocker.NewServiceRepository(mockCtrl)
	controllerRepo := NewControllerRepository(mockService)
	router := gin.Default()
	router.GET("/health", controllerRepo.HandleHealth)
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
