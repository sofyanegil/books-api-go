package book

import (
	"books-api/app/entity"
	repo "books-api/app/repository/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookByIdHandler(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var book entity.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid bookID",
		})
		return
	}

	result, err := repo.GetBookById(idBook, book)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
