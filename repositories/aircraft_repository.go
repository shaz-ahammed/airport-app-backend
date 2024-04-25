package repositories

import (
	"airport-app-backend/models"
)

type IAircraftRepository interface {
	RetrieveAllAircrafts(int) ([]models.Aircraft, error)
}

func (sr *ServiceRepository) RetrieveAllAircrafts(page int) ([]models.Aircraft, error) {
	return nil, nil
}
