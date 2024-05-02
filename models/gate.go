package models

type Gate struct {
	Id          string `json:"id" gorm:"primaryKey;autoIncrement"`
	GateNumber  int    `json:"gate_number" binding:"required" gorm:"unique;not null"`
	FloorNumber int    `json:"floor_number" binding:"required" gorm:"not null"`
}
