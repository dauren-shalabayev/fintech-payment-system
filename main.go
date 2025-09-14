package main

import "github.com/gofiber/fiber/v2"

func main() {
	db := InitializeDB()
	//Create a new fiber app
	app := fiber.New(fiber.Config{
		AppName:   "My Fiber App",
		ETag:      true,
		BodyLimit: 1024 * 1024 * 5,
	})

	// Define a new route
	AuthHandlers(app.Group("/auth"), db)
	Download(app.Group("/download"), db)

	protected := app.Use(AuthMiddleware(db))

	BookHandlers(protected.Group("/book"), db)

	//download csv

	//Create a new route
	app.Listen(":3000")

}
