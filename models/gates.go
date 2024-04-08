package models

type Gates struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	GateNumber  int  `json:"gateNumber" gorm:"notNull"`
	FloorNumber int  `json:"floorNumber" gorm:"notNull"`
}
