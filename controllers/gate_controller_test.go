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
var CREATE_NEW_GATE = "/gate"

var gateMockService *mocks.MockIGateRepository
var gateController *GateControllerRepository
var gateContext *gin.Context

func beforeEachGateTest(t *testing.T) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()

	gateMockService = mocks.NewMockIGateRepository(mockControl)
	gateController = NewGateRepository(gateMockService)
	recorder := httptest.NewRecorder()
	gateContext, _ = gin.CreateTestContext(recorder)
}

func TestHandleGetGates(t *testing.T) {
	beforeEachGateTest(t)
	mockGates := make([]models.Gate, 3)
	mockGates = append(mockGates, models.Gate{FloorNumber: 2, GateNumber: 1})
	gateContext.Request, _ = http.NewRequest("GET", GET_ALL_GATES, nil)
	gateMockService.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(mockGates, nil)

	gateController.HandleGetGates(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleGetGatesWhenServiceReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	gateMockService.EXPECT().GetGates(gomock.Any(), gomock.Any()).Return(nil, errors.New("Invalid"))

	gateContext.Request, _ = http.NewRequest("GET", GET_ALL_GATES, nil)

	gateController.HandleGetGates(gateContext)

	assert.Equal(t, http.StatusInternalServerError, gateContext.Writer.Status())
}

func TestHandleGetGateById(t *testing.T) {
	beforeEachGateTest(t)
	mockGates := models.Gate{FloorNumber: 2, GateNumber: 1}
	gateMockService.EXPECT().GetGateById(gomock.Any()).Return(&mockGates, nil)

	gateContext.Request, _ = http.NewRequest("GET", GET_GATE_BY_ID, nil)

	gateController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusOK, gateContext.Writer.Status())
}

func TestHandleGetGateByIdWhenGateIdDoesNotExist(t *testing.T) {
	beforeEachGateTest(t)
	gateMockService.EXPECT().GetGateById(gomock.Any()).Return(nil, errors.New("SQLSTATE 22P02"))

	gateContext.Request, _ = http.NewRequest("GET", GET_GATE_BY_ID, nil)

	gateController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusNotFound, gateContext.Writer.Status())
}

func TestHandleGetGateByIdWhenServiceReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	gateMockService.EXPECT().GetGateById(gomock.Any()).Return(nil, errors.New("invalid"))
	gateContext.Request, _ = http.NewRequest("GET", GET_GATE_BY_ID, nil)

	gateController.HandleGetGateById(gateContext)

	assert.Equal(t, http.StatusInternalServerError, gateContext.Writer.Status())
}

func TestHandleCreateNewGate(t *testing.T) {
	beforeEachGateTest(t)
	ExpectedGateNumber := 1
	ExpectedFloorNumber := 1
	gateMockService.EXPECT().CreateNewGate(gomock.Any()).Return(nil)
	reqBody := `{"gate_number" : 1, "floor_number" : 1}`
	var gate models.Gate
	gateContext.Request, _ = http.NewRequest("POST", "CREATE_NEW_GATE", strings.NewReader(reqBody))
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
	gateContext.Request, _ = http.NewRequest("POST", CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhenTheRequestPayloadIsEmpty(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{}`
	gateContext.Request, _ = http.NewRequest("POST", CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhenTheMandatoryKeyIsAbsent(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":2}`
	gateContext.Request, _ = http.NewRequest("POST", CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":"one", "floor_number":20}`
	gateContext.Request, _ = http.NewRequest("POST", CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleCreateNewGateWhereErrorIsThrownInServiceLayer(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":3, "floor_number":6}`
	gateMockService.EXPECT().CreateNewGate(gomock.Any()).Return(errors.New("invalid Request"))
	gateContext.Request, _ = http.NewRequest("POST", CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}
