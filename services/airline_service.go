package services

import (
	"airport-app-backend/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
	"regexp"
)

type IAirlineRepository interface {
	GetAirline(pageNum int, c context.Context, ctx *gin.Context) ([]models.Airlines, error)
	GetAirlineById(context.Context, *gin.Context, string) (*models.Airlines, error)
	CreateNewAirline(c context.Context, ctx *gin.Context, airline *models.Airlines) error
}

var DEFAULT_PAGE_LIMIT int = 10

func (sr *ServiceRepository) GetAirline(pageNum int) ([]models.Airlines, error) {
	var airlines []models.Airlines
	result := sr.db.Limit(DEFAULT_PAGE_LIMIT).Offset(pageNum * DEFAULT_PAGE_LIMIT).Find(&airlines)
	if result.Error != nil {
		return nil, result.Error
	}
	return airlines, nil
}

func (sr *ServiceRepository) GetAirlineById(id string) (*models.Airlines, error) {
	var airlines *models.Airlines
	result := sr.db.First(&airlines, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return airlines, nil
}

func (sr *ServiceRepository) CreateNewAirline(c context.Context, ctx *gin.Context, airline *models.Airlines) error {
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
