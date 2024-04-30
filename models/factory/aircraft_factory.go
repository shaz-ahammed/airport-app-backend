package factory

import (
	"airport-app-backend/models"

	"github.com/go-faker/faker/v4"
)

func ConstructAircraft() models.Aircraft {
	return models.Aircraft{TailNumber: faker.CCType()}
}
