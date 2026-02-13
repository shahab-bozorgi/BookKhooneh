package handlers

import (
	"BookKhoone/internal/dto"
	"BookKhoone/internal/models"
	"BookKhoone/internal/services"
	"BookKhoone/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// CreateBookHandler godoc
// @Summary Create a new book
// @Description Create a new book (requires authentication)
// @Tags Books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body dto.CreateBookRequest true "Book data"
// @Success 200 {object} dto.CreateBookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 401 {object} dto.BookErrorResponse
// @Router /books/create [post]
// @Security ApiKeyAuth
func CreateBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDValue, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, dto.BookErrorResponse{Error: "User not found in context"})
			return
		}

		userID, ok := userIDValue.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{Error: "Invalid user ID format"})
			return
		}

		var input dto.CreateBookRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, dto.BookErrorResponse{Error: err.Error()})
			return
		}

		userIDPtr := userID
		book := models.Book{
			Title:       input.Title,
			Author:      input.Author,
			Description: input.Description,
			UserID:      &userIDPtr,
		}

		createdBook, err := services.CreateBook(db, book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{Error: "Failed to create book"})
			return
		}

		c.JSON(http.StatusOK, dto.CreateBookResponse{
			ID:          createdBook.ID,
			Title:       createdBook.Title,
			Author:      createdBook.Author,
			Description: createdBook.Description,
			UserID:      userID,
		})
	}
}

// GetAllBooksHandler godoc
// @Summary Get all books
// @Description Retrieve all books. Requires authentication via Bearer token.
// @Tags Books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.BookResponse
// @Failure 401 {object} dto.BookErrorResponse
// @Failure 500 {object} dto.BookErrorResponse
// @Router /books/get_all [get]
func GetAllBooksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := utils.Paginate[models.Book](c, db)
		if err != nil {
			c.JSON(500, dto.BookErrorResponse{Error: err.Error()})
			return
		}

		booksWithStats, err := services.GetAllBooksService(db, res.Data)
		if err != nil {
			c.JSON(500, dto.BookErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"page":  res.Page,
			"size":  res.Size,
			"total": res.Total,
			"books": booksWithStats,
		})
	}
}

// GetBookHandler godoc
// @Summary Get a book by id
// @Description Retrieve a single book details by providing its id. Requires authentication.
// @Tags Books
// @Accept json
// @Produce json
// @Param id path string true "id of the book"
// @Security ApiKeyAuth
// @Success 200 {array} dto.BookResponse
// @Failure 401 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Router /books/get/{id} [get]
func GetBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")

		bookWithStats, err := services.GetBook(db, idParam)
		if err != nil {
			c.JSON(404, gin.H{"error": "Book not found in library"})
			return
		}

		c.JSON(200, gin.H{"book": bookWithStats})
	}
}

// UpdateBookHandler godoc
// @Summary Update a book
// @Description Update book details by book ID. Requires authentication.
// @Tags Books
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path int true "Book ID"
// @Param data body dto.UpdateBookRequest true "Book update data"
// @Security ApiKeyAuth
// @Success 200 {object} dto.UpdateBookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 401 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Router /books/update/{id} [patch]
func UpdateBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		bookID, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.BookErrorResponse{Error: "Invalid book ID"})
			return
		}

		var input dto.UpdateBookRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, dto.BookErrorResponse{Error: "Invalid JSON"})
			return
		}

		updateData := map[string]interface{}{}
		if input.Title != "" {
			updateData["title"] = input.Title
		}
		if input.Author != "" {
			updateData["author"] = input.Author
		}
		if input.Description != "" {
			updateData["description"] = input.Description
		}

		updatedBook, err := services.UpdateBook(db, uint(bookID), updateData)
		if err != nil {
			c.JSON(http.StatusNotFound, dto.BookErrorResponse{Error: "Book not found or update failed"})
			return
		}

		c.JSON(http.StatusOK, dto.UpdateBookResponse{
			Title:       updatedBook.Title,
			Author:      updatedBook.Author,
			Description: updatedBook.Description,
			UserID:      *updatedBook.UserID,
		})
	}
}

// DeleteBookHandler godoc
// @Summary Delete a book
// @Description Delete a book by its ID. Requires authentication.
// @Tags Books
// @Security BearerAuth
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.DeleteBookResponse
// @Failure 400 {object} dto.DeleteErrorResponse
// @Failure 401 {object} dto.DeleteErrorResponse
// @Failure 404 {object} dto.DeleteErrorResponse
// @Router /books/delete/{id} [delete]
func DeleteBookHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		bookID, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		if err := services.DeleteBook(db, uint(bookID)); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found or deletion failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}

// FilterBooksHandler godoc
// @Summary Filter a book
// @Description Filter books by author or title. At least one filter is required. Requires authentication.
// @Tags Books
// @Security BearerAuth
// @Produce json
// @Param author query []string false "Filter by author (can be multiple)"
// @Param title query []string false "Filter by title (can be multiple)"
// @Success 200 {object} dto.BookResponse
// @Failure 400 {object} dto.BookErrorResponse
// @Failure 401 {object} dto.BookErrorResponse
// @Failure 404 {object} dto.BookErrorResponse
// @Router /books/search/ [get]
func FilterBooksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter dto.FilterBooksRequest
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, dto.BookErrorResponse{
				Error: err.Error(),
			})
			return
		}

		query := db.Model(&models.Book{})
		if len(filter.Author) > 0 {
			query = query.Where("author IN ?", filter.Author)
		}
		if len(filter.Title) > 0 {
			query = query.Where("title IN ?", filter.Title)
		}

		paginationResult, err := utils.Paginate[models.Book](c, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.BookErrorResponse{
				Error: err.Error(),
			})
			return
		}

		var response []dto.BookResponse
		for _, book := range paginationResult.Data {
			response = append(response, dto.BookResponse{
				Title:       book.Title,
				Author:      book.Author,
				Description: book.Description,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"page":  paginationResult.Page,
			"size":  paginationResult.Size,
			"total": paginationResult.Total,
			"books": response,
		})
	}
}
