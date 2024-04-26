package models

type Airline struct {
	Id    string `json:"id"   gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" binding:"required" gorm:"unique;notNull;size:100"`
	Count int    `json:"count" gorm:"null"`
}

func (airline *Airline) SetName(name string) Airline {
	(*airline).Name = name
	return *airline
}

func (airline *Airline) SetCount(count int) Airline {
	(*airline).Count = count
	return *airline
}
