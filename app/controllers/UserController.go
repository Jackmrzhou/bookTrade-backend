package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/app/services"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"bookTrade-backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type updateUserReq struct {
	Username     string `json:"username" binding:"required"`
	AvatarKey    string `json:"avatar_key" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address string 	`json:"address" binding:"required"`
}

type updateUserResp struct {
	Message string `json:"message"`
}

func UpdateUser(c *gin.Context) {
	var query updateUserReq
	if err := c.ShouldBindJSON(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	var accountID int
	if p, err := utils.GetPrinciple(c); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
		return
	} else {
		accountID = p.AccountID
	}

	if ok, err := services.TestUsername(query.Username, accountID); err != nil {
		log.WithError(err).Error("test username failed")
		c.Set(error2.CodeKey, error2.ServerError)
	} else if ok {
		// username exists
		c.Set(error2.CodeKey, error2.UsernameDup)
	} else {
		user := models.User{
			AccountID:    accountID,
			Username:     query.Username,
			AvatarKey:    query.AvatarKey,
			Introduction: query.Introduction,
			PhoneNumber:query.PhoneNumber,
			Address:query.Address,
		}
		// create if not exists
		if err := dao.UpdateOrCreateUser(&user); err != nil {
			log.WithError(err).Error("update user failed")
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			c.JSON(http.StatusOK, updateUserResp{
				Message: "update successful",
			})
		}
	}
}

type getAvatarResp struct {
	AvatarKey string `json:"avatar_key"`
}

func GetAvatar(c *gin.Context) {
	if p, err := utils.GetPrinciple(c); err != nil {
		log.WithError(err).Error("get principle failed")
		c.Set(error2.CodeKey, error2.ServerError)
		return
	} else {
		if user, err := dao.GetUserByAccountID(p.AccountID); err != nil {
			log.WithError(err).Error("get user by account id failed")
			c.Set(error2.CodeKey, error2.ServerError)
			return
		} else {
			c.JSON(http.StatusOK, getAvatarResp{
				AvatarKey: user.AvatarKey,
			})
		}
	}
}

type getUserProfileResp struct {
	Username     string `json:"username"`
	AvatarKey    string `json:"avatar_key"`
	Introduction string `json:"introduction"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
}

func GetUserProfile(c *gin.Context) {
	if p, err := utils.GetPrinciple(c); err != nil {
		log.WithError(err).Error("get principle failed")
		c.Set(error2.CodeKey, error2.ServerError)
		return
	} else {
		if user, err := dao.GetUserByAccountID(p.AccountID); err != nil {
			log.WithError(err).Error("get user by account id failed")
			c.Set(error2.CodeKey, error2.ServerError)
			return
		} else {
			c.JSON(http.StatusOK, getUserProfileResp{
				Username:     user.Username,
				AvatarKey:    user.AvatarKey,
				Introduction: user.Introduction,
				PhoneNumber:user.PhoneNumber,
				Address:user.Address,
			})
		}
	}
}

type getUserProdileCommonReq struct {
	UserID int `form:"user_id"`
}

func GetUserProfileCommon(c *gin.Context) {
	var query getUserProdileCommonReq
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}
	if user, err := dao.GetUserByUserID(query.UserID); err != nil {
		log.WithError(err).Error("get user by user id failed")
		c.Set(error2.CodeKey, error2.ServerError)
		return
	} else {
		c.JSON(http.StatusOK, getUserProfileResp{
			Username:     user.Username,
			AvatarKey:    user.AvatarKey,
			Introduction: user.Introduction,
			PhoneNumber:user.PhoneNumber,
			Address:user.Address,
		})
	}
}

type getViewedUsersReq struct {
	BookID int `form:"book_id" binding:"required"`
}

type getViewedUserItem struct {
	UserID int `json:"user_id"`
	AvatarKey string `json:"avatar_key"`
	Username string `json:"username"`
	ViewedTime string `json:"viewed_time"`
}

type getViewedUserResp struct {
	Users []getViewedUserItem `json:"users"`
}

const ReturnLimit = 5

func GetViewedUsers(c *gin.Context) {
	var query getViewedUsersReq
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	if rs, err := dao.GetRecordsByBookID(query.BookID); err != nil {
		log.WithError(err).Error("get records failed")
		c.Set(error2.CodeKey, error2.ServerError)
	}else {
		if len(rs) >= ReturnLimit{
			rs = rs[0:ReturnLimit]
		}
		var resp getViewedUserResp
		for _, r := range rs {
			if u, err := dao.GetUserByUserID(r.UserID); err != nil {
				log.WithError(err).Error("get user by user id failed")
				c.Set(error2.CodeKey, error2.ServerError)
				return
			}else {
				diff := time.Now().Sub(r.ViewedTime).Minutes()
				resp.Users = append(resp.Users, getViewedUserItem{
					UserID: u.ID,
					AvatarKey:u.AvatarKey,
					Username:u.Username,
					ViewedTime:fmt.Sprintf("%.0f", diff),
				})
			}
		}

		c.JSON(http.StatusOK, resp)
	}
}