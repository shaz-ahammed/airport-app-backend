package controllers

import (
	"airport-app-backend/middleware"
	"airport-app-backend/services"
	"context"
	"go.opencensus.io/trace"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AirlineControllerRepository struct {
	service services.IAirlineRepository
}

func NewAirlineControllerRepository(service services.IAirlineRepository) *AirlineControllerRepository {
	return &AirlineControllerRepository{
		service: service,
	}
}

func (hcr *AirlineControllerRepository) HandleAirline(ctx *gin.Context) {
	c, span := trace.StartSpan(context.Background(), "handle_get_airline")
	defer span.End()
	middleware.TraceSpanTags(span)(ctx)
	log.Debug().Msg("Getting application health information")
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page < 0 {
		ctx.JSON(400, gin.H{"msg": "Page number must be greater than 0"})
		return
	}
	appAirline, _ := hcr.service.GetAirline(page, c, ctx)
	ctx.JSON(http.StatusOK, appAirline)
}
