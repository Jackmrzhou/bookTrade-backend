package models

import "time"

type UserRecord struct {
	ID int `gorm:"primary_key"`
	UserID int
	BookID int
	ViewedTime time.Time
}

func (UserRecord) TableName() string {
	return "user_record"
}