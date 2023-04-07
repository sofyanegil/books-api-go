package book

import (
	repo "books-api/app/repository/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteBookHandler(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid bookID",
		})
		return
	}

	err = repo.DeleteBook(idBook)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book with id %d not found", idBook),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %d successfully deleted", idBook),
	})
}
