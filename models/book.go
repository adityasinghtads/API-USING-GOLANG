package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Author    string
	Name      string
	PageCount int
}
