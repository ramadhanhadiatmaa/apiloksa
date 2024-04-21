package middlewares

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func AuthMiddleware(c *fiber.Ctx) error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading .env files: %v", err)
	}

	secret_key := os.Getenv("SECRET_KEY")

	header := c.Get("apiKey")
	if header == "" || header != secret_key {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized Access Request",
		})
	}
	return c.Next()
}