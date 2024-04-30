package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"airport-app-backend/models/factory"

	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var AIRCRAFTS = "/aircrafts"
var AIRCRAFT = "/aircraft"

var mockAircraftRepository *mocks.MockIAircraftRepository
var aircraftController *AircraftController
var aircraftContext *gin.Context
var aircraftResponseRecorder *httptest.ResponseRecorder

func beforeEachAircraftTest(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	mockAircraftRepository = mocks.NewMockIAircraftRepository(gomockController)
	aircraftController = NewAircraftController(mockAircraftRepository)
	aircraftResponseRecorder = httptest.NewRecorder()
	aircraftContext, _ = gin.CreateTestContext(aircraftResponseRecorder)
}

func TestHandleGetAllAircrafts(t *testing.T) {
	beforeEachAircraftTest(t)
	var aircrafts []models.Aircraft
	aircraft1 := factory.ConstructAircraft()
	aircrafts = append(aircrafts, aircraft1)
	aircraft2 := factory.ConstructAircraft()
	aircrafts = append(aircrafts, aircraft2)
	aircraft3 := factory.ConstructAircraft()
	aircrafts = append(aircrafts, aircraft3)

	mockAircraftRepository.EXPECT().RetrieveAllAircrafts(gomock.Any(), gomock.Any(), gomock.Any()).Return(aircrafts, nil)
	aircraftContext.Request, _ = http.NewRequest(http.MethodGet, AIRCRAFTS, nil)

	aircraftController.HandleGetAllAircrafts(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var aircraftsFromResponse []models.Aircraft
	json.Unmarshal([]byte(responseBody), &aircraftsFromResponse)

	assert.Equal(t, 3, len(aircraftsFromResponse))
	assert.Contains(t, aircraftsFromResponse, aircraft1)
	assert.Contains(t, aircraftsFromResponse, aircraft2)
	assert.Contains(t, aircraftsFromResponse, aircraft3)
}

func TestHandleGetAllAircraftsWhenServiceReturnsError(t *testing.T) {
	beforeEachAircraftTest(t)
	mockAircraftRepository.EXPECT().RetrieveAllAircrafts(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("Invalid"))
	aircraftContext.Request, _ = http.NewRequest(http.MethodGet, AIRCRAFTS, nil)

	aircraftController.HandleGetAllAircrafts(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, "{\"Error\":\"Internal server error\"}", string(responseBody))
}
