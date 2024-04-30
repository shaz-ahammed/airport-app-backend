package factory

import (
	"airport-app-backend/models"

	"github.com/go-faker/faker/v4"
)

var DEFAULT_CAPACITY = 30

func ConstructAircraft() models.Aircraft {
	return models.Aircraft{TailNumber: faker.CCType(), Capacity: GenerateRandomInt() + DEFAULT_CAPACITY}
}
