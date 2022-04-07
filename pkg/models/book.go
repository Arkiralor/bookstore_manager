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

func GetBookByID(book_Id int64) (*Book, *gorm.DB) {
	var found_book Book
	db := db.Where("ID=?", book_Id)

	return &found_book, db
}

func DeleteBook(book_ID int64) *Book {
	var deleted_book Book
	db.Where("ID=?", book_ID).Delete(deleted_book)

	return &deleted_book
}
