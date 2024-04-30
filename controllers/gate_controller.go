package controllers

import (
	"airport-app-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"airport-app-backend/repositories"
)

type GateController struct {
	repository repositories.IGateRepository
}

func NewGateController(repository repositories.IGateRepository) *GateController {
	return &GateController{
		repository: repository,
	}
}

// @Summary Get all gates
// @Router /gates [get]
// @Description get all the gate details
// @ID get-all-gate
// @Tags gate
// @Produce  json
// @Param   page        query    int     false        "Page number (default = 0)"
// @Param   floor       query    int     false        "filter by floor (default = all floor)"
// @Success 200  "ok"
// @Failure 500 "Internal server error"
func (gc *GateController) HandleGetAllGates(ctx *gin.Context) {
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
	gates, err := gc.repository.GetAllGates(page, floor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gates"})
		return
	}
	ctx.JSON(http.StatusOK, gates)
}

// @Router /gate/{id} [get]
// @Summary Get gate by Id
// @Description Retrieve a gate by its Id
// @ID get-gate-by-id
// @Tags gate
// @Produce  json
// @Param id path string true "Gate Id"
// @Success 200  "ok"
// @Failure 400  "Gate not found"
func (gc *GateController) HandleGetGate(ctx *gin.Context) {
	log.Debug().Msg("controller layer for retrieving gate details by id")

	gateId := ctx.Param("id")
	gate, err := gc.repository.GetGate(gateId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Incorrect gate id: " + gateId})
		return
	}
	ctx.JSON(http.StatusOK, gate)
}

// @Router /gate [POST]
// @Summary Create gate
// @Description Create a new gate
// @ID create-gate
// @Tags gate
// @Accept json
// @Produce  json
// @Param gate body models.Gate true "New gate object"
// @Success 200  "ok"
// @Failure 400  "Bad request"
func (gc *GateController) HandleCreateNewGate(ctx *gin.Context) {
	var gate models.Gate

	err := ctx.ShouldBindWith(&gate, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repositoryError := gc.repository.CreateNewGate(&gate)
	if repositoryError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": repositoryError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Successfully created a gate")
}

// @Router /gate/{id} [PUT]
// @Summary Update gate
// @Description Update gate of given id
// @ID update-gate
// @Tags gate
// @Accept json
// @Produce  json
// @Param gate body models.Gate true "Updated gate object"
// @Param id path string true "Gate Id"
// @Success 200  "ok"
// @Failure 400  "Gate not found"
func (gc *GateController) HandleUpdateGate(ctx *gin.Context) {
	log.Debug().Msg("controller layer for updating gate info")

	var gate models.Gate
	gateId := ctx.Param("id")
	err := ctx.ShouldBindWith(&gate, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = gc.repository.UpdateGate(gateId, gate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Successfully updated gate details")
}
