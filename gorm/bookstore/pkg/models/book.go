package models

import (
	"github.com/go/gorm/bookstore/pkg/config"
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

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBookById(bookId int64) *Book {
	var book Book
	db.Where("ID=?", bookId).Find(&book)

	return &book
}

func DeleteBookById(bookId int64) Book {
	var book Book
	db.Where("ID=?", bookId).Delete(book)

	return book
}

// my implementation
func UpdateBookById(bookId int64, book *Book) {
	db.Where("ID=?", bookId).Update(book)
}
