package middlewares

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/dao"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("sessionID")
		if err != nil {
			if c.Request.URL.Path == "/api/passport/ping" {
				c.Set(error2.CodeKey, error2.PingUnauth)
			} else {
				c.Set(error2.CodeKey, error2.Unauth)
			}
			c.Abort()
			return
		} else if ok, err := utils.ValidateSessionID(sessionID); err != nil {
			c.Set(error2.CodeKey, error2.ValidateSession)
			c.Abort()
			return
		} else if ok {
			s, err := dao.GetSessionBySessionID(sessionID)
			if err == nil {
				// service down grade
				if user, e := dao.GetUserByAccountID(s.AccountID); e == nil {
					c.Set("principle", utils.Principle{
						AccountID: s.AccountID,
						UserId:user.ID,
					})
				} else if c.Request.URL.Path == "/api/user/" {
					c.Set("principle", utils.Principle{
						AccountID: s.AccountID,
						UserId:0,
					})
				} else {
					logrus.WithError(e).Error("get user by accountID failed")
				}
			} else {
				logrus.WithError(err).Warn("get session in database failed")
			}
			c.Set("sessionID", sessionID)
			c.Next()
			return
		} else {
			c.Set(error2.CodeKey, error2.Unauth)
			c.Abort()
			return
		}
	}
}
