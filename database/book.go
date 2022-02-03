package database

import (
	"github.com/adityasinghtads/api_go/models"
	"github.com/jinzhu/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book, error) {

	books := []models.Book{}
	query := db.Select("books.*")
	err := query.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
