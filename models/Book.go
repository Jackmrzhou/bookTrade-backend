package models

type Book struct {
	ID           int `gorm:"primary_key"`
	UserID  int
	Name         string
	Author       string
	ISBN         string
	Price        float32
	CoverKey     string
	Introduction string
	Type         int    // for a sale or for a request
	OutLink      string // only amazon
	CatalogID    int
}

func (Book) TableName() string {
	return "book"
}

const (
	SELL = 1
	REQUEST = 1 << 1
)

const OUTLINK = "https://isbnsearch.org/isbn/"