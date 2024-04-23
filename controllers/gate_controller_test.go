package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHandleCreateNewGate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIGateRepository(mockCtrl)
	controllerRepo := NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	gate := models.Gate{
		GateNumber:  1,
		FloorNumber: 1,
	}
	mockService.EXPECT().CreateNewGate(&gate).Return(nil)
	reqBody := `{"gate_number" : 1,"floor_number" : 1}`
	var gateModel models.Gate

	ctx.Request, _ = http.NewRequest("POST", "/gate", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewGate(ctx)
	err := json.Unmarshal([]byte(reqBody), &gateModel)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
	assert.NoError(t, err)
	assert.Equal(t, gate.GateNumber, gateModel.GateNumber)
	assert.Equal(t, gate.FloorNumber, gateModel.FloorNumber)
}

func TestHandleCreateNewGateWhenTheMandatoryValueIsAbsent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIGateRepository(mockCtrl)
	controllerRepo := NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{"gate_number":,"floor_number":2}`

	ctx.Request, _ = http.NewRequest("POST", "/gate", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewGate(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestHandleCreateNewGateWhenTheRequestPayloadIsEmpty(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIGateRepository(mockCtrl)
	controllerRepo := NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{}`

	ctx.Request, _ = http.NewRequest("POST", "/gate", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewGate(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestHandleCreateNewGateWhenTheMandatoryKeyIsAbsent(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIGateRepository(mockCtrl)
	controllerRepo := NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{"gate_number":2}`

	ctx.Request, _ = http.NewRequest("POST", "/gate", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewGate(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestHandleCreateNewGateWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIGateRepository(mockCtrl)
	controllerRepo := NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{"gate_number":"one","floor_number":20}`

	ctx.Request, _ = http.NewRequest("POST", "/gate", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewGate(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}

func TestHandleCreateNewGateWhereErrorIsThrownInServiceLayer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIGateRepository(mockCtrl)
	controllerRepo := NewGateRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	gate := models.Gate{
		GateNumber:  3,
		FloorNumber: 6,
	}
	reqBody := `{"gate_number":3, "floor_number":6}`

	mockService.EXPECT().CreateNewGate(&gate).Return(errors.New("invalid Request"))
	ctx.Request, _ = http.NewRequest("POST", "/gate", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewGate(ctx)

	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}
