package services

import (
	"airport-app-backend/models"
	"fmt"

	"github.com/rs/zerolog/log"
)

var DEFAULT_PAGE_SIZE = 10

type IGateRepository interface {
	GetGates(int, int) ([]models.Gate, error)
}

func (sr *ServiceRepository) GetGates(page, floor int) ([]models.Gate, error) {
	log.Debug().Msg("Fetching list of gates")
	var gates []models.Gate
    offset := (page - 1) * DEFAULT_PAGE_SIZE
    query := sr.db.Offset(offset).Limit(DEFAULT_PAGE_SIZE)
	fmt.Printf("offset : %v page : %v floor: %v", offset,page,floor)
    if floor != -1 {
        query = query.Where("floor_number = ?", floor)
    }
    if err := query.Find(&gates).Error; err != nil {
        return nil, err
    }
    return gates, nil
}
