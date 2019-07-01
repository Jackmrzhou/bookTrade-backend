package dao

import (
	"bookTrade-backend/models"
	"github.com/jinzhu/gorm"
)

func CreateMessage(message *models.Message) error {
	return db.Create(message).Error
}

func CreateContactIfNotExists(contact *models.Contact) error {
	var c models.Contact
	if err := db.Where("self_id = ? AND counterpart_id = ?", contact.SelfID, contact.CounterpartID).First(&c).Error; err != nil {
		if gorm.IsRecordNotFoundError(err){
			return db.Create(contact).Error
		}
		return err
	}else {
		contact = &c
		return nil
	}
}

func GetAllContactsBySelfID(selfID int) ([]models.Contact, error) {
	var cs []models.Contact
	err := db.Where("self_id = ?", selfID).Find(&cs).Error
	return cs, err
}

func GetAllMessageByContactID(contactID int) ([]models.Message, error) {
	var msgs []models.Message
	err := db.Where("contact_id = ?", contactID).Find(&msgs).Error
	return msgs,err
}

func CountMessageUnreadByContactID(contactID int) (int, error) {
	var count int
	err := db.Model(&models.Message{}).
		Where("is_read = ? AND contact_id = ?", models.MESSAGE_UNREAD, contactID).
		Count(&count).Error
	return count, err
}

func CountMessageUnread(userID int) (int, error) {
	var count int
	err := db.Model(&models.Message{}).Where("is_read = ? AND to_id = ?", models.MESSAGE_UNREAD, userID).
		Count(&count).Error
	return count, err
}

func SaveAllMsgs(msgs []models.Message) error {
	return db.Model(&models.Message{}).Updates(msgs).Error
}