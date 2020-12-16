package handlers

import (
	"context"
	"encoding/json"
	"fiber-api/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	collection, err := helpers.GetMongoCollection(
		helpers.GetEnv("DB_NAME"), helpers.GetEnv("COLLECTION_NAME"),
	)

	if err != nil {
		c.Status(500).Send([]byte("Error"))
		log.Fatal(err)
	}

	// get data from request body
	var book helpers.Book
	json.Unmarshal([]byte(c.Body()), &book)

	// insert data to mongodb
	res, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		c.Status(500).Send([]byte("Error"))
		log.Fatal(err)
	}

	// send back the data
	// TODO: change this to success message instead
	c.Set("Content-Type", "application/json")
	resp, _ := json.Marshal(res)
	return c.Send(resp)
}
