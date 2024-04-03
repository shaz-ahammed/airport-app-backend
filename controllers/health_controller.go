package controllers

import (
	"airport-app-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
	log.Debug().Msg("Getting application health information")
	appHealth := repo.service.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}

func (hcr *HealthControllerRepository) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home page : AIRPORT MANAGEMENT SYSTEM")
}
