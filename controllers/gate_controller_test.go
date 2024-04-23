package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockService *mocks.MockIGateRepository
var mockController *GateControllerRepository
var ctx *gin.Context

func beforeEach(t *testing.T) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()
 
	mockService = mocks.NewMockIGateRepository(mockControl)
	mockController = NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)
}

func TestHandleGetGates(t *testing.T) {
	beforeEach(t)
	mockGates := make([]models.Gate, 3)
	mockGates = append(mockGates, models.Gate{FloorNumber: 2, GateNumber: 1})

	mockService.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(mockGates, nil)
	ctx.Request, _ = http.NewRequest("GET", "/gates", nil)
	mockController.HandleGetGates(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestHandleGetGateById(t *testing.T) {
	beforeEach(t)
	mockGates := models.Gate{FloorNumber: 2, GateNumber: 1}

	mockService.EXPECT().GetGateById(gomock.Any()).Return(&mockGates, nil)
	ctx.Request, _ = http.NewRequest("GET", "/gates/123", nil)
	mockController.HandleGetGateById(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
