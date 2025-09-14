package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthHandlers(router fiber.Router, db *gorm.DB) {
	router.Post("/register", func(c *fiber.Ctx) error {
		user := &User{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}
		if user.Username == "" || user.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username and password are required"})
		}
		hedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		user.Password = string(hedPassword)
		if err := db.Create(user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
		}

		token, err := GenerateToken(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	})

	router.Post("/login", func(c *fiber.Ctx) error {

		authUser := &User{
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
		}
		if authUser.Username == "" || authUser.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username and password are required"})
		}

		dbUser := new(User)
		if err := db.Where("username = ?", authUser.Username).First(dbUser).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(authUser.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		token, err := GenerateToken(dbUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
	})
}
