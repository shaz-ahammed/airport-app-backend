package controllers

import (
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

// @Router /gates [get]
// @Summary Get all gates
// @Description get all the gate details
// @ID get-all-gate
// @Tags gate
// @Produce  json
// @Param   page        query    int     false        "Page number (default = 0)"
// @Param   floor       query    int     false        "filter by floor (default = all floor)"
// @Success 200  "ok"
// @Failure 500 "Internal server error"
func (gcr *GateControllerRepository) HandleGetGates(ctx *gin.Context) {
	log.Debug().Msg("Getting list of gates")

	pageStr := ctx.Query("page")
	floorStr := ctx.Query("floor")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 0 {
		page = 0
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

// @Router /gate/{id} [get]
// @Summary Get gate by ID
// @Description Retrieve a gate by its ID
// @ID get-gate-by-id
// @Tags gate
// @Produce  json
// @Param id path string true "Gate ID"
// @Success 200  "ok"
// @Failure 400  "Gate not found"
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
