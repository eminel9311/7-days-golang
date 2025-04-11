package controller

import (
	"go-bookstore/internal/service"
	"net/http"

	"go-bookstore/pkg/response"

	"go-bookstore/internal/model"

	"github.com/gin-gonic/gin"
)

// BookController handles HTTP requests related to books
type BookController struct {
	bookService service.BookService
}

// NewBookController creates a new BookController instance
func NewBookController(bookService service.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

// GetAllBooks handles GET request to fetch all books
func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.bookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error("Failed to fetch books", err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success("Books retrieved successfully", books))
}

// GetBookByID handles GET request to fetch a book by ID
func (c *BookController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	book, err := c.bookService.GetBookByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Error("Book not found", err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success("Book retrieved successfully", book))
}

// CreateBook handles POST request to create a new book
func (c *BookController) CreateBook(ctx *gin.Context) {
	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error("Invalid input", err))
		return
	}

	createdBook, err := c.bookService.CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error("Failed to create book", err))
		return
	}

	ctx.JSON(http.StatusCreated, response.Success("Book created successfully", createdBook))
}

// UpdateBook handles PUT request to update an existing book
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error("Invalid input", err))
		return
	}

	updatedBook, err := c.bookService.UpdateBook(id, book)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Error("Failed to update book", err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success("Book updated successfully", updatedBook))
}

// DeleteBook handles DELETE request to remove a book
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.bookService.DeleteBook(id); err != nil {
		ctx.JSON(http.StatusNotFound, response.Error("Failed to delete book", err))
		return
	}

	ctx.JSON(http.StatusOK, response.Success("Book deleted successfully", nil))
}
