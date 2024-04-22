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

func TestHandleGetGatesController(t *testing.T) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()

	mockService := mocks.NewMockIGateRepository(mockControl)
	mockController := NewGateRepository(mockService)
	mockGates := make([]models.Gate, 3)
	mockGates = append(mockGates, models.Gate{FloorNumber: 2, GateNumber: 1})
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(mockGates, nil)
	ctx.Request, _ = http.NewRequest("GET", "/gates", nil)
	mockController.HandleGetGates(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestHandleGetGatesByIDController(t *testing.T) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()
	
	mockService := mocks.NewMockIGateRepository(mockControl)
	mockController := NewGateRepository(mockService)
	mockGates := models.Gate{FloorNumber: 2, GateNumber: 1}
	mockService.EXPECT().GetGateByID(gomock.Any()).Return(&mockGates, nil)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request, _ = http.NewRequest("GET", "/gates/123", nil)
	mockController.HandleGetGateByID(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}