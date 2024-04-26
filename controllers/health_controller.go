package controllers

import (
	"airport-app-backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthControllerRepository struct {
	service repositories.IHealthRepository
}

func NewControllerRepository(service repositories.IHealthRepository) *HealthControllerRepository {
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
