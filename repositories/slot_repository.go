package repositories

import (
	"airport-app-backend/models"
)

type ISlotRepository interface {
	RetrieveAllSlots(int, bool) ([]models.Slot, error)
}

func (sr *ServiceRepository) RetrieveAllSlots(page int, status bool) ([]models.Slot, error) {
	var slots []models.Slot
	statusArr := []string{}

	if status {
		statusArr = append(statusArr, "Available")
	} else {
		statusArr = append(statusArr, "Reserved", "Booked")
	}

	result := sr.db.Where("status IN (?)", statusArr).
		Offset(page * DEFAULT_PAGE_SIZE).Limit(DEFAULT_PAGE_SIZE).
		Find(&slots)
	if result.Error != nil {
		return nil, result.Error
	}
	return slots, nil
}
