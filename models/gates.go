package models

import (
    "github.com/google/uuid"
    _"gorm.io/gorm"
)

type Gate struct {
    ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
    GateNumber   int       `gorm:"unique;not null"`
    FloorNumber  int       `gorm:"not null"`
}

// TableName specifies the table name for the Gate model
func (Gate) TableName() string {
    return "gates"
}
