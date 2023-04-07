package book

import (
	"books-api/app/entity"
	"books-api/pkg/database"
)

func CreateBook(book entity.Book) (result entity.Book, err error) {
	db, err := database.ConnectDB()
	if err != nil {
		return
	}
	book = entity.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	err = db.Create(&book).Error

	if err != nil {
		return
	}

	result = book
	return
}
