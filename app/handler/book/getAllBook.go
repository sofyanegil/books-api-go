package book

import (
	"books-api/app/entity"
	repo "books-api/app/repository/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBookHandler(ctx *gin.Context) {
	var books []entity.Book

	result, err := repo.GetAllBook(books)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid Request %+v", err.Error()),
		})
		return
	}

	if len(result) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Books empty",
			"data":    []entity.Book{},
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
