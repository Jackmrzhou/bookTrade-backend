package models

import "time"

type Order struct {
	ID int `gorm:"primary_key"`
	BookID int
	TransportType int
	OrderType int
	SellerID int
	BuyerID int
	Status int
	CreateTime time.Time
}

func (Order) TableName() string {
	return "book_order"
}

const (
	TRANSPORT_OFFLINE = 1
	TRANSPORT_SEND = 2
)

const (
	ORDERTYPE_SELL = 1
	ORDERTYPE_REQUEST = 2
)

const (
	ORDER_DONE = 1
	ORDER_UNDONE = 0
)
