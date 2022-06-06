package models

import (
	"fmt"
	"github.com/leehaowei/go-dbapp/src/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Price  int    `json:"price"`
	Author string `json:"author"`
}

func init() {
	fmt.Println("models.book.init()")
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
