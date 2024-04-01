package models

import "gorm.io/gorm"

type Aircrafts struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	WingNumber string `json:"wing_number" gorm:"unique;size:24"`
	Capacity   int    `json:"capacity" gorm:"notNull"`
}

func AircraftsTable(db *gorm.DB) error {
	err := db.AutoMigrate(&Aircrafts{})
	if err != nil {
		return err
	}
	result := db.Exec("ALTER TABLE aircrafts DROP COLUMN IF EXISTS name;")
    if result.Error != nil {
        return result.Error
    }
	return nil
}
