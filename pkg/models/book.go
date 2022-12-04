package models

import (
	"time"

	"github.com/gareisdev/go-books-restapi/pkg/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"column:title" json:"title"`
	Author      string    `gorm:"column:author" json:"author"`
	Publication string    `gorm:"column:publication" json:"publication"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime:true"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:true"`
}

func init() {
	time.Sleep(5 * time.Second)
	config.ConnectDB()
	db = config.GetConnectionDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book

	db.Find(&books)

	return books
}

func GetBook(id int64) (Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", id).First(&book)

	return book, db
}

func DeleteBook(id int64) Book {
	var books []Book
	db.Clauses(clause.Returning{}).Where("ID=?", id).Delete(&books)

	if len(books) > 1 {
		return books[0]
	} else {
		return Book{}
	}
}
