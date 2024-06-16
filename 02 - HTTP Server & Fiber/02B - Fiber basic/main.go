package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello WOrld!")
	})

	app.Listen("localhost:8080")
}
