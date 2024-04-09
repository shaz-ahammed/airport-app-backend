package services

import (
	"airport-app-backend/models"

	"github.com/rs/zerolog/log"
)

type IGateRepository interface {
	GetGates() []models.Gate
}

func (sr *ServiceRepository) GetGates() []models.Gate {
	log.Debug().Msg("Fetching list of gates")
	var gates []models.Gate
	// Query all gates from the database
	if err := sr.db.Find(&gates).Error; err != nil {
		log.Printf("Failed to fetch gates: %v", err)
		return nil
	}
	return gates
}
