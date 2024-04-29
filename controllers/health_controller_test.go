package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandleHealth(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	mockHealthRepository := mocks.NewMockIHealthRepository(gomockController)
	healthController := NewController(mockHealthRepository)

	appHealth := models.AppHealth{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	mockHealthRepository.EXPECT().GetAppHealth().Return(appHealth)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/health", nil)

	healthController.HandleHealth(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
