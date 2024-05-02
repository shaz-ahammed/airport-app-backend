package factory

import (
	"airport-app-backend/models"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
)

func ConstructSlot() models.Slot {
	return models.Slot{StartTime: GenerateRandomTimestamp(),
		EndTime: GenerateRandomTimestamp(),
		Status:  GenerateRandomStatus()}
}

func GenerateRandomTimestamp() time.Time {
	format := "2006-01-02 15:04:05"
	timeObject, _ := time.Parse(format, faker.Timestamp())
	return timeObject
}

func GenerateRandomStatus() string {
	statusList := []string{"Available", "Booked", "Reserved"}
	return statusList[rand.Intn(len(statusList))]
}
