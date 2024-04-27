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

var GATE_BY_ID = "/gate/123"
var GET_ALL_GATES = "gates"
var CREATE_NEW_GATE = "/gate"

var gateMockRepository *mocks.MockIGateRepository
var gateController *GateControllerRepository
var gateContext *gin.Context

func beforeEachGateTest(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	gateMockRepository = mocks.NewMockIGateRepository(gomockController)
	gateController = NewGateRepository(gateMockRepository)
	gateContext, _ = gin.CreateTestContext(httptest.NewRecorder())
}

func TestHandleGetGates(t *testing.T) {
	beforeEachGateTest(t)
	mockGates := make([]models.Gate, 3)
	mockGates = append(mockGates, models.Gate{FloorNumber: 2, GateNumber: 1})
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GET_ALL_GATES, nil)
	gateMockRepository.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(mockGates, nil)

	gateController.HandleGetGates(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleGetGatesWhenRepositoryReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	gateMockRepository.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(nil, errors.New("Invalid"))
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GET_ALL_GATES, nil)

	gateController.HandleGetGates(gateContext)

	assert.Equal(t, http.StatusInternalServerError, gateContext.Writer.Status())
}

func TestHandleGetGateById(t *testing.T) {
	beforeEachGateTest(t)
	mockGates := models.Gate{FloorNumber: 2, GateNumber: 1}
	gateMockRepository.EXPECT().GetGateById(gomock.Any()).Return(&mockGates, nil)
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GATE_BY_ID, nil)

	gateController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleGetGateByIdWhenGateIdDoesNotExist(t *testing.T) {
	beforeEachGateTest(t)
	gateMockRepository.EXPECT().GetGateById(gomock.Any()).Return(nil, errors.New("SQLSTATE 22P02"))
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GATE_BY_ID, nil)

	gateController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusNotFound, gateContext.Writer.Status())
}

func TestHandleGetGateByIdWhenRepositoryReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	gateMockRepository.EXPECT().GetGateById(gomock.Any()).Return(nil, errors.New("invalid"))
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GATE_BY_ID, nil)

	gateController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusInternalServerError, gateContext.Writer.Status())
}

func TestHandleCreateNewGate(t *testing.T) {
	beforeEachGateTest(t)
	ExpectedGateNumber := 1
	ExpectedFloorNumber := 1
	gateMockRepository.EXPECT().CreateNewGate(gomock.Any()).Return(nil)
	reqBody := `{"gate_number" : 1, "floor_number" : 1}`
	var gate models.Gate
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))
	err := json.Unmarshal([]byte(reqBody), &gate)

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
	assert.NoError(t, err)
	assert.Equal(t, ExpectedGateNumber, gate.GateNumber)
	assert.Equal(t, ExpectedFloorNumber, gate.FloorNumber)
}

func TestHandleCreateNewGateWhenTheMandatoryValueIsAbsent(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":, "floor_number":2}`
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhenTheRequestPayloadIsEmpty(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{}`
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhenTheMandatoryKeyIsAbsent(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":2}`
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":"one", "floor_number":20}`
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhereErrorIsThrownInRepositoryLayer(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":3, "floor_number":6}`
	gateMockRepository.EXPECT().CreateNewGate(gomock.Any()).Return(errors.New("invalid Request"))
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleUpdateGate(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":3, "floor_number":6}`
	gateMockRepository.EXPECT().UpdateGate(gomock.Any(), gomock.Any()).Return(nil)
	gateContext.Request, _ = http.NewRequest(http.MethodPut, GATE_BY_ID, strings.NewReader(reqBody))

	gateController.HandleUpdateGate(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleUpdateGateWhenRequiredFieldIsNotGiven(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":3}`
	gateMockRepository.EXPECT().UpdateGate(gomock.Any(), gomock.Any()).Return(nil)
	gateContext.Request, _ = http.NewRequest(http.MethodPut, GATE_BY_ID, strings.NewReader(reqBody))

	gateController.HandleUpdateGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleUpdateGateWhenRepositoryThrowsError(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":3, "floor_number":6}`
	gateMockRepository.EXPECT().UpdateGate(gomock.Any(), gomock.Any()).Return(errors.New("Invalid"))
	gateContext.Request, _ = http.NewRequest(http.MethodPut, GATE_BY_ID, strings.NewReader(reqBody))

	gateController.HandleUpdateGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}
