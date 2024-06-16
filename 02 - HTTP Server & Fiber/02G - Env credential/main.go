package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
}

var books []Book

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Setup routes
	app.Get("/api/config", getConfig)

	// Use the environment variable for the port
	port := os.Getenv("PORT")
	// port := os.Getenv("_PORT")
	if port == "" {
		port = "8089" // Default port if not specified
	}

	app.Listen(":" + port)
}

func getConfig(c *fiber.Ctx) error {
	// Example: Return a configuration value from environment variable
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		secretKey = "defaultSecret" // Default value if not specified
	}

	return c.JSON(fiber.Map{
		"secret_key": secretKey,
	})
}
