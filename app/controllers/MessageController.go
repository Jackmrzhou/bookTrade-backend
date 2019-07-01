package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type contactWithUser struct {
	ContactID int `json:"contact_id"`
	UserID int `json:"user_id"`
	AvatarKey string `json:"avatar_key"`
	Username string `json:"username"`
	Unread string `json:"unread"`
}

type getContactResp struct {
	Contacts []contactWithUser `json:"contacts"`
}

func GetContact(c *gin.Context) {
	if p, err := utils.GetPrinciple(c); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		if cs, err := dao.GetAllContactsBySelfID(p.UserId); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			var resp getContactResp
			for _, contact := range cs {
				if user, err := dao.GetUserByUserID(contact.CounterpartID); err != nil {
					c.Set(error2.CodeKey, error2.ServerError)
					return
				}else if count, err := dao.CountMessageUnreadByContactID(contact.ID); err != nil{
					c.Set(error2.CodeKey, error2.ServerError)
					return
				} else {
					resp.Contacts = append(resp.Contacts, contactWithUser{
						ContactID:contact.ID,
						UserID:user.ID,
						AvatarKey:user.AvatarKey,
						Username:user.Username,
						Unread: strconv.Itoa(count),
					})
				}
			}
			c.JSON(http.StatusOK, resp)
		}
	}
}

type createContactReq struct {
	UserID int `json:"user_id" binding:"required"`
}

type createContactResp struct {
	ContactID int `json:"contact_id"`
}

func CreateContact(c *gin.Context) {
	var query createContactReq
	if err := c.ShouldBindJSON(&query); err != nil{
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	if p, err := utils.GetPrinciple(c); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	}else {
		contact := models.Contact{
			SelfID:p.UserId,
			CounterpartID:query.UserID,
		}

		if err := dao.CreateContactIfNotExists(&contact); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			c.JSON(http.StatusOK, createContactResp{
				ContactID:contact.ID,
			})
		}
	}
}

type getAllMessageReq struct {
	ContactID int `form:"contact_id"`
}

type getAllMessageResp struct {
	Messages []models.Message `json:"messages"`
}

func GetAllMessage(c *gin.Context) {
	var query getAllMessageReq
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}
	if msgs, err := dao.GetAllMessageByContactID(query.ContactID); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		var unreadMsgs []models.Message
		for _, msg := range msgs{
			msg.IsRead = models.MESSAGE_READ
			unreadMsgs = append(unreadMsgs, msg)
		}
		if err := dao.SaveAllMsgs(unreadMsgs); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
			return
		}
		c.JSON(http.StatusOK, getAllMessageResp{
			Messages:msgs,
		})
	}
}

type sendMessageReq struct {
	ContactID int `json:"contact_id" binding:"required"`
	CounterpartID int `json:"counterpart_id" bind:"required"`
	Content string `json:"content" binding:"required"`
}

type sendMessageResp struct {
	MessageID int `json:"message_id"`
}

func SendMessage(c *gin.Context) {
	var query sendMessageReq
	if err := c.ShouldBindJSON(&query); err != nil {
		logrus.WithError(err).Info("bind json failed")
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	if p, err := utils.GetPrinciple(c); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		msg := models.Message{
			ContactID:query.ContactID,
			FromID:p.UserId,
			ToID:query.CounterpartID,
			Content:query.Content,
			CreateTime:time.Now(),
			IsRead:models.MESSAGE_UNREAD,
		}
		if err := dao.CreateMessage(&msg); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			c.JSON(http.StatusOK, sendMessageResp{
				MessageID:msg.ID,
			})
		}
	}
}

type getUnreadMsgCountReq struct {
	ContactID int `form:"contact_id" binding:"required"`
}

type getUnreadMsgCountResp struct {
	Count int `json:"count"`
}

func GetUnreadMsgCount(c *gin.Context) {
	var query getUnreadMsgCountReq
	if err := c.ShouldBindQuery(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	var p utils.Principle
	var err error
	if p, err = utils.GetPrinciple(c); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}

	if query.ContactID == -1 {
		// get all unread message
		if count, err := dao.CountMessageUnread(p.UserId); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
			return
		}else {
			c.JSON(http.StatusOK, getUnreadMsgCountResp{
				Count: count,
			})
		}
	} else {
		if count, err := dao.CountMessageUnreadByContactID(query.ContactID); err != nil {
			c.Set(error2.CodeKey, error2.ServerError)
			return
		} else {
			c.JSON(http.StatusOK, getUnreadMsgCountResp{
				Count:count,
			})
		}
	}
}