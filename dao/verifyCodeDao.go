package dao

import "bookTrade-backend/models"

func CreateVerifyCode(code *models.VerifyCode) error {
	return db.Create(code).Error
}

func GetVerifyCodeByEmailAndCode(email, code string) (*models.VerifyCode, error) {
	verifyCode := &models.VerifyCode{}
	err := db.Where("email = ? AND code = ?", email, code).First(verifyCode).Error
	return verifyCode, err
}

func DeleteVerifyCode(code *models.VerifyCode) error {
	return db.Delete(code).Error
}
