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

func TestHandleAirlineController(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	airlinesMock := make([]models.Airlines, 0)
	mockService.EXPECT().GetAirline(gomock.Any()).Return(airlinesMock, nil)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Request, _ = http.NewRequest("GET", "/airline", nil)
	controllerRepo.HandleAirline(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
