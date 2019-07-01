package dao

import "bookTrade-backend/models"

func CreateBook(book *models.Book) error {
	return db.Create(book).Error
}

func GetBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	err := db.Where("id = ? AND is_deleted = ?", bookID, models.BOOK_UNDELETED).First(&book).Error
	return &book, err
}

func GetBooksByCatalogID(start, limit, catalogID, Type int) ([]models.Book, error) {
	var books []models.Book
	err := db.Offset(start).Limit(limit).Where("catalog_id = ? AND type = ? AND is_deleted = ?",
		catalogID, Type, models.BOOK_UNDELETED).Find(&books).Error
	return books, err
}

func GetBooks(start, limit, Type int) ([]models.Book, error) {
	var books []models.Book
	err := db.Offset(start).Limit(limit).Where("type = ? AND is_deleted = ?", Type, models.BOOK_UNDELETED).Find(&books).Error
	return books, err
}

func GetBookByIDDeletedOrNot(bookID int) (*models.Book, error) {
	var book models.Book
	err := db.Where("id = ?", bookID).First(&book).Error
	return &book, err
}

func DeleteBookByID(bookID int) error {
	return db.Model(&models.Book{}).Where("id = ?", bookID).Update("is_deleted", models.BOOK_DELETED).Error
}

func RecoverBookByID(bookID int) error {
	return db.Model(&models.Book{}).Where("id = ?", bookID).Update("is_deleted", models.BOOK_UNDELETED).Error
}