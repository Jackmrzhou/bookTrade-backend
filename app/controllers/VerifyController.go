package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/app/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type sendVerifyCodeQuery struct {
	Email string `json:"email" binding:"required"`
}

type sendVerifyCodeResp struct {
	VerifyCode string `json:"verify_code"`
}

func SendVerifyCode(c *gin.Context) {
	var query sendVerifyCodeQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	if code, err := services.SendCode(query.Email); err != nil {
		log.WithError(err).Error("send code failed")
		c.Set(error2.CodeKey, error2.SendCode)
		return
	} else {
		c.JSON(http.StatusOK, sendVerifyCodeResp{
			VerifyCode: code,
		})
	}
}
