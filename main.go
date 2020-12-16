package main

import (
	// "fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, 世界!")
	})

	log.Fatal(app.Listen(":3000"))
}
