package handlers

import (
	"context"
	"encoding/json"
	"github.com/elianiva/fiber-api/helpers"
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
	json.Unmarshal(c.Body(), &book)

	// insert data to mongodb
	_, insertErr := collection.InsertOne(context.Background(), book)
	if insertErr != nil {
		c.Status(500).Send([]byte("Error"))
		log.Fatal(insertErr)
	}

	// send back the data
	// TODO: change this to success message instead
	c.Set("Content-Type", "application/json")
	result := make([]helpers.Book, 0)
	jsonResp, _ := json.Marshal(helpers.Result{
		Status: "201",
		Data:   result,
	})
	return c.Send(jsonResp)
}
