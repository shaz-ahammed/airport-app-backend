package services

import (
	"airport-app-backend/models"
)



type IAirlineRepository interface {
	GetAirline(int) ([]models.Airlines, error)
}

var DEFAULT_PAGE_LIMIT int = 10

func (sr *ServiceRepository) GetAirline(pageNum int) ([]models.Airlines, error) {
	var airlines []models.Airlines
	result := sr.db.Limit(DEFAULT_PAGE_LIMIT).Offset(pageNum*DEFAULT_PAGE_LIMIT).Find(&airlines)
	if result.Error != nil {
		return nil, result.Error
	}
	return airlines,nil
}