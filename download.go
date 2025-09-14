package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Download(router fiber.Router, db *gorm.DB) {
	router.Get("/", func(c *fiber.Ctx) error {
		format := c.Query("format", "json") //json or csv
		books := new([]Book)
		if err := db.Find(&books).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get books"})
		}
		var filename string

		switch format {
		case "json":
			filename = "books.json"
			file, err := os.Create(filename)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create file"})
			}
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(books); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to encode books"})
			}
		case "csv":
			filename = "books.csv"
			file, err := os.Create(filename)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create file"})
			}
			defer file.Close()
			csvWriter := csv.NewWriter(file)
			if err := csvWriter.Write([]string{"Title", "Author", "Year", "Status"}); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to write books"})
			}
			for _, book := range *books {
				if err := csvWriter.Write([]string{book.Title, book.Author, strconv.Itoa(book.Year), string(book.Status)}); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to write books"})
				}
			}
			csvWriter.Flush()
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid format"})

		}

		defer os.Remove(filename)
		return c.Download(filename)
	})
}
