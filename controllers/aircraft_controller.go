package controllers

import (
	"airport-app-backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AircraftControllerRepository struct {
	service services.IAircraftRepository
}

func NewAircraftControllerRepository(service services.IAircraftRepository) *AircraftControllerRepository {
	return &AircraftControllerRepository{
		service: service,
	}
}

// @Summary Get aircrafts
// @Router /aircrafts [get]
// @Description get all the aircrafts details
// @ID get-all-aircrafts
// @Tags aircraft
// @Produce  json
// @Param   page        query    int     false        "Page number (default = 1)"
// @Param   type       query    int     false        "filter by type of aircraft (default = all type) options : [passenger, cargo, helicopter]"
// @Param   year       query    int     false        "filter by manufacturing (default = all year)"
// @Param   capacity       query    int     false        "condition by capacity grater than given value (default = 0)"
// @Success 200  "ok"
// @Failure 500 "Internal server error"
func (acr *AircraftControllerRepository) HandleGetAircrafts(context *gin.Context) {
	pageStr := context.Query("page")
	yearStr := context.Query("year")
	capacityStr := context.Query("capacity")
	aircraftType := context.Query("type")


	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 1970 {
		year = -1
	}
	capacity, err := strconv.Atoi(capacityStr)
	if err != nil || capacity < 0 {
		capacity = -1
	}
}
