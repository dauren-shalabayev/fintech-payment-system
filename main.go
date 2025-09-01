package main

import "github.com/gofiber/fiber/v2"

func main() {
	//Create a new fiber app
	app := fiber.New(fiber.Config{
		AppName:   "My Fiber App",
		ETag:      true,
		BodyLimit: 1024 * 1024 * 5,
	})

	//Create a new route
	app.Listen(":3000")

}
