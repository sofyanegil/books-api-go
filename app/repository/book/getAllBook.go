package book

import (
	"books-api/app/entity"
	"books-api/pkg/database"
	"errors"
	"log"

	"gorm.io/gorm"
)

func GetAllBook(books []entity.Book) (result []entity.Book, err error) {
	db, err := database.ConnectDB()
	if err != nil {
		return
	}
	err = db.Find(&books).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Book data not found")
		err = errors.New("error getting data")
		return
	}

	result = append(result, books...)
	return
}
