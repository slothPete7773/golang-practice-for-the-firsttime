package main

import "github.com/gofiber/fiber/v2"

type Book struct {
	ID     int    `json:"id"`
	TITLE  string `json:"title"`
	AUTHOR string `json:"author"`
}

var books []Book

func main() {

	books = append(books, Book{ID: 1, TITLE: "hello", AUTHOR: "slothpete"})
	books = append(books, Book{ID: 2, TITLE: "introactive", AUTHOR: "zelsa"})
	app := fiber.New()
	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello WOrld!")
	// })
	app.Get("/books", getAllBooks)
	app.Get("/books/:id", getBookByID)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadImage)

	app.Listen("localhost:8080")
}

func uploadImage(c *fiber.Ctx) error {
	// Read file from request
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Save the file to the server
	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully: " + file.Filename)

}
