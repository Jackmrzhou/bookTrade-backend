package models

import "time"

type Session struct {
	ID         int `gorm:"primary_key"`
	SessionID  string
	AccountID  int
	ExpireTime time.Time
}

func (Session) TableName() string {
	return "session"
}
