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
	mockRepository := mocks.NewMockIHealthRepository(gomockController)
	controllerRepo := NewControllerRepository(mockRepository)
	appHealthMock := models.AppHealth{}
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	mockRepository.EXPECT().GetAppHealth().Return(appHealthMock)
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/health", nil)

	controllerRepo.HandleHealth(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
