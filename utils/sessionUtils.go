package utils

import (
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

func ValidateSessionID(sessionID string) (bool, error) {
	if session, err := dao.GetSessionBySessionID(sessionID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		} else {
			return false, err
		}
	} else if time.Now().After(session.ExpireTime) {
		return false, nil
	}
	return true, nil
}

func NewSession(accountID uint) *models.Session {
	uid, _ := uuid.NewRandom()
	//logrus.Info(uid.String())
	//logrus.Info(int(accountID))
	return &models.Session{
		AccountID:  int(accountID),
		SessionID:  uid.String(),
		ExpireTime: time.Now().Add(48 * time.Hour),
	}
}

func newSeesionID() string {
	s, _ := uuid.NewRandom()
	return s.String()
}

func RenewSession(session *models.Session) {
	// allow login in multiple devices
	// session.SessionID = newSeesionID()
	session.ExpireTime = time.Now().Add(48 * time.Hour)
}

func GetUserIDBySessionID(session string) (int, error) {
	if session, err := dao.GetSessionBySessionID(session); err != nil {
		return 0, err
	}else {
		if user, err := dao.GetUserByAccountID(session.AccountID); err != nil {
			return 0, err
		}else {
			return user.ID, nil
		}
	}
}