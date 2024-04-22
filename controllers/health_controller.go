package controllers

import (
	"airport-app-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthControllerRepository struct {
	service services.IHealthRepository
}

func NewControllerRepository(service services.IHealthRepository) *HealthControllerRepository {
	return &HealthControllerRepository{
		service: service,
	}
}

func (repo *HealthControllerRepository) HandleHealth(ctx *gin.Context) {
	appHealth := repo.service.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}

func (hcr *HealthControllerRepository) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home page : AIRPORT MANAGEMENT SYSTEM")
}
