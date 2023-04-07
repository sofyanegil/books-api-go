package book

import (
	"books-api/app/entity"
	"books-api/pkg/database"
	"errors"
	"log"

	"gorm.io/gorm"
)

func UpdateBook(id int, book entity.Book) (result entity.Book, err error) {
	db, err := database.ConnectDB()
	if err != nil {
		return
	}

	var findBook entity.Book
	err = db.Where("id = ?", id).First(&findBook).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book not found")
		err = errors.New("book not found")
		return
	}

	err = db.Model(&book).Where("id = ?", id).Updates(entity.Book{
		ID:        uint(id),
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: findBook.CreatedAt,
	}).Error

	if err != nil {
		log.Println("Error updating book data", err)
		err = errors.New("error updating book data")
		return
	}

	result = book
	return
}
