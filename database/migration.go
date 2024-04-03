package database

import (
	"airport-app-backend/models"

	"gorm.io/gorm"
)

func Aircrafts(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Aircrafts{})
	if err != nil {
		return err
	}
	result := db.Exec("ALTER TABLE aircrafts DROP COLUMN IF EXISTS name;")
	if result.Error != nil {
		return result.Error
	}
	return nil
}
