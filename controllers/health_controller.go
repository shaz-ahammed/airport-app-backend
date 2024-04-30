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

func (hc *HealthController) HandleHealth(ctx *gin.Context) {
	appHealth := hc.repository.GetAppHealth()
	ctx.JSON(http.StatusOK, appHealth)
}

func (hc *HealthController) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home page : AIRPORT MANAGEMENT SYSTEM")
}
