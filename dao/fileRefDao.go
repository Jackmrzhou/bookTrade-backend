package dao

import "bookTrade-backend/models"

func CreateFileRef(ref *models.FileRef) error {
	return db.Create(ref).Error
}

func GetFileRefByFileKey(key string) (*models.FileRef, error) {
	fileRef := models.FileRef{}
	err := db.Where("file_key = ?", key).First(&fileRef).Error
	return &fileRef, err
}
