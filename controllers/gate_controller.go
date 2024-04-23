package controllers

import (
	"airport-app-backend/models"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"strings"

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

func (gcr *GateControllerRepository) HandleGetGateById(ctx *gin.Context) {
	log.Debug().Msg("controller layer for retrieving gate details by id")

	gateID := ctx.Param("id")
	gate, err := gcr.service.GetGateById(gateID)
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

func (gcr *GateControllerRepository) HandleCreateNewGate(ctx *gin.Context) {
	var gate models.Gate

	err := ctx.ShouldBindWith(&gate, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceError := gcr.service.CreateNewGate(&gate)
	if serviceError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Successfully created a gate")
}
