package main

import (
	"log"

	"github.com/elianiva/fiber-api/handlers"
	"github.com/gofiber/fiber/v2"
)

const PORT string = ":3000"

func main() {
	// TODO: figure out the correct status code for each responses
	app := fiber.New()

	// routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 世界！")
	})

	app.Get("/api/book/", handlers.GetBooks)
	app.Get("/api/book/id/:id", handlers.GetBooks)
	app.Get("/api/book/name/:name", handlers.GetBooks)
	app.Post("/api/book", handlers.AddBook)
	app.Delete("/api/book/id/:id", handlers.DeleteBook)
	app.Patch("/api/book/id/:id", handlers.UpdateBook)

	// listen to this port
	log.Fatal(app.Listen(PORT))
}
