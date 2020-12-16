package main

import (
	"log"

	"fiber-api/handlers"
	"github.com/gofiber/fiber/v2"
)

const PORT string = ":3000"

func main() {
	app := fiber.New()

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Henlo")
	})

	app.Get("/api/book/", handlers.GetBooks)
	app.Get("/api/book/id/:id", handlers.GetBooks)
	app.Get("/api/book/name/:name", handlers.GetBooks)
	app.Post("/api/book", handlers.AddBook)

	// listen to this port
	log.Fatal(app.Listen(PORT))
}
