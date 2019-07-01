package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/app/services"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type createNewAccountReq struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	VerifyCode string `json:"verify_code" binding:"required"`
	// verification code: means you whether owns that email address
	// captcha: means distinguish human or machine
}

type createNewAccountResp struct {
	SessionID string `json:"session_id"`
}

func CreateNewAccount(c *gin.Context) {
	var query createNewAccountReq
	if err := c.ShouldBindJSON(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}
	if ok, err := services.ValidateCode(query.Email, query.VerifyCode); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
		return
	} else if ok {
		// check if account exists
		if _, err := dao.GetAccountByEmail(query.Email); err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				c.Set(error2.CodeKey, error2.ServerError)
				return
			}
		} else {
			c.Set(error2.CodeKey, error2.AccountExists)
			return
		}

		// crate account
		account := &models.Account{
			Status:   models.STATUS_ACTIVE,
			Email:    query.Email,
			Password: utils.StrToSha1(query.Password),
		}
		if err := dao.CreateAccount(account); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			session := utils.NewSession(account.ID)
			// todo: error check
			dao.CreateSession(session)
			c.JSON(http.StatusOK, createNewAccountResp{
				SessionID: session.SessionID,
			})
		}
	} else {
		c.Set(error2.CodeKey, error2.InvalidCode)
	}
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
