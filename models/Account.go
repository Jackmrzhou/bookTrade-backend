package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Status   int
	Email    string
	Password string
}

func (Account) TableName() string {
	return "account"
}

const (
	STATUS_IVALID = 1 << 2
	STATUS_BANNED = 1 << 1
	STATUS_ACTIVE = 1
)

type VerifyCode struct {
	ID         int `gorm:"primary_key"`
	Code       string
	Email      string
	ExpireTime time.Time
}

func (VerifyCode) TableName() string {
	return "verify_code"
}
