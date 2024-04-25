package controllers

import (
	"airport-app-backend/services"

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

func (acr * AircraftControllerRepository) HandleGetAircrafts(context *gin.Context) {
	
}