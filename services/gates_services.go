package services

import (
	"airport-app-backend/middleware"
	"airport-app-backend/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"

	"github.com/rs/zerolog/log"
)

var DEFAULT_PAGE_SIZE = 10

type IGateRepository interface {
	GetGates(page int, floor int, c context.Context, ctx *gin.Context) ([]models.Gate, error)
}

func (sr *ServiceRepository) GetGates(page, floor int, c context.Context, ctx *gin.Context) ([]models.Gate, error) {
	_, span := trace.StartSpan(c, "get_gates")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)
	log.Debug().Msg("Fetching list of gates")
	var gates []models.Gate
	offset := (page - 1) * DEFAULT_PAGE_SIZE
	query := sr.db.Offset(offset).Limit(DEFAULT_PAGE_SIZE)
	fmt.Printf("offset : %v page : %v floor: %v", offset, page, floor)
	if floor != -1 {
		query = query.Where("floor_number = ?", floor)
	}
	if err := query.Find(&gates).Error; err != nil {
		return nil, err
	}
	return gates, nil
}
