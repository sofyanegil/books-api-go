package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Book represents the model for an book
type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null;type:varchar(200)" json:"name_book"`
	Author    string    `gorm:"not null;type:varchar(200)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Book) TableName() string {
	return "book"
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {

	if len(b.Title) <= 2 {
		err = errors.New("book title is too short")
	}

	if len(b.Author) <= 2 {
		err = errors.New("author is too short")
	}

	return
}
