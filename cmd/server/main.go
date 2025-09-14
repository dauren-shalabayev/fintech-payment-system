package main

import (
	"fiber-golang-json-api/internal/config"
	"fiber-golang-json-api/internal/database"
	"fiber-golang-json-api/internal/handler"
	"fiber-golang-json-api/internal/middleware"
	"fiber-golang-json-api/internal/repository"
	"fiber-golang-json-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db := database.InitializeDB(cfg)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	bookService := service.NewBookService(bookRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(userService)
	bookHandler := handler.NewBookHandler(bookService)
	downloadHandler := handler.NewDownloadHandler(bookService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:   "My Fiber App",
		ETag:      true,
		BodyLimit: 1024 * 1024 * 5,
	})

	// Public routes
	app.Post("/auth/register", authHandler.Register)
	app.Post("/auth/login", authHandler.Login)
	app.Get("/download/", downloadHandler.DownloadBooks)

	// Protected routes
	protected := app.Use(middleware.AuthMiddleware())
	protected.Get("/book/", bookHandler.GetBooks)
	protected.Get("/book/:id", bookHandler.GetBookByID)
	protected.Post("/book/", bookHandler.CreateBook)
	protected.Put("/book/:id", bookHandler.UpdateBook)
	protected.Delete("/book/:id", bookHandler.DeleteBook)

	// Start server
	app.Listen(":" + cfg.Server.Port)
}
