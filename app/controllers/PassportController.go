package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/dao"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type loginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResp struct {
	AccountID uint   `json:"account_id"`
	SessionID string `json:"session_id"`
}

func Login(c *gin.Context) {
	var query loginReq
	if err := c.ShouldBindJSON(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	p := utils.StrToSha1(query.Password)

	if account, err := dao.GetAccount(query.Email, p); err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			log.WithError(err).Error("get account failed")
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			c.Set(error2.CodeKey, error2.LoginError)
		}
	} else {
		if session, err := dao.GetSessionByAccountID(account.ID); err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				log.WithError(err).Error("get session failed")
				c.Set(error2.CodeKey, error2.ServerError)
			} else {
				// create session
				newSession := utils.NewSession(account.ID)
				//log.Info(*newSession)
				if err := dao.CreateSession(newSession); err != nil {
					log.WithError(err).Error("create session failed")
					c.Set(error2.CodeKey, error2.ServerError)
				} else {
					// return session
					c.JSON(http.StatusOK, loginResp{
						AccountID: account.ID,
						SessionID: newSession.SessionID,
					})
				}
			}
		} else {
			// got session, do update
			utils.RenewSession(session)
			if err := dao.UpdateSession(session); err != nil {
				log.WithError(err).Error("update session failed")
				c.Set(error2.CodeKey, error2.ServerError)
			} else {
				c.JSON(http.StatusOK, loginResp{
					AccountID: account.ID,
					SessionID: session.SessionID,
				})
			}
		}
	}
}

func Logout(c *gin.Context) {

}
