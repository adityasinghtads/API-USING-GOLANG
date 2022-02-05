package database

import (
	"github.com/adityasinghtads/api_go/models"
	"github.com/jinzhu/gorm"
)

//getBooks is creating connection and interacting from golang app to db server via db variable
func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func GetBookByID(db *gorm.DB, id string) (*models.Book, error) {
	book := models.Book{}
	err := db.Select("book.*").Group("books.id").Where("books.id= ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func DeleteBookByID(db *gorm.DB, id string) error {
	var book models.Book
	err := db.Where("id=?", id).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateBookByID(db *gorm.DB, book *models.Book) error {

	err := db.Save(book).Error // Book is with id
	if err != nil {
		return err
	}
	return nil
}
func SaveBook(db *gorm.DB, book *models.Book) error {
	err := db.Save(book).Error // book here is without ID
	if err != nil {
		return err
	}
	return nil
}

// Add the handler and router..
