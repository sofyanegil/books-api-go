package book

import (
	"books-api/app/entity"
	repo "books-api/app/repository/book"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBookHandler(ctx *gin.Context) {
	var book entity.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := repo.CreateBook(book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprint("Failed add book because ", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
