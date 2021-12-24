package main

import (
	"github.com/boconlonton/ho-golang-fiber-library-app/pkg/configs"
	"github.com/boconlonton/ho-golang-fiber-library-app/pkg/middleware"
	"github.com/boconlonton/ho-golang-fiber-library-app/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

// @title API
// @version 1.0
// @description This is an auto-generated API docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// Define Fiber config
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app)

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.
}
