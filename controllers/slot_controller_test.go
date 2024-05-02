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

var SLOTS = "/slots"

var mockSlotRepository *mocks.MockISlotRepository
var slotController *SlotController
var slotContext *gin.Context
var slotResponseRecorder *httptest.ResponseRecorder

func beforeEachslotTest(t *testing.T) {
	gomockController := gomock.NewController(t)
	defer gomockController.Finish()

	mockSlotRepository = mocks.NewMockISlotRepository(gomockController)
	slotController = NewSlotController(mockSlotRepository)
	slotResponseRecorder = httptest.NewRecorder()
	slotContext, _ = gin.CreateTestContext(slotResponseRecorder)
}

func TestHandleGetAllSlots(t *testing.T) {
	beforeEachslotTest(t)
	var slots []models.Slot
	slot1 := factory.ConstructSlot()
	slots = append(slots, slot1)
	slot2 := factory.ConstructSlot()
	slots = append(slots, slot2)
	slot3 := factory.ConstructSlot()
	slots = append(slots, slot3)

	mockSlotRepository.EXPECT().RetrieveAllSlots(gomock.Any(), true).Return(slots, nil)
	slotContext.Request, _ = http.NewRequest(http.MethodGet, SLOTS, nil)

	slotController.HandleGetAllSlots(slotContext)

	response := slotResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var slotsFromResponse []models.Slot
	json.Unmarshal([]byte(responseBody), &slotsFromResponse)

	assert.Equal(t, 3, len(slotsFromResponse))
	assert.Contains(t, slotsFromResponse, slot1)
	assert.Contains(t, slotsFromResponse, slot2)
	assert.Contains(t, slotsFromResponse, slot3)
}

func TestHandleGetAllSlotsWhenServiceReturnsError(t *testing.T) {
	beforeEachslotTest(t)
	mockSlotRepository.EXPECT().RetrieveAllSlots(gomock.Any(), gomock.Any()).Return(nil, errors.New("Invalid"))
	slotContext.Request, _ = http.NewRequest(http.MethodGet, SLOTS, nil)

	slotController.HandleGetAllSlots(slotContext)

	response := slotResponseRecorder.Result()
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, "{\"Error\":\"Internal server error\"}", string(responseBody))
}

func TestHandleGetAllSlotsWhenIsAvailableIsFalse(t *testing.T) {
	beforeEachslotTest(t)
	var slots []models.Slot
	slot1 := factory.ConstructSlot()
	slot1.SetStatus("Booked")
	slots = append(slots, slot1)
	slot2 := factory.ConstructSlot()
	slot2.SetStatus("Reserved")
	slots = append(slots, slot2)
	slot3 := factory.ConstructSlot()
	slot3.SetStatus("Booked")
	slots = append(slots, slot3)

	mockSlotRepository.EXPECT().RetrieveAllSlots(gomock.Any(), false).Return(slots, nil)
	slotContext.Request, _ = http.NewRequest(http.MethodGet, SLOTS+"?is_available=false", nil)

	slotController.HandleGetAllSlots(slotContext)

	response := slotResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var slotsFromResponse []models.Slot
	json.Unmarshal([]byte(responseBody), &slotsFromResponse)

	assert.Equal(t, 3, len(slotsFromResponse))
	assert.Contains(t, slotsFromResponse, slot1)
	assert.Equal(t, slotsFromResponse[0].Status, "Booked")
	assert.Contains(t, slotsFromResponse, slot2)
	assert.Equal(t, slotsFromResponse[1].Status, "Reserved")
	assert.Contains(t, slotsFromResponse, slot3)
	assert.Equal(t, slotsFromResponse[2].Status, "Booked")
}

func TestHandleGetAllSlotsWhenIsAvailableIsFalseAndPageIsGiven(t *testing.T) {
	beforeEachslotTest(t)
	var slots []models.Slot
	slot1 := factory.ConstructSlot()
	slot1.SetStatus("Available")
	slots = append(slots, slot1)
	slot2 := factory.ConstructSlot()
	slot2.SetStatus("Available")
	slots = append(slots, slot2)

	mockSlotRepository.EXPECT().RetrieveAllSlots(1, true).Return(slots, nil)
	slotContext.Request, _ = http.NewRequest(http.MethodGet, SLOTS+"?page=1&is_available=true", nil)

	slotController.HandleGetAllSlots(slotContext)

	response := slotResponseRecorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	responseBody, _ := io.ReadAll(response.Body)
	var slotsFromResponse []models.Slot
	json.Unmarshal([]byte(responseBody), &slotsFromResponse)

	assert.Equal(t, 2, len(slotsFromResponse))
	assert.Contains(t, slotsFromResponse, slot1)
	assert.Equal(t, slotsFromResponse[0].Status, "Available")
	assert.Contains(t, slotsFromResponse, slot2)
	assert.Equal(t, slotsFromResponse[1].Status, "Available")
}
