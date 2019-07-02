package controllers

import (
	error2 "bookTrade-backend/app/error"
	"bookTrade-backend/dao"
	"bookTrade-backend/models"
	"bookTrade-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type createOrderReq struct {
	BookID int `json:"book_id" binding:"required"`
	TransportType int `json:"transport_type" binding:"required"`
	OrderType int `json:"order_type" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
}

type createOrderResp struct {
	OrderID int `json:"order_id"`
}

func CreateOrder(c *gin.Context) {
	var query createOrderReq
	if err := c.ShouldBindJSON(&query); err != nil {
		logrus.WithError(err).Info("binding failed")
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	var principle utils.Principle
	var err error
	if principle, err = utils.GetPrinciple(c); err != nil {
		logrus.WithError(err).Error("get principle failed")
		c.Set(error2.CodeKey, error2.ServerError)
		return
	}

	if principle.UserId == query.UserID {
		c.Set(error2.CodeKey, error2.SelfOrder)
		return
	}

	if _, err := dao.GetOrderByBookID(query.BookID); err == nil || !gorm.IsRecordNotFoundError(err) {
		c.Set(error2.CodeKey, error2.BookOrdered)
		return
	}

	order := models.Order{
		BookID:query.BookID,
		TransportType:query.TransportType,
		OrderType:query.OrderType,
		Status:models.ORDER_UNDONE,
		CreateTime:time.Now(),
	}

	if query.OrderType == models.ORDERTYPE_SELL {
		order.SellerID = query.UserID
		order.BuyerID = principle.UserId
	} else {
		order.SellerID = principle.UserId
		order.BuyerID = query.UserID
	}

	if err := dao.DeleteBookByID(query.BookID); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {

		if err = dao.CreateOrder(&order); err != nil {
			logrus.WithError(err).Error("create order failed")
			dao.RecoverBookByID(query.BookID)
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			c.JSON(http.StatusOK, createOrderResp{
				OrderID: order.ID,
			})
		}
	}

}

type typicalOrder struct {
	OrderID int `json:"order_id"`
	CreateTime string `json:"create_time"`
	OrderType int `json:"order_type"`
	BookID int `json:"book_id"`
	BookName string `json:"book_name"`
	Status int `json:"status"`
	TransportType int `json:"transport_type"`
}

type getOrdersResp struct {
	Orders []typicalOrder `json:"orders"`
}

func GetOrders(c *gin.Context) {
	if p, err := utils.GetPrinciple(c); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		if orders, err := dao.GetOrdersByUserID(p.UserId); err != nil {
			logrus.WithError(err).Error("get orders failed")
			c.Set(error2.CodeKey, error2.ServerError)
		} else {
			var resp getOrdersResp
			for _, order := range orders {

				if book, err := dao.GetBookByIDDeletedOrNot(order.BookID); err != nil {
					c.Set(error2.CodeKey, error2.ServerError)
					return
				} else {

					resp.Orders = append(resp.Orders, typicalOrder{
						OrderID:    order.ID,
						CreateTime: order.CreateTime.Format("2006-01-02"),
						BookID:book.ID,
						OrderType:  order.OrderType,
						BookName: book.Name,
						Status: order.Status,
						TransportType:order.TransportType,
					})
				}
			}
			c.JSON(http.StatusOK, resp)
		}
	}
}

type updateOrderReq struct {
	OrderID int `json:"order_id"`
	Status int `json:"status"`
}

type updateOrderResp struct {
	OrderID	int `json:"order_id"`
}

func UpdateOrder(c *gin.Context) {
	var query updateOrderReq
	if err := c.ShouldBindJSON(&query); err != nil {
		c.Set(error2.CodeKey, error2.BadRequest)
		return
	}

	if err := dao.UpdateOrderStatus(query.OrderID, query.Status); err != nil {
		c.Set(error2.CodeKey, error2.ServerError)
	} else {
		c.JSON(http.StatusOK, updateOrderResp{
			OrderID:query.OrderID,
		})
	}
}