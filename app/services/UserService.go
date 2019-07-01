package services

import (
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

func CreateDefaultProfile(accountID uint) error {
	user := models.User{
		AccountID:    int(accountID),
		Username:     "用户" + strconv.Itoa(int(accountID)+10000),
		AvatarKey:    "",
		Introduction: "",
	}

	return dao.CreateUser(&user)
}

func TestUsername(username string, accountID int) (bool, error) {
	if user, err := dao.GetUserByUsername(username); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	} else if accountID != user.AccountID {
		return true, nil
	}
	return false, nil
}
