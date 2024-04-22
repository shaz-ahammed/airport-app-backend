package controllers

import (
	"airport-app-backend/models"
	"airport-app-backend/services"
	"github.com/gin-gonic/gin/binding"
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

func (acr *AirlineControllerRepository) HandleGetAirlines(ctx *gin.Context) {
	log.Debug().Msg("Getting application health information")

	page, _ := strconv.Atoi(ctx.Query("page"))
	if page < 0 {
		ctx.JSON(400, gin.H{"msg": "Page number must be greater than 0"})
		return
	}
	airline, err := acr.service.GetAirline(page)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Airlines Details Not found"})
	}
	ctx.JSON(http.StatusOK, airline)
}

func (acr *AirlineControllerRepository) HandleGetAirlineById(ctx *gin.Context) {
	airlineId := ctx.Param(`id`)
	airline, err := acr.service.GetAirlineById(airlineId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "error: Incorrect Airlines Id")
	}
	ctx.JSON(http.StatusOK, airline)
}

func (acr *AirlineControllerRepository) HandleCreateNewAirline(ctx *gin.Context) {
	var airline models.Airline

	err := ctx.ShouldBindWith(&airline, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	serviceError := acr.service.CreateNewAirline(&airline)
	if serviceError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, "Created Successfully")
}

func (acr *AirlineControllerRepository) HandleCreateNewAirline(ctx *gin.Context) {
	var airline models.Airline

	err := ctx.ShouldBindWith(&airline, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	serviceError := acr.service.CreateNewAirline(&airline)
	if serviceError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": serviceError.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, "Created Successfully")
}
