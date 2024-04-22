package services

import (
	"airport-app-backend/models"
	"errors"

	"regexp"
)

type IAirlineRepository interface {
	GetAirline(int) ([]models.Airline, error)
	GetAirlineById(string) (*models.Airline, error)
	CreateNewAirline(*models.Airline) error
}

var DEFAULT_PAGE_LIMIT int = 10

func (sr *ServiceRepository) GetAirline(pageNum int) ([]models.Airline, error) {
	var airline []models.Airline
	result := sr.db.Limit(DEFAULT_PAGE_LIMIT).Offset(pageNum * DEFAULT_PAGE_LIMIT).Find(&airline)
	if result.Error != nil {
		return nil, result.Error
	}
	return airline, nil
}

func (sr *ServiceRepository) GetAirlineById(id string) (*models.Airline, error) {
	var airline *models.Airline
	result := sr.db.First(&airline, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return airline, nil
}

func (sr *ServiceRepository) CreateNewAirline(airline *models.Airline) error {

	if !(containsOnlyCharacters(airline.Name)) {
		return errors.New("name should not contain numbers")
	}
	result := sr.db.Save(airline)
	return result.Error
}

func containsOnlyCharacters(s string) bool {
	re := regexp.MustCompile("^[A-Za-z ]+$")
	return re.MatchString(s)
}

func (sr *ServiceRepository) CreateNewAirline(c context.Context, ctx *gin.Context, airline *models.Airline) error {
	_, span := trace.StartSpan(c, "get_airline_by_id")
	defer span.End()
	middleware.TraceSpanTags(span)(ctx)

	if !(containsOnlyCharacters(airline.Name)) {
		return errors.New("name should not contain Numbers")
	}
	result := sr.db.Save(airline)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func containsOnlyCharacters(s string) bool {
	re := regexp.MustCompile("^[A-Za-z ]+$")
	return re.MatchString(s)
}
