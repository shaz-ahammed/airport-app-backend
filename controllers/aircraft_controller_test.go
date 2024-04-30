package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"airport-app-backend/models/factory"
	"fmt"
	"strings"

	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
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

func TestHandleGetAircraft(t *testing.T) {
	beforeEachAircraftTest(t)
	aircraft := factory.ConstructAircraft()
	aircraftId := "123"
	mockAircraftRepository.EXPECT().RetrieveAircraft(aircraftId).Return(&aircraft, nil)
	aircraftContext.Request, _ = http.NewRequest(http.MethodGet, AIRCRAFT, nil)
	aircraftContext.AddParam("id", aircraftId)

	aircraftController.HandleGetAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var aircraftFromResponse models.Aircraft
	json.Unmarshal([]byte(responseBody), &aircraftFromResponse)

	assert.Equal(t, aircraft, aircraftFromResponse)
}

func TestHandleGetAircraftWhenRecordDoesntExist(t *testing.T) {
	beforeEachAircraftTest(t)
	nonExistentAircraftId := "-23243"
	mockAircraftRepository.EXPECT().RetrieveAircraft(nonExistentAircraftId).Return(nil, errors.New("foo bar"))
	aircraftContext.Request, _ = http.NewRequest("GET", AIRCRAFT, nil)
	aircraftContext.AddParam("id", nonExistentAircraftId)

	aircraftController.HandleGetAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, fmt.Sprintf("{\"Error\":\"Incorrect aircraft id: %s\"}", nonExistentAircraftId), string(responseBody))
}

func TestHandleCreateNewAircraft(t *testing.T) {
	beforeEachAircraftTest(t)
	airlineId := faker.UUIDHyphenated()
	aircraft := factory.ConstructAircraft()
	reqBody, _ := json.Marshal(&aircraft)
	aircraftContext.Request, _ = http.NewRequest(http.MethodPost, AIRCRAFT, strings.NewReader(string(reqBody)))
	mockAircraftRepository.EXPECT().InsertAircraft(aircraft, airlineId).Return(nil)
	aircraftContext.AddParam("airline_id", airlineId)

	aircraftController.HandleCreateNewAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestHandleCreateNewAircraftWhenTheRequestPayloadIsEmpty(t *testing.T) {
	beforeEachAircraftTest(t)
	reqBody := `{}`
	aircraftContext.Request, _ = http.NewRequest(http.MethodPost, AIRCRAFT, strings.NewReader(reqBody))

	aircraftController.HandleCreateNewAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, `{"error":"Key: 'Aircraft.TailNumber' Error:Field validation for 'TailNumber' failed on the 'required' tag\nKey: 'Aircraft.Capacity' Error:Field validation for 'Capacity' failed on the 'required' tag"}`,
		string(responseBody))
}

func TestHandleCreateNewAircraftWhenTheMandatoryValueIsNull(t *testing.T) {
	beforeEachAircraftTest(t)
	aircraft := factory.ConstructAircraft()
	aircraft.SetTailNumber("")
	reqBody, _ := json.Marshal(&aircraft)
	aircraftContext.Request, _ = http.NewRequest(http.MethodPost, AIRCRAFT, strings.NewReader(string(reqBody)))

	aircraftController.HandleCreateNewAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, "{\"error\":\"Key: 'Aircraft.TailNumber' Error:Field validation for 'TailNumber' failed on the 'required' tag\"}", string(responseBody))
}

func TestHandleCreateNewAircraftWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {
	beforeEachAircraftTest(t)
	reqBody := `{"tail_number":123, "capacity":"eleven"}`
	aircraftContext.Request, _ = http.NewRequest(http.MethodPost, AIRCRAFT, strings.NewReader(reqBody))

	aircraftController.HandleCreateNewAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, "{\"error\":\"json: cannot unmarshal number into Go struct field Aircraft.tail_number of type string\"}", string(responseBody))
}

func TestHandleCreateNewAircraftWhereErrorIsThrownInRepositoryLayer(t *testing.T) {
	beforeEachAircraftTest(t)
	airlineId := faker.UUIDHyphenated()
	aircraft := factory.ConstructAircraft()
	reqBody, _ := json.Marshal(&aircraft)
	aircraftContext.Request, _ = http.NewRequest(http.MethodPost, AIRCRAFT, strings.NewReader(string(reqBody)))
	mockAircraftRepository.EXPECT().InsertAircraft(aircraft, airlineId).Return(errors.New("invalid request"))
	aircraftContext.AddParam("airline_id", airlineId)

	aircraftController.HandleCreateNewAircraft(aircraftContext)

	response := aircraftResponseRecorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, "{\"error\":\"invalid request\"}", string(responseBody))
}
