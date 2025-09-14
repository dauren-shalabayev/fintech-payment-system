package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookHandlers(router fiber.Router, db *gorm.DB) {
	router.Get("/", func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(float64)
		books := []Book{}

		// Start with base query for current user
		query := db.Where("user_id = ?", uint(userID))

		// Add filters based on query parameters
		if status := c.Query("status"); status != "" {
			query = query.Where("status = ?", status)
		}

		if author := c.Query("author"); author != "" {
			query = query.Where("author LIKE ?", "%"+author+"%")
		}
		fmt.Println("title", c.Query("title"))
		if title := c.Query("title"); title != "" {
			query = query.Where("title LIKE ?", "%"+title+"%")
		}

		if year := c.Query("year"); year != "" {
			query = query.Where("year = ?", year)
		}

		// Execute query
		if err := query.Find(&books).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get books"})
		}

		return c.Status(fiber.StatusOK).JSON(books)
	})

	router.Get("/:id", func(c *fiber.Ctx) error {
		bookID, _ := c.ParamsInt("id")
		userID := c.Locals("userID").(float64)
		var book Book
		if err := db.Where("id = ? AND user_id = ?", bookID, userID).First(&book).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		fmt.Println("book", book)
		return c.Status(fiber.StatusOK).JSON(book)
	})

	router.Post("/", func(c *fiber.Ctx) error {
		book := new(Book)
		userID := c.Locals("userID").(float64)

		// Parse book data from request body
		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Set user ID
		book.UserID = uint(userID)

		// Save to database
		if err := db.Create(book).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
		}

		return c.Status(fiber.StatusCreated).JSON(book)
	})

	router.Put("/:id", func(c *fiber.Ctx) error {
		bookID, _ := c.ParamsInt("id")
		userID := c.Locals("userID").(float64)
		book := new(Book)
		// Find existing book
		if err := db.Where("id = ? AND user_id = ?", bookID, uint(userID)).First(&book).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		// Parse update data from request body

		if err := c.BodyParser(book); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if err := db.Save(&book).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
		}
		return c.Status(fiber.StatusOK).JSON(book)
	})

	router.Delete("/:id", func(c *fiber.Ctx) error {
		bookID, _ := c.ParamsInt("id")
		userID := c.Locals("userID").(float64)
		book := new(Book)
		if err := db.Where("id = ? AND user_id = ?", bookID, uint(userID)).First(&book).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		if err := db.Delete(&book).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete book"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
	})
}
