package book

import (
	"books-api/app/entity"
	"books-api/pkg/database"
	"errors"

	"gorm.io/gorm"
)

func DeleteBook(id int) (err error) {
	db, err := database.ConnectDB()
	if err != nil {
		return
	}
	var book entity.Book

	err = db.Where("id = ?", id).First(&book).Delete(&book).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return
}
