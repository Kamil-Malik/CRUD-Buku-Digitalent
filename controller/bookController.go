package controller

import (
	"Belajar-Golang-7/dto"
	"Belajar-Golang-7/mapper"
	"Belajar-Golang-7/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// insert book godoc
// @Summary      Insert a Book
// @Description  insert a new Book into postgres database
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GenericResponseDTO
// @Failure      400  {object}  dto.GenericResponseDTO
// @Router       /books [post]
func InsertBook(ctx *gin.Context) {
	var newBookDTO dto.BookDTO

	if err := ctx.ShouldBindJSON(&newBookDTO); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid data",
			})
		return
	}

	newBookEntity := mapper.BookDtoToNewEntity(newBookDTO)
	if err := service.InsertBook(newBookEntity); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		dto.GenericResponseDTO{
			Code:    http.StatusOK,
			Message: "OK",
		})
}

// Get all Books godoc
// @Summary      Get all books
// @Description  get all books from postgres database
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.BookDTO
// @Failure      400  {object}  dto.GenericResponseDTO
// @Router       /books [get]
func GetAllBook(ctx *gin.Context) {
	var bookEntities = service.GetAllBook()
	bookDTO := []dto.BookDTO{}

	for _, entity := range bookEntities {
		bookDTO = append(bookDTO, mapper.BookEntityToDto(entity))
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status": dto.GenericResponseDTO{
				Code:    http.StatusOK,
				Message: "OK",
			},
			"data": bookDTO,
		})
}

// Get a book by ID godoc
// @Summary      Get Book by ID
// @Description  get a book by its ID from postgres database
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  dto.BookDTO
// @Failure      400  {object}  dto.GenericResponseDTO
// @Failure      404  {object}  dto.GenericResponseDTO
// @Router       /books/{id} [get]
func GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Please provide a valid ID",
			})
		return
	}

	book, err := service.GetBookByID(uint(bookID))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Book not found",
			})
		return
	}

	bookDTO := mapper.BookEntityToDto(book)
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"Status": dto.GenericResponseDTO{
				Code:    http.StatusOK,
				Message: "OK",
			},
			"data": bookDTO,
		})
}

// Update a book godoc
// @Summary      Update a book data
// @Description  update a book from postgres database
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GenericResponseDTO
// @Failure      400  {object}  dto.GenericResponseDTO
// @Failure      404  {object}  dto.GenericResponseDTO
// @Router       /books [put]
func UpdateBook(ctx *gin.Context) {
	var updatedBookDTO dto.BookDTO

	if err := ctx.ShouldBindJSON(&updatedBookDTO); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid data",
			})
		return
	}

	updatedBookEntity, err := mapper.BookDtoToEntity(updatedBookDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid data",
			})
		return
	}

	if err := service.UpdateBook(updatedBookEntity); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Book not found",
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		dto.GenericResponseDTO{
			Code:    http.StatusOK,
			Message: "OK",
		})
}

// Delete a book by ID godoc
// @Summary      Delete Book by ID
// @Description  delete a book by its ID from postgres database
// @Tags         books
// @Produce      json
// @Param        id   path      int  true  "Book ID"
// @Success      200  {object}  dto.GenericResponseDTO
// @Failure      400  {object}  dto.GenericResponseDTO
// @Failure      404  {object}  dto.GenericResponseDTO
// @Router       /books/{id} [delete]
func DeleteBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			dto.GenericResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "Please provide a valid ID",
			})
		return
	}

	if err := service.DeleteBookByID(uint(bookID)); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			dto.GenericResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Book not found",
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		dto.GenericResponseDTO{
			Code:    http.StatusOK,
			Message: "OK",
		})
}
