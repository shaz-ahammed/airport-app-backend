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
	airlineContext.Request, _ = http.NewRequest("GET", GET_ALL_AIRLINES, nil)

	airlineController.HandleGetAllAirlines(airlineContext)

	response := responseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var responseAirlines []models.Airline
	json.Unmarshal([]byte(responseBody), &responseAirlines)

	assert.Equal(t, 3, len(responseAirlines))
	assert.Contains(t, responseAirlines, airline1)
	assert.Contains(t, responseAirlines, airline2)
	assert.Contains(t, responseAirlines, airline3)
}

// TODO: InternalServerError scenario for GetAllAirlines

func TestHandleAirlineById(t *testing.T) {
	beforeEachAirlineTest(t)
	newAirline := factory.ConstructAirline()
	airlines := newAirline.SetName("Jet Airways")
	airlineMockRepository.EXPECT().GetAirlineById(gomock.Any()).Return(&airlines, nil)
	airlineContext.Request, _ = http.NewRequest("GET", AIRLINE_BY_ID, nil)

	airlineController.HandleGetAirlineById(airlineContext)

	assert.Equal(t, http.StatusOK, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirline(t *testing.T) {
	beforeEachAirlineTest(t)
	airline := factory.ConstructAirline()
	airlineName := "XYZAirline"
	airline = airline.SetName(airlineName)
	airlineMockRepository.EXPECT().CreateNewAirline(&airline).Return(nil)
	reqBody := fmt.Sprintf("{\"name\":\"%s\"}", airlineName)
	var response models.Airline
	err := json.Unmarshal([]byte(reqBody), &response)
	airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	assert.Equal(t, http.StatusCreated, airlineContext.Writer.Status())
	assert.NoError(t, err)
	assert.Equal(t, airline.Name, response.Name)
}

func TestHandleCreateNewAirlineWhenTheMandatoryValueIsAbsent(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"Name":""}`
	airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhenTheRequestPayloadIsEmpty(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{}`
	airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhenTheMandatoryKeyIsAbsent(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"Count":2}`
	airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"name":123}`
	airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhereErrorIsThrownInRepositoryLayer(t *testing.T) {
	beforeEachAirlineTest(t)
	reqBody := `{"name":"Test"}`
	airlineMockRepository.EXPECT().CreateNewAirline(gomock.Any()).Return(errors.New("invalid Request"))
	airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))

	airlineController.HandleCreateNewAirline(airlineContext)

	assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleDeleteAirlineById(t *testing.T) {
	beforeEachAirlineTest(t)
	airlineMockRepository.EXPECT().DeleteAirlineById(gomock.Any()).Return(nil)
	airlineContext.Request, _ = http.NewRequest("DELETE", AIRLINE_BY_ID, nil)

	airlineController.HandleDeleteAirlineById(airlineContext)

	assert.Equal(t, http.StatusOK, airlineContext.Writer.Status())
}

func TestHandleDeleteNewAirlineWhereErrorIsThrownInRepositoryLayer(t *testing.T) {
	beforeEachAirlineTest(t)
	airlineMockRepository.EXPECT().DeleteAirlineById(gomock.Any()).Return(errors.New("invalid Request"))
	airlineContext.Request, _ = http.NewRequest("DELETE", AIRLINE_BY_ID, nil)

	airlineController.HandleDeleteAirlineById(airlineContext)

	assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}
