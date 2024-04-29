package factory

import (
	"airport-app-backend/models"

	"github.com/go-faker/faker/v4"
)

func ConstructGate() models.Gate {
	return models.Gate{GateNumber: GenerateRandomInt(), FloorNumber: GenerateRandomInt()}
}

func GenerateRandomInt() int {
	arr, _ := faker.RandomInt(1, 20, 1)
	return arr[0]
}
