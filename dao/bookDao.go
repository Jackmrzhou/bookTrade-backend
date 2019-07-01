package dao

import "bookTrade-backend/models"

func CreateBook(book *models.Book) error {
	return db.Create(book).Error
}

func GetBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	err := db.Where("id = ?", bookID).First(&book).Error
	return &book, err
}

func GetBooksByCatalogID(start, limit, catalogID, Type int) ([]models.Book, error) {
	var books []models.Book
	err := db.Offset(start).Limit(limit).Where("catalog_id = ? AND type = ?", catalogID, Type).Find(&books).Error
	return books, err
}

func GetBooks(start, limit, Type int) ([]models.Book, error) {
	var books []models.Book
	err := db.Offset(start).Limit(limit).Where("type = ?", Type).Find(&books).Error
	return books, err
}