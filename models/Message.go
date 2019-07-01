package models

import "time"

type Message struct {
	ID int `gorm:"primary_key" json:"id"`
	ContactID int `json:"contact_id"`
	FromID int `json:"from_id"`
	ToID int `json:"to_id"`
	Content string `json:"content"`
	CreateTime time.Time `json:"create_time"`
	IsRead int `json:"-"`
}

func (Message) TableName() string {
	return "message"
}

const (
	MESSAGE_READ = 1
	MESSAGE_UNREAD = 0
)