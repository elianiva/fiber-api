package main

import (
	"log"

	"github.com/elianiva/fiber-api/handlers"
	"github.com/elianiva/fiber-api/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// use cors
	app.Use(cors.New())

	// default routes, nothing interesting here
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 世界！")
	})

	// routes
	app.Get("/api/book/", handlers.GetBooks)
	app.Get("/api/book/id/:id", handlers.GetBooks)
	app.Get("/api/book/name/:name", handlers.GetBooks)
	app.Post("/api/book", handlers.AddBook)
	app.Delete("/api/book/id/:id", handlers.DeleteBook)
	app.Patch("/api/book/id/:id", handlers.UpdateBook)

	// listen to this port
	log.Fatal(app.Listen(":" + helpers.GetEnv("PORT")))
}
