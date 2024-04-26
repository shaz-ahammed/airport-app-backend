package factory

import (
	"airport-app-backend/models"

	"github.com/go-faker/faker/v4"
)

func ConstructAirline() *models.Airline {
	airline := models.Airline{}
	return airline.SetName(faker.Name())
}
