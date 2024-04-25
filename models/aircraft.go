package models

import "time"

type AircraftType string

const (
	Passenger  AircraftType = "passenger"
	Cargo      AircraftType = "cargo"
	Helicopter AircraftType = "helicopter"
)

type Aircraft struct {
	Id                string       `json:"id"   gorm:"primaryKey;autoIncrement"`
	WingNumber        string       `json:"wing_number" binding:"required" gorm:"unique;notNull;size:100"`
	Type              AircraftType `json:"aircraft_type" binding:"oneof=passenger cargo helicopter" gorm:"notNull"`
	Capacity          int          `json:"capacity" binding:"required" gorm:"notNull"`
	YearOfManufacture int          `json:"year_of_manufacture" gorm:"notNull"`
	CreatedAt         time.Time    `json:"created_at"`
}
