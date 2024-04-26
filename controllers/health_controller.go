package controllers

import (
	"airport-app-backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthControllerRepository struct {
	repository repositories.IHealthRepository
}

func NewControllerRepository(repository repositories.IHealthRepository) *HealthControllerRepository {
	return &HealthControllerRepository{
		repository: repository,
	}
}

func (repo *HealthControllerRepository) HandleHealth(ctx *gin.Context) {
	appHealth := repo.repository.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}

func (hcr *HealthControllerRepository) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home page : AIRPORT MANAGEMENT SYSTEM")
}
