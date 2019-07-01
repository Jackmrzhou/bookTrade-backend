package services

import (
	"bookTrade-backend/app"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func SendCode(email string) (string, error) {
	code := generateCode()
	verifyCode := models.VerifyCode{
		Code:       code,
		Email:      email,
		ExpireTime: time.Now().Add(10 * time.Minute),
	}
	if err := dao.CreateVerifyCode(&verifyCode); err != nil {
		log.Error("store verification code failed")
		return "", err
	} else if err := app.App.MailManager.SendMail(email, code); err != nil {
		// ok to fail
		dao.DeleteVerifyCode(&verifyCode)
		return "", err
	}
	return code, nil
}

func ValidateCode(email, code string) (bool, error) {
	verifyCode, err := dao.GetVerifyCodeByEmailAndCode(email, code)
	if gorm.IsRecordNotFoundError(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	if time.Now().After(verifyCode.ExpireTime) {
		return false, nil
	}
	return true, nil
}

func generateCode() string {
	var source = rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	num := r.Intn(900000) + 100000
	// map to [10 0000, 100 0000)
	return fmt.Sprint(num)
}
