package models

type Airlines struct {
	Id         string   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"unique;notNull;size:100"`
}