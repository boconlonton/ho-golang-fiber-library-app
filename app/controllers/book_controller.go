package controllers

import (
	"github.com/boconlonton/ho-golang-fiber-library-app/app/models"
	"github.com/boconlonton/ho-golang-fiber-library-app/app/utils"
	"github.com/boconlonton/ho-golang-fiber-library-app/platform/database"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"time"
)

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary Get all exists books
// @Tags Books
// @Accept json
// @Product json
// @Success 200 {array} models.Book
// @Router /v1/books [get]
func GetBooks(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all books
	books, err := db.GetBooks()
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "books were not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(books),
		"books": books,
	})
}

// GetBook func gets book by given ID or 404 error.
// @Description Get book by given ID.
// @Summary Get book by given ID.
// @Tags Book
// @Accept json
// @Product json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Router /v1/book/{id} [get]
func GetBook(c *fiber.Ctx) error {
	// Catch book ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get book by ID.
	book, err := db.GetBook(id)
	if err != nil {
		// Return, if nook not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with the given ID is not found",
			"book":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary created a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_attrs body models.BookAttrs true "Book attributes"
// @Success 201 {object} models.Book
// @Security ApiKeyAuth
// @Router /v1/book [post]
func CreateBook(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Book struct
	book := &models.Book{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(book); err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create new validator for a Book model.
	validate := utils.NewValidator()

	// Set initialized default data for book:
	book.ID = uuid.New()
	book.CreatedAt = time.Now()
	book.BookStatus = 1 // 0 == draft, 1 == active

	// Validate book fields.
	if err := validate.Struct(book); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Send query to create book
	if err := db.CreateBook(book); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}
