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
var GET_ALL_AIRLINES = "/airlines"
var GET_AIRLINE_BY_ID = "/airline/123"
var POST_AIRLINE = "/airline"

var airlineMockService *mocks.MockIAirlineRepository
var airlineMockController *AirlineControllerRepository
var airlineContext *gin.Context

func beforeEachAirlineTest(t *testing.T) {
  mockControl := gomock.NewController(t)
  defer mockControl.Finish()

  airlineMockService = mocks.NewMockIAirlineRepository(mockControl)
  airlineMockController = NewAirlineControllerRepository(airlineMockService)
  recorder := httptest.NewRecorder()
  airlineContext, _ = gin.CreateTestContext(recorder)
}

func TestHandleAirline(t *testing.T) {
  beforeEachAirlineTest(t)
  mockAirline := make([]models.Airline, 3)
  mockAirline = append(mockAirline, models.Airline{Name: "Kingfisher"})
  
  airlineMockService.EXPECT().GetAirline(gomock.Any()).Return(mockAirline, nil)
  airlineContext.Request, _ = http.NewRequest("GET", GET_ALL_AIRLINES, nil)
  airlineMockController.HandleGetAirlines(airlineContext)

  assert.Equal(t, http.StatusOK, airlineContext.Writer.Status())
}

func TestHandleAirlineById(t *testing.T) {
  beforeEachAirlineTest(t)
  mockAirline := models.Airline{Name: "Jet Airways"}

  airlineMockService.EXPECT().GetAirlineById(gomock.Any()).Return(&mockAirline, nil)
  airlineContext.Request, _ = http.NewRequest("GET", GET_AIRLINE_BY_ID, nil)
  airlineMockController.HandleGetAirlineById(airlineContext)

  assert.Equal(t, http.StatusOK, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirline(t *testing.T) {
  beforeEachAirlineTest(t)
  airline := models.Airline{
    Name: "XYZAirline",
  }
  airlineMockService.EXPECT().CreateNewAirline(&airline).Return(nil)
  reqBody := `{"name":"XYZAirline"}`
  airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))
  airlineMockController.HandleCreateNewAirline(airlineContext)

  var response models.Airline
  err := json.Unmarshal([]byte(reqBody), &response)

  assert.Equal(t, http.StatusCreated, airlineContext.Writer.Status())
  assert.NoError(t, err)
  assert.Equal(t, airline.Name, response.Name)
}

func TestHandleCreateNewAirlineWhenTheMandatoryValueIsAbsent(t *testing.T) {
  beforeEachAirlineTest(t)
  reqBody := `{"Name":""}`

  airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))
  airlineMockController.HandleCreateNewAirline(airlineContext)

  assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhenTheRequestPayloadIsEmpty(t *testing.T) {
  beforeEachAirlineTest(t)
  reqBody := `{}`

  airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))
  airlineMockController.HandleCreateNewAirline(airlineContext)

  assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhenTheMandatoryKeyIsAbsent(t *testing.T) {
  beforeEachAirlineTest(t)
  reqBody := `{"Count":2}`

  airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))
  airlineMockController.HandleCreateNewAirline(airlineContext)

  assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
  beforeEachAirlineTest(t)
  reqBody := `{"name":123}`

  airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))
  airlineMockController.HandleCreateNewAirline(airlineContext)

  assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}

func TestHandleCreateNewAirlineWhereErrorIsThrownInServiceLayer(t *testing.T) {
  beforeEachAirlineTest(t)
  airline := models.Airline{
    Name: "Test",
  }
  reqBody := `{"name":"Test"}`

  airlineMockService.EXPECT().CreateNewAirline(&airline).Return(errors.New("invalid Request"))
  airlineContext.Request, _ = http.NewRequest("POST", POST_AIRLINE, strings.NewReader(reqBody))
  airlineMockController.HandleCreateNewAirline(airlineContext)

  assert.Equal(t, http.StatusBadRequest, airlineContext.Writer.Status())
}
