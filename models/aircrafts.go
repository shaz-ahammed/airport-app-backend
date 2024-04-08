package models

type Aircrafts struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	WingNumber string `json:"wing_number" gorm:"unique;size:24"`
	Capacity   int    `json:"capacity" gorm:"notNull"`
}
