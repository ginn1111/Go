package models

import (
	"log"

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

func (Book) TableName() string {
	return "books"
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
func UpdateBookById(bookId uint, book *Book) {
	updatingBook := Book{}
	updatingBook.ID = bookId

	db.First(&updatingBook)

	if book.Author != "" {
		updatingBook.Author = book.Author
	}
	if book.Name != "" {
		updatingBook.Name = book.Name
	}
	if book.Publication != "" {
		updatingBook.Publication = book.Publication
	}

	log.Println(updatingBook.ID, bookId)

	db.Save(&updatingBook)
}
