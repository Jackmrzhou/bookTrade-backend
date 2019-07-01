package models

type Contact struct {
	ID int `gorm:"primary_key" json:"id"`
	SelfID int `json:"self_id"`
	CounterpartID int `json:"counterpart_id"`
}

func (Contact) TableName() string {
	return "contact"
}