package services

import (
	"airport-app-backend/middleware"
	"airport-app-backend/models"

	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
)

var DEFAULT_PAGE_SIZE = 10

type IGateRepository interface {
  
  GetGates(page int, floor int, c context.Context, ctx *gin.Context) ([]models.Gate, error)
	GetGateByID(context.Context, *gin.Context, string) (*models.Gate, error)
	
}

func (sr *ServiceRepository) GetGates(page, floor int, c context.Context, ctx *gin.Context) ([]models.Gate, error) {
	_, span := trace.StartSpan(c, "get_gates")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)
	log.Debug().Msg("Fetching list of gates")
	var gates []models.Gate
	offset := (page - 1) * DEFAULT_PAGE_SIZE
	query := sr.db.Offset(offset).Limit(DEFAULT_PAGE_SIZE)
	if floor != -1 {
		query = query.Where("floor_number = ?", floor)
	}
	if err := query.Find(&gates).Error; err != nil {
		return nil, err
	}
	return gates, nil
}

func (sr *ServiceRepository) GetGateByID(c context.Context, ctx *gin.Context, id string) (*models.Gate, error) {
	log.Debug().Msg("service layer for retrieving gate details by id")
	_, span := trace.StartSpan(c, "get_gate_by_id")
	defer span.End()

	middleware.TraceSpanTags(span)(ctx)

	var gate models.Gate
	if err := sr.db.Where("id = ?", id).First(&gate).Error; err != nil {
		return nil, err
	}
	return &gate, nil
}
