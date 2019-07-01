package dao

import "bookTrade-backend/models"

func CreateAccount(account *models.Account) error {
	return db.Create(account).Error
}

func GetAccount(email, passwd string) (*models.Account, error) {
	account := models.Account{}
	err := db.Where("email = ? AND password = ?", email, passwd).First(&account).Error
	return &account, err
}

func GetAccountByEmail(email string) (*models.Account, error) {
	account := models.Account{}
	err := db.Where("email = ?", email).First(&account).Error
	return &account, err
}
