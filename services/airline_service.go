package services

import (
	"airport-app-backend/middleware"
	"airport-app-backend/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
)

type IAirlineRepository interface {
	GetAirline(pageNum int, c context.Context, ctx *gin.Context) ([]models.Airlines, error)
	GetAirlineById(context.Context, *gin.Context, string) (*models.Airlines, error)
}

var DEFAULT_PAGE_LIMIT int = 10

func (sr *ServiceRepository) GetAirline(pageNum int, c context.Context, ctx *gin.Context) ([]models.Airlines, error) {
	_, span := trace.StartSpan(c, "get_airline")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)

	var airlines []models.Airlines
	result := sr.db.Limit(DEFAULT_PAGE_LIMIT).Offset(pageNum * DEFAULT_PAGE_LIMIT).Find(&airlines)
	if result.Error != nil {
		return nil, result.Error
	}
	return airlines, nil
}
func (sr *ServiceRepository) GetAirlineById(c context.Context, ctx *gin.Context, id string) (*models.Airlines, error) {
	_, span := trace.StartSpan(c, "get_airline_by_id")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)
	var airlines *models.Airlines
	result := sr.db.First(&airlines, "id=?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return airlines, nil
}
