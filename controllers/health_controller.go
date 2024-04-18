package controllers

import (
	"airport-app-backend/middleware"
	"airport-app-backend/services"
	"context"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
	"net/http"
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
	c, span := trace.StartSpan(context.Background(), "handle_get_health")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)
	appHealth := repo.service.GetAppHealth(c, ctx)
	ctx.JSON(http.StatusOK, appHealth)
}

func (hcr *HealthControllerRepository) Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home page : AIRPORT MANAGEMENT SYSTEM")
}
