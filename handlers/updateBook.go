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

func UpdateBook(c *fiber.Ctx) error {
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

	// get id params
	var filter bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	filter = bson.M{"_id": objID}

	var result bson.M
	jsonErr := json.Unmarshal(c.Body(), &result)
	if jsonErr != nil {
		errResp, _ := json.Marshal(helpers.Result{
			Status: "500",
			Data:   make([]helpers.Book, 0),
		})
		c.Status(500).Send(errResp)
		log.Fatal(jsonErr)
	}

	update := bson.M{
		"$set": result,
	}

	_, updateErr := collection.UpdateOne(context.Background(), filter, update)
	if updateErr != nil {
		errResp, _ := json.Marshal(helpers.Result{
			Status: "500",
			Data:   make([]helpers.Book, 0),
		})
		c.Status(500).Send(errResp)
		log.Fatal(updateErr)
	}

	// send back the data
	// TODO: change this to message instead
	c.Set("Content-Type", "application/json")
	resp := make([]helpers.Book, 0)
	jsonResp, _ := json.Marshal(helpers.Result{
		Status: "204",
		Data:   resp,
	})
	return c.Send(jsonResp)
}
