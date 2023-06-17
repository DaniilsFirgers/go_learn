package app

import (
	"go-learn/tests"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World5!")
	})
	tests.InsertDummyDocumentIntoDb()

	app.Listen(":4444")
}
