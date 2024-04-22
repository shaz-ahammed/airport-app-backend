package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandleAirlineController(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	mockAirline := make([]models.Airline, 3)
	mockAirline = append(mockAirline, models.Airline{Name: "Kingfisher"})
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService.EXPECT().GetAirline(gomock.Any()).Return(mockAirline, nil)
	ctx.Request, _ = http.NewRequest("GET", "/airline", nil)
	controllerRepo.HandleGetAirline(ctx)

	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestHandleAirlineByIdController(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockAirline := models.Airline{Name: "Jet Airways"}
	mockService.EXPECT().GetAirlineById(gomock.Any()).Return(&mockAirline, nil)
	ctx.Request, _ = http.NewRequest("GET", "airline/12332", nil)
	controllerRepo.HandleGetAirlineById(ctx)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestHandleCreateNewAirline(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	airline := models.Airline{
		Name: "XYZAirline",
	}
	mockService.EXPECT().CreateNewAirline(&airline).Return(nil)
	reqBody := `{"name":"XYZAirline"}`

	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusCreated, ctx.Writer.Status())

	var response models.Airline
	err := json.Unmarshal([]byte(reqBody), &response)
	assert.NoError(t, err)

	assert.Equal(t, airline.Name, response.Name)

}

func TestHandleCreateNewAirlineWhenTheMandatoryValueIsAbsent(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{"Name":""}`

	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())

}

func TestHandleCreateNewAirlineWhenThePayloadIsEmpty(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{}`

	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())

}

func TestHandleCreateNewAirlineWhenTheMandatoryKeyIsAbsent(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{"Count":2}`
	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())

}

func TestHandleCreateNewAirlineWhenDataOfDifferentDatatypeIsGiven(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	reqBody := `{"name":123}`

	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())

}

func TestHandleCreateNewAirlineWhereErrorIsThrownInServiceLayer(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	airline := models.Airline{
		Name: "Test",
	}
	reqBody := `{"name":"Test"}`
	mockService.EXPECT().CreateNewAirline(&airline).Return(errors.New("invalid Request"))
	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusBadRequest, ctx.Writer.Status())
}
