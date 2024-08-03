package handler

import (
	"Book/internal/models"
	"Book/internal/service"
	"net/http"
	"strconv"
	_ "Book/docs"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept  json
// @Produce  json
// @Param Book body models.Book true "Book to create"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books [post]
func (b *BookHandler) CreateBooks(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	Book, err := b.Service.CreateBooks(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": Book.Id})
}

// GetBookById godoc
// @Summary Get a book by ID
// @Description Get a book by its ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [get]
func (b *BookHandler) GetbyIdBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	book, err := b.Service.GetbyIdBooks(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Produce  json
// @Success 200 {array} models.Book
// @Failure 500 {object} string
// @Router /books [get]
func (b *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := b.Service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param Book body models.Book true "Book to update"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [put]
func (b *BookHandler) UpdateBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	book.Id = id
	
	err = b.Service.UpdateBooks(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [delete]
func (b *BookHandler) DeleteBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = b.Service.DeleteBooks(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
