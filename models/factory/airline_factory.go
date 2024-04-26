package factory

import (
	"airport-app-backend/models"

	"github.com/go-faker/faker/v4"
)

func ConstructAirline() models.Airline {
	return models.Airline{Name: faker.Name()}
}
