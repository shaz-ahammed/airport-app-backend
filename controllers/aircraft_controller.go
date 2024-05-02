package controllers

import (
	"airport-app-backend/models"
	"airport-app-backend/repositories"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	START_YEAR = 1970
	NO_FILTER  = -1
)

type AircraftController struct {
	repository repositories.IAircraftRepository
}

func NewAircraftController(repository repositories.IAircraftRepository) *AircraftController {
	return &AircraftController{
		repository: repository,
	}
}

// @Summary Get all aircrafts
// @Router /aircrafts [get]
// @Description get all the aircrafts
// @ID get-all-aircrafts
// @Tags aircraft
// @Produce  json
// @Param   page        query    int     false        "Page number (default = 1)"
// @Param   year       query    int     false        "filter by manufacturing (default = all year)"
// @Param   capacity       query    int     false        "condition by capacity grater than given value (default = 0)"
// @Success 200  "ok"
// @Failure 500 "Internal server error"
func (ac *AircraftController) HandleGetAllAircrafts(ctx *gin.Context) {
	// TODO: Convert to using a pagination library to handle this and other edge cases
	page, _ := strconv.Atoi(ctx.Query("page"))
	if page < 0 {
		ctx.JSON(400, gin.H{"msg": "Page number must be greater than 0"})
		return
	}
	year, _ := strconv.Atoi(ctx.Query("year"))
	if year < START_YEAR {
		year = NO_FILTER
	}
	capacity, _ := strconv.Atoi(ctx.Query("capacity"))
	if capacity < 0 {
		capacity = NO_FILTER
	}

	aircrafts, err := ac.repository.RetrieveAllAircrafts(page, capacity, year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, aircrafts)
}

// @Summary			Get aircraft by Id
// @Router			/aircraft/{id} [get]
// @Description 	Gets aircraft by Id
// @ID 				get-aircraft-by-id
// @Tags 			aircraft
// @Produce  		json
// @Param   		id		path		string		true		"Aircraft Id"
// @Success 		200		"ok"
// @Failure 		400		"Aircraft not found"
func (ac *AircraftController) HandleGetAircraft(ctx *gin.Context) {
	aircraftId := ctx.Param("id")
	aircraft, err := ac.repository.RetrieveAircraft(aircraftId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Incorrect aircraft id: " + aircraftId})
		return
	}
	ctx.JSON(http.StatusOK, aircraft)
}

// @Summary Create aircraft
// @Router /airlines/{airline_id}/aircraft [post]
// @Description Create a new aircraft for the given airline
// @ID create-aircraft
// @Tags aircraft
// @Produce  json
// @Param   airline_id        path    string     true        "Airline ID"
// @Param gate body models.Aircraft true "New Aircraft object"
// @Success 201  "Aircraft created"
// @Failure 500 "Internal server error"
func (ac *AircraftController) HandleCreateNewAircraft(ctx *gin.Context) {
	var aircraft models.Aircraft

	airlineId := ctx.Param("airline_id")
	err := ctx.ShouldBindWith(&aircraft, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ac.repository.InsertAircraft(aircraft, airlineId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, "Aircraft created Successfully")
}

// @Router /airlines/{airline_id}/aircraft/{id} [put]
// @Summary Update aircraft by ID
// @Description update the aircraft details by its ID
// @ID update-aircraft
// @Tags aircraft
// @Produce  json
// @Param id path string true "Aircraft ID"
// @Param airline_id path string true "Airlines ID"
// @Param aircraft body models.Aircraft true "Updated aircraft object"
// @Success 200  "ok"
// @Failure 400 "Bad request"
func (ac *AircraftController) HandleUpdateAircraft(ctx *gin.Context) {
	var aircraft models.Aircraft
	airlineId := ctx.Param(`airline_id`)
	aircraftId := ctx.Param(`id`)

	err := ctx.ShouldBindWith(&aircraft, binding.JSON)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	repositoryError := ac.repository.UpdateAircraft(&aircraft, aircraftId, airlineId)
	if repositoryError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": repositoryError.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "update success"})
}

// @Summary Delete aircraft by Id
// @Router /aircraft/{id} [delete]
// @Summary Delete aircraft by Id
// @Description Delete the aircraft details of the particular id
// @ID delete-aircraft-by-id
// @Tags aircraft
// @Param id path string true "Aircraft Id"
// @Success 200  "ok"
// @Failure 400 "Aircraft not found"
func (ac *AircraftController) HandleDeleteAircraft(ctx *gin.Context) {
	aircraftId := ctx.Param("id")
	err := ac.repository.DeleteAircraft(aircraftId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Incorrect aircraft id: " + aircraftId})
		return
	}
	ctx.JSON(http.StatusOK, "Deleted the aircraft successfully")
}
