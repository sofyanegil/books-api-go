package book

import (
	"books-api/app/entity"
	repo "books-api/app/repository/book"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
