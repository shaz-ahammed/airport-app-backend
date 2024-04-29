package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"airport-app-backend/models/factory"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

var mockGateRepository *mocks.MockIGateRepository
var gateController *GateController
var gateContext *gin.Context
var gateResponseRecorder *httptest.ResponseRecorder

func beforeEachGateTest(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	mockGateRepository = mocks.NewMockIGateRepository(gomockController)
	gateController = NewGateController(mockGateRepository)
	gateResponseRecorder = httptest.NewRecorder()
	gateContext, _ = gin.CreateTestContext(gateResponseRecorder)
}

func TestHandleGetAllGates(t *testing.T) {
	beforeEachGateTest(t)
	var gates []models.Gate
	gate1 := factory.ConstructGate()
	gates = append(gates, gate1)
	gate2 := factory.ConstructGate()
	gates = append(gates, gate2)
	gate3 := factory.ConstructGate()
	gates = append(gates, gate3)
	mockGateRepository.EXPECT().GetAllGates(gomock.Any(), gomock.Any()).Return(gates, nil)

	gateController.HandleGetAllGates(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var gatesFromResponse []models.Gate
	json.Unmarshal([]byte(responseBody), &gatesFromResponse)

	assert.Equal(t, 3, len(gatesFromResponse))
	assert.Contains(t, gatesFromResponse, gate1)
	assert.Contains(t, gatesFromResponse, gate2)
	assert.Contains(t, gatesFromResponse, gate3)
}

// TODO: All tests beyond this line need to be verified/rewritten

func TestHandleGetAllGatesWhenRepositoryReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	mockGateRepository.EXPECT().GetAllGates(gomock.Any(), gomock.Any()).Return(nil, errors.New("Invalid"))
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GET_ALL_GATES, nil)

	gateController.HandleGetAllGates(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, fmt.Sprintf("{\"error\":\"Failed to fetch gates\"}"), string(responseBody))
}

func TestHandleGetGate(t *testing.T) {
	beforeEachGateTest(t)
	gateId := "123"
	gate := factory.ConstructGate()
	mockGateRepository.EXPECT().GetGate(gateId).Return(&gate, nil)
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GATE_BY_ID, nil)
	gateContext.AddParam("id", gateId)

	gateController.HandleGetGate(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var gateFromResponse models.Gate
	json.Unmarshal([]byte(responseBody), &gateFromResponse)

	assert.Equal(t, gate, gateFromResponse)
}

func TestHandleGetGateWhenRepositoryReturnsError(t *testing.T) {
	beforeEachGateTest(t)
	InvalidGateId := "123"
	mockGateRepository.EXPECT().GetGate(InvalidGateId).Return(nil, errors.New("Invalid"))
	gateContext.Request, _ = http.NewRequest(http.MethodGet, GATE_BY_ID, nil)
	gateContext.AddParam("id", InvalidGateId)

	gateController.HandleGetGate(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, fmt.Sprintf("{\"Error\":\"Incorrect gate id: %s\"}", InvalidGateId), string(responseBody))
}

func TestHandleCreateNewGate(t *testing.T) {
	beforeEachGateTest(t)
	gate := factory.ConstructGate()
	reqBody, _ := json.Marshal(&gate)
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(string(reqBody)))
	mockGateRepository.EXPECT().CreateNewGate(gomock.Any()).Return(nil)

	gateController.HandleCreateNewGate(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
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

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, fmt.Sprintf("{\"error\":\"Key: 'Gate.GateNumber' Error:Field validation for 'GateNumber' failed on the 'required' tag\\nKey: 'Gate.FloorNumber' Error:Field validation for 'FloorNumber' failed on the 'required' tag\"}"), string(responseBody))
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
	mockGateRepository.EXPECT().CreateNewGate(gomock.Any()).Return(errors.New("invalid Request"))
	gateContext.Request, _ = http.NewRequest(http.MethodPost, CREATE_NEW_GATE, strings.NewReader(reqBody))

	gateController.HandleCreateNewGate(gateContext)

	assert.Equal(t, http.StatusBadRequest, gateContext.Writer.Status())
}

func TestHandleUpdateGate(t *testing.T) {
	beforeEachGateTest(t)
	gateId := "1"
	gate := factory.ConstructGate()
	reqBody, _ := json.Marshal(gate)
	gateContext.AddParam("id", gateId)
	mockGateRepository.EXPECT().UpdateGate(gateId, gate).Return(nil)
	gateContext.Request, _ = http.NewRequest(http.MethodPut, GATE_BY_ID, strings.NewReader(string(reqBody)))

	gateController.HandleUpdateGate(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestHandleUpdateGateWhenRequiredFieldIsNotGiven(t *testing.T) {
	beforeEachGateTest(t)
	reqBody := `{"gate_number":3}`
	gateContext.Request, _ = http.NewRequest(http.MethodPut, GATE_BY_ID, strings.NewReader(reqBody))

	gateController.HandleUpdateGate(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, fmt.Sprintf("{\"error\":\"Key: 'Gate.FloorNumber' Error:Field validation for 'FloorNumber' failed on the 'required' tag\"}"), string(responseBody))
}

func TestHandleUpdateGateWhenRepositoryThrowsError(t *testing.T) {
	beforeEachGateTest(t)
	invalidId := "1"
	gate := factory.ConstructGate()
	gateContext.AddParam("id", invalidId)
	reqBody, _ := json.Marshal(gate)
	mockGateRepository.EXPECT().UpdateGate(invalidId, gate).Return(errors.New("Invalid"))
	gateContext.Request, _ = http.NewRequest(http.MethodPut, GATE_BY_ID, strings.NewReader(string(reqBody)))

	gateController.HandleUpdateGate(gateContext)

	response := gateResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}
