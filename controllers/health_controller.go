package controllers

import (
	"airport-app-backend/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	repository repositories.IHealthRepository
}

func NewController(repository repositories.IHealthRepository) *HealthController {
	return &HealthController{
		repository: repository,
	}
}

func (repo *HealthController) HandleHealth(ctx *gin.Context) {
	appHealth := repo.repository.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}

func (hcr *HealthController) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home page : AIRPORT MANAGEMENT SYSTEM")
}
