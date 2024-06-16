package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

type Book struct {
	ID     int    `json:"id"`
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
}

// Dummy user for login
var user = struct {
	Email    string
	Password string
}{
	Email:    "user@example.com",
	Password: "password123",
}

// UserData represents the user data extracted from the JWT token
type UserData struct {
	Email string
	Role  string
}

// userContextKey is the key used to store user data in the Fiber context
const userContextKey = "user"

var books []Book

func main() {
	books = append(books, Book{ID: 1, TITLE: "hello", AUTHOR: "slothpete"})
	books = append(books, Book{ID: 2, TITLE: "introactive", AUTHOR: "zelsa"})

	app := fiber.New()

	// JWT Secret Key
	secretKey := "secret"

	// Login route
	app.Post("/login", login(secretKey))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secretKey),
	}))

	// Middleware to extract user data from JWT
	app.Use(extractUserFromJWT)

	app.Get("/books", getAllBooks)
	app.Get("/books/:id", getBookByID)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Listen("localhost:8080")
}
