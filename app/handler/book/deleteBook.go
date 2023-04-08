package book

import (
	repo "books-api/app/repository/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteBooks godoc
// @Summary Delete book identified by the given bookID
// @Description Delete the order corresponding to the input bookID
// @Tags books
// @Accept json
// @Produces json
// @Param bookID path int true "bookID of the book to be deleted"
// @Success 204 "No Content"
// @Router /books/{bookID} [delete]
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
