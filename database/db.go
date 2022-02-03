package database

import (
	"log"

	"github.com/adityasinghtads/api_go/models"
	"github.com/jinzhu/gorm"
)

// gorm can be used with postgres mysql orace mssql

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "postgres"
	args := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{}) // create table autoatically in the DB
	DB = db
}
