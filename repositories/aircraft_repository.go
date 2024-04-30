package repositories

import (
	"airport-app-backend/models"
)

type IAircraftRepository interface {
	RetrieveAllAircrafts(int, int, int) ([]models.Aircraft, error)
	RetrieveAircraft(id string) (*models.Aircraft, error)
	InsertAircraft(models.Aircraft, string) error
	UpdateAircraft(*models.Aircraft, string, string) error
}

func (sr *ServiceRepository) RetrieveAllAircrafts(page, capacity, year int) ([]models.Aircraft, error) {
	var aircrafts []models.Aircraft
	query := sr.db.Offset(page * DEFAULT_PAGE_SIZE).Limit(DEFAULT_PAGE_SIZE)
	if capacity != -1 {
		query = query.Where("capacity >= ?", capacity)
	}
	if year != -1 {
		query = query.Where("year_of_manufacture = ?", year)
	}
	if err := query.Find(&aircrafts).Error; err != nil {
		return nil, err
	}
	return aircrafts, nil
}

func (sr ServiceRepository) RetrieveAircraft(id string) (*models.Aircraft, error) {
	var aircraft *models.Aircraft
	result := sr.db.First(&aircraft, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return aircraft, nil
}

func (sr *ServiceRepository) InsertAircraft(aircraft models.Aircraft, airlineId string) error {
	aircraft.AirlineId = airlineId
	result := sr.db.Save(&aircraft)
	return result.Error
}

func (sr ServiceRepository) UpdateAircraft(aircraft *models.Aircraft, aircraftId string, airlineId string) error {
	aircraft.Id = aircraftId
	aircraft.AirlineId = airlineId
	result := sr.db.Updates(aircraft)
	return result.Error
}
