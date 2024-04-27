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

var GET_ALL_AIRLINES = "/airlines"
var AIRLINE_BY_ID = "/airline/123"
var POST_AIRLINE = "/airline"

var airlineMockRepository *mocks.MockIAirlineRepository
var airlineController *AirlineControllerRepository
var airlineContext *gin.Context
var responseRecorder *httptest.ResponseRecorder

func beforeEachAirlineTest(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	airlineMockRepository = mocks.NewMockIAirlineRepository(gomockController)
	airlineController = NewAirlineControllerRepository(airlineMockRepository)
	responseRecorder = httptest.NewRecorder()
	airlineContext, _ = gin.CreateTestContext(responseRecorder)
}

func TestHandleGetAllAirlines(t *testing.T) {
	beforeEachAirlineTest(t)
	var airlines []models.Airline
	airline1 := factory.ConstructAirline()
	airlines = append(airlines, airline1)
	airline2 := factory.ConstructAirline()
	airlines = append(airlines, airline2)
	airline3 := factory.ConstructAirline()
	airlines = append(airlines, airline3)

	airlineMockRepository.EXPECT().GetAllAirlines(gomock.Any()).Return(airlines, nil)
	airlineContext.Request, _ = http.NewRequest(http.MethodGet, GET_ALL_AIRLINES, nil)

	airlineController.HandleGetAllAirlines(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var airlinesFromResponse []models.Airline
	json.Unmarshal([]byte(responseBody), &airlinesFromResponse)

	assert.Equal(t, 3, len(airlinesFromResponse))
	assert.Contains(t, airlinesFromResponse, airline1)
	assert.Contains(t, airlinesFromResponse, airline2)
	assert.Contains(t, airlinesFromResponse, airline3)
}

// TODO: InternalServerError scenario for GetAllAirlines

func TestHandleGetAirlineById(t *testing.T) {
	beforeEachAirlineTest(t)
	airline := factory.ConstructAirline()
	airlineId := "123"
	airlineMockRepository.EXPECT().GetAirlineById(airlineId).Return(&airline, nil)
	airlineContext.Request, _ = http.NewRequest(http.MethodGet, AIRLINE_BY_ID, nil)
	airlineContext.AddParam("id", airlineId)

	airlineController.HandleGetAirlineById(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var airlineFromResponse models.Airline
	json.Unmarshal([]byte(responseBody), &airlineFromResponse)

	assert.Equal(t, airlineFromResponse, airline)
}

func TestHandleGetAirlineByIdWhenRecordDoesntExist(t *testing.T) {
	beforeEachAirlineTest(t)
	nonExistentAirlineId := "-23243"
	airlineMockRepository.EXPECT().GetAirlineById(nonExistentAirlineId).Return(nil, errors.New("foo bar"))
	airlineContext.Request, _ = http.NewRequest("GET", AIRLINE_BY_ID, nil)
	airlineContext.AddParam("id", nonExistentAirlineId)

	airlineController.HandleGetAirlineById(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)

	assert.Equal(t, fmt.Sprintf("{\"Error\":\"Incorrect airline id: %s\"}", nonExistentAirlineId), string(responseBody))
}

// TODO: All tests beyond this line have to be reviewed
func TestHandleCreateNewAirline(t *testing.T) {
	beforeEachAirlineTest(t)
	airline := factory.ConstructAirline()
	airlineName := "XYZAirline"
	airline = airline.SetName(airlineName)
	airlineMockRepository.EXPECT().CreateNewAirline(&airline).Return(nil)
	reqBody := fmt.Sprintf("{\"name\":\"%s\"}", airlineName)
	var airlineFromResponse models.Airline
	// TODO: How is this getting it from the response?
	json.Unmarshal([]byte(reqBody), &airlineFromResponse)
	airlineContext.Request, _ = http.NewRequest(http.MethodPost, POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	assert.Equal(t, airline.Name, airlineFromResponse.Name)
}

func TestHandleCreateNewAirlineWhenTheRequestPayloadIsEmpty(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{}`
	airlineContext.Request, _ = http.NewRequest(http.MethodPost, POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	// TODO: More assertions needed?
}

func TestHandleCreateNewAirlineWhenTheMandatoryValueIsAbsent(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"Name":""}`
	airlineContext.Request, _ = http.NewRequest(http.MethodPost, POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	// TODO: More assertions needed?
}

func TestHandleCreateNewAirlineWhenTheMandatoryKeyIsAbsent(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"Count":2}`
	airlineContext.Request, _ = http.NewRequest(http.MethodPost, POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	// TODO: More assertions needed?
}

func TestHandleCreateNewAirlineWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"name":123}`
	airlineContext.Request, _ = http.NewRequest(http.MethodPost, POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	// TODO: More assertions needed?
}

func TestHandleCreateNewAirlineWhereErrorIsThrownInRepositoryLayer(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"name":"Test"}`
	airlineMockRepository.EXPECT().CreateNewAirline(gomock.Any()).Return(errors.New("invalid request"))
	airlineContext.Request, _ = http.NewRequest(http.MethodPost, POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	// TODO: More assertions needed?
}

func TestHandleDeleteAirlineById(t *testing.T) {
	beforeEachAirlineTest(t)
	airlineMockRepository.EXPECT().DeleteAirlineById(gomock.Any()).Return(nil)
	airlineContext.Request, _ = http.NewRequest(http.MethodDelete, AIRLINE_BY_ID, nil)

	airlineController.HandleDeleteAirlineById(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
	// TODO: More assertions needed?
}

func TestHandleDeleteNewAirlineWhereErrorIsThrownInRepositoryLayer(t *testing.T) {
	beforeEachAirlineTest(t)
	airlineMockRepository.EXPECT().DeleteAirlineById(gomock.Any()).Return(errors.New("invalid request"))
	airlineContext.Request, _ = http.NewRequest(http.MethodDelete, AIRLINE_BY_ID, nil)

	airlineController.HandleDeleteAirlineById(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	// TODO: More assertions needed?
}
