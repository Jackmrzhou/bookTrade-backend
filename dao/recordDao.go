package dao

import (
	"bookTrade-backend/models"
	"github.com/jinzhu/gorm"
)

func CreateUserRecord(record *models.UserRecord) error {
	return db.Create(record).Error
}

func UpdateOrCreateUserRecord(record *models.UserRecord) error {
	var r models.UserRecord
	if err := db.Where("user_id = ? AND book_id = ?", record.UserID, record.BookID).First(&r).Error; err != nil{
		if gorm.IsRecordNotFoundError(err){
			return CreateUserRecord(record)
		}else {
			return err
		}
	} else {
		r.ViewedTime = record.ViewedTime
		return db.Save(r).Error
	}
}

func GetRecordsByBookID(bookID int) ([]models.UserRecord, error) {
	var rs []models.UserRecord
	if err := db.Where("book_id = ?", bookID).Find(&rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}