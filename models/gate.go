package models

import (
	_ "gorm.io/gorm"
)

type Gate struct {
	Id          string `json:"id" gorm:"primaryKey;autoIncrement"`
	GateNumber  int    `json:"gate_number" binding:"required" gorm:"unique;not null"`
	FloorNumber int    `json:"floor_number" binding:"required" gorm:"not null"`
}

// TableName specifies the table name for the Gate model
func (Gate) TableName() string {
	return "gates"
}
