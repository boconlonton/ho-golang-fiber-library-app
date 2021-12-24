package routes

import (
	"github.com/boconlonton/ho-golang-fiber-library-app/app/controllers"
	"github.com/boconlonton/ho-golang-fiber-library-app/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book

	// Route for PUT method:
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update a book by ID

	// Routes for DELETE method
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete a book by ID
}
