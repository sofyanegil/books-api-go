package app

import (
	"books-api/app/handler/book"
	_ "books-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title Book API
//
// @version 1.0
// @description This is a sample service for managing books
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email sofyanegil@gmail.com
// @license.name Apache 2.0
// @licence.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func StartServer() *gin.Engine {
	route := gin.Default()
	// Create
	route.POST("/books", book.CreateBookHandler)
	// Read All
	route.GET("/books", book.GetAllBookHandler)
	// Read
	route.GET("/books/:bookID", book.GetBookByIdHandler)
	// Update
	route.PUT("/books/:bookID", book.UpdateBookHandler)
	// Delete
	route.DELETE("/books/:bookID", book.DeleteBookHandler)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return route
}
