package models

import (
	"bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (new_book *Book) CreateBook() *Book {
	db.NewRecord(new_book)
	db.Create(&new_book)
	return new_book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
