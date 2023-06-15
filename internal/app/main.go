package app

import (
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World5!")
	})

	app.Listen(":4444")
}
