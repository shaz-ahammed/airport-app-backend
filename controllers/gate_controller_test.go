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

var GET_GATE_BY_ID = "/gate/123"
var GET_ALL_GATES = "gates"

var gateMockService *mocks.MockIGateRepository
var gateMockController *GateControllerRepository
var gateContext *gin.Context

func beforeEachGateTest(t *testing.T) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()
 
	gateMockService = mocks.NewMockIGateRepository(mockControl)
	gateMockController = NewGateRepository(gateMockService)
	recorder := httptest.NewRecorder()
	gateContext, _ = gin.CreateTestContext(recorder)
}

func TestHandleGetGates(t *testing.T) {
	beforeEachGateTest(t)
	mockGates := make([]models.Gate, 3)
	mockGates = append(mockGates, models.Gate{FloorNumber: 2, GateNumber: 1})
	gateMockService.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(mockGates, nil)
	
	gateContext.Request, _ = http.NewRequest("GET", GET_ALL_GATES, nil)
	gateMockController.HandleGetGates(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleGetGatesWhenServiceReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	gateMockService.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(nil, errors.New("Invalid"))

	gateContext.Request, _ = http.NewRequest("GET", GET_ALL_GATES, nil)
	gateMockController.HandleGetGates(gateContext)

	assert.Equal(t, http.StatusInternalServerError, gateContext.Writer.Status())
}

func TestHandleGetGateById(t *testing.T) {
	beforeEachGateTest(t)
	mockGates := models.Gate{FloorNumber: 2, GateNumber: 1}
	gateMockService.EXPECT().GetGateById(gomock.Any()).Return(&mockGates, nil)

	gateContext.Request, _ = http.NewRequest("GET", GET_GATE_BY_ID, nil)
	gateMockController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleGetGateByIdWhenGateIdDoesNotExist(t *testing.T) {
	beforeEachGateTest(t)
	gateMockService.EXPECT().GetGateById(gomock.Any()).Return(nil, errors.New("SQLSTATE 22P02"))

	gateContext.Request, _ = http.NewRequest("GET", GET_GATE_BY_ID, nil)
	gateMockController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusNotFound, gateContext.Writer.Status())
}

func TestHandleGetGateByIdWhenServiceReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	gateMockService.EXPECT().GetGateById(gomock.Any()).Return(nil, errors.New("Invalid"))

	gateContext.Request, _ = http.NewRequest("GET", GET_GATE_BY_ID, nil)
	gateMockController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusInternalServerError, gateContext.Writer.Status())
}
