package services

import (
	"airport-app-backend/models"

	"github.com/rs/zerolog/log"
)

var DEFAULT_PAGE_SIZE = 10

type IGateRepository interface {
	GetGates(int) ([]models.Gate, error)
}

func (sr *ServiceRepository) GetGates(page int) ([]models.Gate, error) {
	log.Debug().Msg("Fetching list of gates")
	var gates []models.Gate
	offset := (page - 1) * DEFAULT_PAGE_SIZE

	if err := sr.db.Offset(offset).Limit(DEFAULT_PAGE_SIZE).Find(&gates).Error; err != nil {
		return nil, err
	}
	return gates, nil
}
