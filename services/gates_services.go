package services

import (
	"airport-app-backend/models"

	"github.com/rs/zerolog/log"
)

type IGateRepository interface {
	GetGates() []models.Gates
}

func (sr *ServiceRepository) GetGates() []models.Gates {
	log.Debug().Msg("Fetching list of gates")
	listOfGates := make([]models.Gates, 3)
	listOfGates = append(listOfGates, models.Gates{GateNumber: 1, FloorNumber: 1})
	listOfGates = append(listOfGates, models.Gates{GateNumber: 2, FloorNumber: 2})
	listOfGates = append(listOfGates, models.Gates{GateNumber: 3, FloorNumber: 3})
	return listOfGates
}
