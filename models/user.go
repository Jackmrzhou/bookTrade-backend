package models

type User struct {
	ID           int `gorm:"primary_key"`
	AccountID    int
	Username     string
	AvatarKey    string
	Introduction string
	PhoneNumber string
	Address string
}

func (User) TableName() string {
	return "user"
}
