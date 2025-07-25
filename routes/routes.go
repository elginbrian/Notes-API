package routes

import (
	"notes-api/handlers"
	"notes-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root endpoint - API information
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":     "Welcome to Notes API",
			"version":     "1.0.0",
			"status":      "running",
			"docs":        "Visit /swagger/ for API documentation",
			"health":      "/health",
			"endpoints": fiber.Map{
				"auth": fiber.Map{
					"register": "POST /api/auth/register",
					"login":    "POST /api/auth/login",
				},
				"notes": fiber.Map{
					"list":   "GET /api/notes",
					"get":    "GET /api/notes/:id",
					"create": "POST /api/notes",
					"update": "PUT /api/notes/:id",
					"delete": "DELETE /api/notes/:id",
				},
			},
		})
	})

	// API group
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// Protected routes
	notes := api.Group("/notes")
	notes.Use(middleware.Protected())
	notes.Get("/", handlers.GetNotes)
	notes.Get("/:id", handlers.GetNote)
	notes.Post("/", handlers.CreateNote)
	notes.Put("/:id", handlers.UpdateNote)
	notes.Delete("/:id", handlers.DeleteNote)
}
