package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"airport-app-backend/models/factory"
	"encoding/json"
	"errors"
	"fmt"
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

func beforeEachAirlineTest(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	airlineMockRepository = mocks.NewMockIAirlineRepository(gomockController)
	airlineController = NewAirlineControllerRepository(airlineMockRepository)
	airlineContext, _ = gin.CreateTestContext(httptest.NewRecorder())
}

func TestHandleAirline(t *testing.T) {
	beforeEachAirlineTest(t)
	mockAirline := make([]models.Airline, 3)
	newAirline := factory.ConstructAirline()
	mockAirline = append(mockAirline, newAirline.SetName("Kingfisher"))

	airlineMockRepository.EXPECT().GetAirline(gomock.Any()).Return(mockAirline, nil)
	airlineContext.Request, _ = http.NewRequest("GET", GET_ALL_AIRLINES, nil)
	airlineController.HandleGetAirlines(airlineContext)

	assert.Equal(t, http.StatusOK, airlineContext.Writer.Status())
}

func TestHandleAirlineById(t *testing.T) {
	beforeEachAirlineTest(t)
	newAirline := factory.ConstructAirline()
	mockAirline := newAirline.SetName("Jet Airways")
	airlineMockRepository.EXPECT().GetAirlineById(gomock.Any()).Return(&mockAirline, nil)
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
