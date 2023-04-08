package book

import (
	"books-api/app/entity"
	repo "books-api/app/repository/book"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateBooks godoc
// @Summary Update book identified by the given Id
// @Description Update the book corresponding to the input bookID
// @Tags books
// @Accept json
// @Produces json
// @Param bookID path int true "bookID of the book to be updated"
// @Param entity.Book body entity.Book true "update book"
// @Success 200 {object} entity.Book
// @Router /books/{bookID} [put]
func UpdateBookHandler(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var book entity.Book

	idBook, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid bookID",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := repo.UpdateBook(idBook, book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid update data because %+v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
