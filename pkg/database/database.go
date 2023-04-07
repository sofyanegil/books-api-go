package database

import (
	"books-api/app/entity"
	"books-api/pkg/common"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", common.DBHost, common.DBUser, common.DBPassword, common.DBName, common.DBPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.AutoMigrate(entity.Book{})

	return db, nil
}
