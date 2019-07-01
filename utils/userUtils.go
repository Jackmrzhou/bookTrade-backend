package utils

import (
	"bookTrade-backend/dao"
	"errors"
	"github.com/gin-gonic/gin"
)

type Principle struct {
	UserId int
	AccountID int
}

func GetPrinciple(c *gin.Context) (Principle, error) {
	if p, ok := c.Get("principle"); ok {
		return p.(Principle), nil
	} else if session, err := dao.GetSessionBySessionID(c.GetString("sessionID")); err != nil {
		return Principle{}, errors.New("principle not found")
	} else if user, e := dao.GetUserByAccountID(session.AccountID); e != nil{
		return Principle{}, errors.New("principle not found")
	} else {
		principle := Principle{
			AccountID: session.AccountID,
			UserId:user.ID,
		}
		c.Set("principle", principle)
		return principle, nil
	}
}
