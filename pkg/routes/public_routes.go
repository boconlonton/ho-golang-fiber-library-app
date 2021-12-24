package routes

import (
	"github.com/boconlonton/ho-golang-fiber-library-app/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/books", controllers.GetBooks)
	route.Get("/book/:id", controllers.GetBook)
	route.Get("/token/new", controllers.GetNewAccessToken)
}
