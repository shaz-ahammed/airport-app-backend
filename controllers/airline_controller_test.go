package controllers

import (
	"airport-app-backend/mocks"
	"airport-app-backend/models"
	"net/http"
	"net/http/httptest"
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
	mockService.EXPECT().GetAirline(gomock.Any(), gomock.Any(), ctx).Return(mockAirlines, nil)
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
	mockService.EXPECT().GetAirlineById(gomock.Any(), ctx, gomock.Any()).Return(&mockAirlines, nil)
	ctx.Request, _ = http.NewRequest("GET", "airline/12332", nil)
	controllerRepo.HandleGetAirlineById(ctx)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())

}
