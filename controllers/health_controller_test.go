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

func TestHandleHealthController(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIHealthRepository(mockCtrl)
	controllerRepo := NewControllerRepository(mockService)
	appHealthMock := models.AppHealth{Goroutines: 5}
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService.EXPECT().GetAppHealth(gomock.Any(), ctx).Return(appHealthMock)

	ctx.Request, _ = http.NewRequest("GET", "/health", nil)
	controllerRepo.HandleHealth(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
