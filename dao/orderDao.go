package dao

import (
	"bookTrade-backend/models"
	models2 "github.com/jackmrzhou/bookTrade-backend/models"
)

func CreateOrder(order *models.Order) error {
	return db.Create(order).Error
}

func GetOrdersByUserID(userID int) ([]models.Order, error) {
	var orders []models.Order
	err := db.Where("buyer_id = ? OR seller_id = ?", userID, userID).Find(&orders).Error
	return orders, err
}

func UpdateOrderStatus(orderID, status int) error{
	return db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func GetOrderByBookID(bookID int) (models.Order, error) {
	var order models.Order
	err := db.Where("book_id = ?", bookID).First(&order).Error
	return order, err
}