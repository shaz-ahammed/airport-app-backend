package repositories

import (
	"airport-app-backend/models"
	"github.com/rs/zerolog/log"
)

var DEFAULT_PAGE_SIZE = 10

type IGateRepository interface {
	GetAllGates(page int, floor int) ([]models.Gate, error)
	GetGate(string) (*models.Gate, error)
	CreateNewGate(*models.Gate) error
	UpdateGate(string, models.Gate) error
}

func (sr *ServiceRepository) GetAllGates(page, floor int) ([]models.Gate, error) {
	log.Debug().Msg("Fetching list of gates")

	var gates []models.Gate
	offset := page * DEFAULT_PAGE_SIZE
	query := sr.db.Offset(offset).Limit(DEFAULT_PAGE_SIZE)
	if floor != -1 {
		query = query.Where("floor_number = ?", floor)
	}
	if err := query.Find(&gates).Error; err != nil {
		return nil, err
	}
	return gates, nil
}

func (sr *ServiceRepository) GetGate(id string) (*models.Gate, error) {
	log.Debug().Msg("repository layer for retrieving gate details by id")

	var gate models.Gate
	if err := sr.db.Where("id = ?", id).First(&gate).Error; err != nil {
		return nil, err
	}
	return &gate, nil
}

func (sr *ServiceRepository) CreateNewGate(gate *models.Gate) error {
	err := sr.db.Save(gate)
	return err.Error
}

func (sr *ServiceRepository) UpdateGate(id string, gate models.Gate) error {
	err := sr.db.Where("id=?", id).Updates(gate)
	return err.Error
}
