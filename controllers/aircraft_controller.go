package controllers

import (
	"airport-app-backend/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
// @Param   type       query    int     false        "filter by type of aircraft (default = all type) options : [passenger, cargo, helicopter]"
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
	// year, err := strconv.Atoi(ctx.Query("year"))
	// if err != nil || year < 1970 {
	// 	year = -1
	// }
	// capacity, err := strconv.Atoi(ctx.Query("capacity"))
	// if err != nil || capacity < 0 {
	// 	capacity = -1
	// }

	aircrafts, err := ac.repository.RetrieveAllAircrafts(page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, aircrafts)
}
