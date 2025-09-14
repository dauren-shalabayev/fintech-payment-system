package handler

import (
	"fiber-golang-json-api/internal/models"
	"fiber-golang-json-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)

	// Get filters from query parameters
	filters := make(map[string]string)
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if author := c.Query("author"); author != "" {
		filters["author"] = author
	}
	if title := c.Query("title"); title != "" {
		filters["title"] = title
	}
	if year := c.Query("year"); year != "" {
		filters["year"] = year
	}

	books, err := h.bookService.GetBooksByUserIDWithFilters(uint(userID), filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get books"})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *BookHandler) GetBookByID(c *fiber.Ctx) error {
	bookID, _ := c.ParamsInt("id")
	userID := c.Locals("userID").(float64)

	book, err := h.bookService.GetBookByID(uint(bookID), uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	userID := c.Locals("userID").(float64)

	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	book.UserID = uint(userID)

	if err := h.bookService.CreateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create book"})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func (h *BookHandler) UpdateBook(c *fiber.Ctx) error {
	bookID, _ := c.ParamsInt("id")
	userID := c.Locals("userID").(float64)

	// Find existing book
	book, err := h.bookService.GetBookByID(uint(bookID), uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Parse update data
	updateData := new(models.Book)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Update only non-zero fields
	if updateData.Title != "" {
		book.Title = updateData.Title
	}
	if updateData.Author != "" {
		book.Author = updateData.Author
	}
	if updateData.Status != "" {
		book.Status = updateData.Status
	}
	if updateData.Year != 0 {
		book.Year = updateData.Year
	}

	if err := h.bookService.UpdateBook(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update book"})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	bookID, _ := c.ParamsInt("id")
	userID := c.Locals("userID").(float64)

	if err := h.bookService.DeleteBook(uint(bookID), uint(userID)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}
