package app

import (
	"books-api/app/handler/book"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	route.POST("/books", book.CreateBookHandler)
	route.GET("/books", book.GetAllBookHandler)
	route.GET("/books/:bookID", book.GetBookByIdHandler)
	route.PUT("/books/:bookID", book.UpdateBookHandler)
	route.DELETE("/books/:bookID", book.DeleteBookHandler)

	return route
}
