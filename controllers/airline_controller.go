package controllers

import (
	"airport-app-backend/middleware"
	"airport-app-backend/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
)

type AirlineControllerRepository struct {
	service services.IAirlineRepository
}

func NewAirlineControllerRepository(service services.IAirlineRepository) *AirlineControllerRepository {
	return &AirlineControllerRepository{
		service: service,
	}
}

func (acr *AirlineControllerRepository) HandleGetAirline(ctx *gin.Context) {
	c, span := trace.StartSpan(context.Background(), "handle_get_airline")
	defer span.End()
	middleware.TraceSpanTags(span)(ctx)
	log.Debug().Msg("Getting application health information")
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page < 0 {
		ctx.JSON(400, gin.H{"msg": "Page number must be greater than 0"})
		return
	}
	appAirline, err := acr.service.GetAirline(page,c,ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Airlines Details Not found"})
	}
	ctx.JSON(http.StatusOK, appAirline)
}

func (acr *AirlineControllerRepository) HandleGetAirlineById(ctx *gin.Context) {
	c, span := trace.StartSpan(context.Background(), "handle_airline_by_id")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)
	airline_id := ctx.Param(`id`)
	appAirline, err := acr.service.GetAirlineById(c, ctx, airline_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Incorrect Airlines Id"})
	}
	ctx.JSON(http.StatusOK, appAirline)
}
