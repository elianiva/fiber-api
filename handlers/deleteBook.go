package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/elianiva/fiber-api/helpers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBook(c *fiber.Ctx) error {
	collection, err := helpers.GetMongoCollection(
		helpers.GetEnv("DB_NAME"), helpers.GetEnv("COLLECTION_NAME"),
	)

	if err != nil {
		errResp, _ := json.Marshal(helpers.Result{
			Status: "500",
			Data:   make([]helpers.Book, 0),
		})
		c.Status(500).Send(errResp)
		log.Fatal(err)
	}

	// get params
	var filter bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	filter = bson.M{"_id": objID}

	// get cursor to iterate through available data
	_, delErr := collection.DeleteOne(context.Background(), filter)
	if delErr != nil {
		errResp, _ := json.Marshal(helpers.Result{
			Status: "500",
			Data:   make([]helpers.Book, 0),
		})
		c.Status(500).Send(errResp)
		log.Fatal(err)
	}

	// TODO: change this to success message instead
	c.Set("Content-Type", "application/json")
	result := make([]helpers.Book, 0)
	jsonResp, _ := json.Marshal(helpers.Result{
		Status: "202",
		Data:   result,
	})
	return c.Send(jsonResp)
}
