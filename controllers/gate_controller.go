package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

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
	gates, err := gcr.service.GetGates(page, floor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gates"})
		return
	}
	ctx.JSON(http.StatusOK, gates)
}