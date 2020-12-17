package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/elianiva/fiber-api/helpers"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	collection, err := helpers.GetMongoCollection(
		helpers.GetEnv("DB_NAME"), helpers.GetEnv("COLLECTION_NAME"),
	)

	if err != nil {
		helpers.ThrowErr(c, err, "500")
	}

	// get file and pages
	pages, _ := strconv.Atoi(c.FormValue("pages"))
	file, fileErr := c.FormFile("img")
	if fileErr != nil {
		helpers.ThrowErr(c, fileErr, "500")
	}

	// get data from request body
	var book helpers.Book
	json.Unmarshal(c.Body(), &book)

	// insert data to mongodb
	_, insertErr := collection.InsertOne(context.Background(), book)
	if insertErr != nil {
		helpers.ThrowErr(c, insertErr, "500")
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
