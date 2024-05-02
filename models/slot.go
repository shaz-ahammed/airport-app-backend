package models

import "time"

type Slot struct {
	Id         string    `json:"id" gorm:"primary_key;autoIncrement"`
	StartTime  time.Time `json:"start_time" binding:"required" gorm:"unique;notNull"`
	EndTime    time.Time `json:"end_time" binding:"required" gorm:"unique;notNull"`
	Status     string    `json:"status" gorm:"notNull"`
	AircraftID string    `json:"aircraft_id" binding:"required" gorm:"foreignKey:Aircraft"`
	GateID     string    `json:"gate_id" binding:"required" gorm:"foreignKey:Gate"`
}

func (slot *Slot) SetStatus(status string) Slot {
	(*slot).Status = status
	return *slot
}
