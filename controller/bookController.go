package controller

import (
	"Belajar-Golang-7/dto"
	"Belajar-Golang-7/mapper"
	"Belajar-Golang-7/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertBook(ctx *gin.Context) {
	var newBookDTO dto.BookDTO

	if err := ctx.ShouldBindJSON(&newBookDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Please provide a valid data",
		})
		return
	}

	newBookEntity := mapper.BookDtoToNewEntity(newBookDTO)
	if err := service.InsertBook(newBookEntity); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
	})
}

func GetAllBook(ctx *gin.Context) {
	var bookEntities = service.GetAllBook()
	bookDTO := []dto.BookDTO{}

	for _, entity := range bookEntities {
		bookDTO = append(bookDTO, mapper.BookEntityToDto(entity))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    bookDTO,
	})
}

func GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Please provide a valid ID",
		})
		return
	}

	book, err := service.GetBookByID(uint(bookID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Book not found",
		})
		return
	}

	bookDTO := mapper.BookEntityToDto(book)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"data":    bookDTO,
	})
}

func UpdateBook(ctx *gin.Context) {
	var updatedBookDTO dto.BookDTO

	if err := ctx.ShouldBindJSON(&updatedBookDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Please provide a valid data",
		})
		return
	}

	updatedBookEntity, err := mapper.BookDtoToEntity(updatedBookDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Please provide a valid ID",
		})
		return
	}

	if err := service.UpdateBook(updatedBookEntity); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Book not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
	})
}

func DeleteBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Please provide a valid ID",
		})
		return
	}

	if err := service.DeleteBookByID(uint(bookID)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Book not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
	})
}
