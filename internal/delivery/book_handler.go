package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/models"
	service "rest-project/internal/services"
	"strconv"
)

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

type BookHandler struct {
	service *service.BookService
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, _ := h.service.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var bookCreate models.BookEdit

	if err := c.ShouldBindJSON(&bookCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newBook, err := h.service.Create(bookCreate.Title, bookCreate.Author, bookCreate.PublishedAt, bookCreate.Pages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var bookEdit models.BookEdit
	if err := c.ShouldBindJSON(&bookEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedBook, err := h.service.Update(id, &bookEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.service.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
