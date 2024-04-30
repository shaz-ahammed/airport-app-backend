package models

// type AircraftType string
// const (
// 	Passenger  AircraftType = "passenger"
// 	Cargo      AircraftType = "cargo"
// 	Helicopter AircraftType = "helicopter"
//   )

type Aircraft struct {
	Id         string `json:"id"   gorm:"primaryKey;autoIncrement"`
	TailNumber string `json:"tail_number" binding:"required" gorm:"unique;notNull;size:100"`
	Capacity   int    `json:"capacity" binding:"required" gorm:"notNull"`
	// Type              AircraftType `json:"aircraft_type" binding:"oneof=passenger cargo helicopter" gorm:"notNull"`
	YearOfManufacture int `json:"year_of_manufacture"`
	// CreatedAt         time.Time    `json:"created_at"`
	AirlineId string `json:"airline_id" gorm:"foreignKey:Airline"`
}
