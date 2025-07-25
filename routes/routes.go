package routes

import (
	"notes-api/handlers"
	"notes-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
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
