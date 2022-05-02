package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tobioye88/go-tutorial-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(&b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) *Book {
	var book Book
	db.Where("Id = ?", Id).Find(&book)
	return &book
}

func DeleteBook(Id int64) {
	db.Delete(Book{}, "Id = ?", Id)
}

func UpdateBook(Id int64, bookModel *Book) *Book {
	var book Book
	db.Where("Id = ?", Id).Find(&book)
	db.Model(&book).Update(bookModel)
	return &book
}
