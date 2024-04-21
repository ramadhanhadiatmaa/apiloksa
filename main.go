package main

import (
	"apiloksa/models"
	"apiloksa/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	models.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(cors.New())

	routes.Routes(app)

	app.Listen(":8801")
}