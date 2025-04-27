package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"INF368/internal/models"
	"INF368/internal/service"
	_ "INF368/docs" 
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/books", h.addBook)
	router.GET("/books", h.getAllBooks)
	router.DELETE("/books/:id", h.deleteBook)

	return router
}

// @Summary Create a new book
// @Description This endpoint creates a new book with a given title.
// @Tags books
// @Accept json
// @Produce json
// @Param title body models.Input true "Book title"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /books [post]
func (h *BookHandler) addBook(c *gin.Context) {
	var input models.Input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "400",
					Message: "Invalid input data",
				},
			},
		})
		return
	}

	h.service.AddBook(input.Title)

	c.JSON(http.StatusOK, models.SuccessResponse{
		Result: gin.H{"message": "Book created successfully"},
	})
}

// getAllBooks handles the HTTP GET request to retrieve all books.
// @Summary Get all books
// @Description This endpoint retrieves all books in the system.
// @Tags books
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func (h *BookHandler) getAllBooks(c *gin.Context) {
	books := h.service.GetAllBooks()

	c.JSON(http.StatusOK, models.SuccessResponse{
		Result: books,
	})
}

// deleteBook handles the HTTP DELETE request to delete a book by its ID.
// @Summary Delete a book
// @Description This endpoint deletes a book by its ID.
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) deleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "400",
					Message: "Invalid ID format",
				},
			},
		})
		return
	}

	if deleted := h.service.DeleteBook(id); !deleted {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "404",
					Message: "Book not found",
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Result: gin.H{"status": "book deleted"},
	})
}
