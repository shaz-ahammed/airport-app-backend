package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"encoding/json"
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
	mockAirlines := make([]models.Airlines, 3)
	mockAirlines = append(mockAirlines, models.Airlines{Name: "Kingfisher"})
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	mockService.EXPECT().GetAirline(gomock.Any()).Return(mockAirlines, nil)
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
	mockAirlines := models.Airlines{Name: "Jet Airways"}
	mockService.EXPECT().GetAirlineById(gomock.Any()).Return(&mockAirlines, nil)
	ctx.Request, _ = http.NewRequest("GET", "airline/12332", nil)
	controllerRepo.HandleGetAirlineById(ctx)
	fmt.Println(ctx.Writer.Status())
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

func TestHandleCreateNewAirline(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockIAirlineRepository(mockCtrl)
	controllerRepo := NewAirlineControllerRepository(mockService)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	airline := models.Airlines{
		Name: "XYZAirlines",
	}
	mockService.EXPECT().CreateNewAirline(gomock.Any(), ctx, &airline).Return(nil)
	reqBody := `{"name":"XYZAirlines"}`

	ctx.Request, _ = http.NewRequest("POST", "/airline", strings.NewReader(reqBody))
	controllerRepo.HandleCreateNewAirline(ctx)
	assert.Equal(t, http.StatusCreated, ctx.Writer.Status())

	var response models.Airlines
	err := json.Unmarshal([]byte(reqBody), &response)
	assert.NoError(t, err)

	assert.Equal(t, airline.Name, response.Name)

}
