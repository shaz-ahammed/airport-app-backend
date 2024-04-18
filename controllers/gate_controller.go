package controllers

import (
	"context"
	"go.opencensus.io/trace"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"

	"airport-app-backend/middleware"
	"airport-app-backend/services"
)

type GateControllerRepository struct {
	service services.IGateRepository
}

func NewGateRepository(service services.IGateRepository) *GateControllerRepository {
	return &GateControllerRepository{
		service: service,
	}
}

func (gcr *GateControllerRepository) HandleGetGates(ctx *gin.Context) {
	c, span := trace.StartSpan(context.Background(), "handle_get_gates")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)
	log.Debug().Msg("Getting list of gates")
	pageStr := ctx.Query("page")
	floorStr := ctx.Query("floor")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	floor, err := strconv.Atoi(floorStr)
	if err != nil || floor < 0 {
		floor = -1
	}
	gates, err := gcr.service.GetGates(page, floor, c, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gates"})
		return
	}
	ctx.JSON(http.StatusOK, gates)
}

func (gcr *GateControllerRepository) HandleGetGateByID(ctx *gin.Context) {
	log.Debug().Msg("controller layer for retrieving gate details by id")
	c, span := trace.StartSpan(context.Background(), "handle_get_gate_by_id")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)

	gateID := ctx.Param("id")
	gate, err := gcr.service.GetGateByID(c, ctx, gateID)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 22P02") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Gate not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gate"})
		return
	}
	ctx.JSON(http.StatusOK, gate)

}
