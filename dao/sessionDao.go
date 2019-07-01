package dao

import "bookTrade-backend/models"

func CreateSession(session *models.Session) error {
	return db.Create(session).Error
}

func GetSessionBySessionID(sessionID string) (*models.Session, error) {
	session := models.Session{}
	err := db.Where("session_id = ?", sessionID).First(&session).Error
	return &session, err
}

func UpdateSession(session *models.Session) error {
	return db.Save(session).Error
}

func GetSessionByAccountID(accountID uint) (*models.Session, error) {
	session := models.Session{}
	err := db.Where("account_id = ?", accountID).First(&session).Error
	return &session, err
}
