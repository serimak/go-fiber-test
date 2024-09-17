package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Go!")
	})

	app.Get("/api/v1/test", func(c *fiber.Ctx) error {
		return c.SendString("Test, Success!")
	})

	app.Listen(":3000")

}
