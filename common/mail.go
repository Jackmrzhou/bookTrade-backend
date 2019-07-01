package common

import (
	"bookTrade-backend/conf"
	"fmt"
	"gopkg.in/gomail.v2"
)

type MailSender interface {
	SendMail(email, body string) error
}

type defaultMailSender struct {
	mailAddress string
	mailToken   string
}

func (mailSender *defaultMailSender) SendMail(email, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mailSender.mailAddress)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "ZJU旧书交易网站验证码")
	m.SetBody("text/html", fmt.Sprintf("%s, 10分钟内有效。", code))

	d := gomail.NewDialer("smtp.qq.com", 587, mailSender.mailAddress, mailSender.mailToken)

	err := d.DialAndSend(m)
	return err
}

func NewDefaultMailSender(config *conf.AppConfig) *defaultMailSender {
	return &defaultMailSender{
		mailAddress: config.MailConifg.Address,
		mailToken:   config.MailConifg.Token,
	}
}
