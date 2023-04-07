package book

import (
	"books-api/app/entity"
	"books-api/pkg/database"
	"errors"

	"gorm.io/gorm"
)

func GetBookById(id int, book entity.Book) (result entity.Book, err error) {
	db, err := database.ConnectDB()
	if err != nil {
		return
	}
	err = db.First(&book, "id = ?", id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	result = book
	return
}
