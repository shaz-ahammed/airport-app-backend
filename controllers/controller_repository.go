package controllers

import (
	"airport-app-backend/services"
)

type ControllerRepository struct {
	service services.ServiceRepository
}

func NewControllerRepository(service services.ServiceRepository) ControllerRepository {
	return ControllerRepository{
		service: service,
	}
}
