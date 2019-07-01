package dao

import (
	"bookTrade-backend/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func CreateUser(user *models.User) error {
	return db.Create(user).Error
}

func UpdateOrCreateUser(user *models.User) error {
	var u models.User
	if err := db.Where("account_id = ?", user.AccountID).First(&u).Error; err != nil {
		logrus.Info(err)
		if gorm.IsRecordNotFoundError(err) {
			return CreateUser(user)
		} else {
			return err
		}
	} else {
		u.Username = user.Username
		u.Introduction = user.Introduction
		u.AvatarKey = user.AvatarKey
		return db.Save(&u).Error
	}
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetUserByAccountID(accountID int) (*models.User, error) {
	var user models.User
	err := db.Where("account_id = ?", accountID).First(&user).Error
	return &user, err
}

func GetUserByUserID(userID int) (*models.User, error) {
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	return &user, err
}