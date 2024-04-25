package controllers

import (
	"airport-app-backend/models"
	"airport-app-backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin/binding"

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

// @Summary			Get all airlines
// @Router 			/airlines [get]
// @Description 	Gets all the airlines
// @ID 				get-all-airlines
// @Tags 			airline
// @Produce  		json
// @Param   		page	query	int		false	"Page number (default = 0)"
// @Success 		200		"ok"
// @Failure 		500		"Internal server error"
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

// @Summary			Get airline by ID
// @Router			/airline/{id} [get]
// @Description 	Gets airline by ID
// @ID 				get-airline-by-id
// @Tags 			airline
// @Produce  		json
// @Param   		id		path		string		true		"Airline ID"
// @Success 		200		"ok"
// @Failure 		500		"Internal server error"
func (acr *AirlineControllerRepository) HandleGetAirlineById(ctx *gin.Context) {
	airlineId := ctx.Param("id")
	airline, err := acr.service.GetAirlineById(airlineId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "error: Incorrect Airlines Id")
		return
	}
	ctx.JSON(http.StatusOK, airline)
}

// @Summary			Create new airline
// @Router			/airline [post]
// @Description 	Create new airline
// @ID 				create-airline
// @Tags 			airline
// @Produce  		json
// @Param   		airline		body		models.Airline		true		"Airline Object"
// @Success 		200		"ok"
// @Failure 		500		"Internal server error"
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
	ctx.JSON(http.StatusCreated, "Created a new airline Successfully")
}

// @Summary Delete airline by Id
// @Router /airline/{id} [delete]
// @Summary Delete airline by Id
// @Description Delete the airline details of the particular id
// @ID delete-airline-by-id
// @Tags airline
// @Param id path string true "Airline ID"
// @Success 200  "ok"
// @Failure 404 "Incorrect airline id"
func (acr *AirlineControllerRepository) HandleDeleteAirlineById(ctx *gin.Context) {
	airlineId := ctx.Param(`id`)
	err := acr.service.DeleteAirlineById(airlineId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error: Incorrect Airlines Id")
		return
	}
	ctx.JSON(http.StatusOK, "Deleted the airline successfully")
}
