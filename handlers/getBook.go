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

func GetBooks(c *fiber.Ctx) error {
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
	if param := c.Params("id"); param != "" {
		objID, _ := primitive.ObjectIDFromHex(param)
		filter = bson.M{"_id": objID}
	} else if param := c.Params("name"); param != "" {
		filter = bson.M{"name": param}
	}

	// get cursor to iterate through available data
	cur, curErr := collection.Find(context.Background(), filter)
	if curErr != nil {
		errResp, _ := json.Marshal(helpers.Result{
			Status: "500",
			Data:   make([]helpers.Book, 0),
		})
		c.Status(500).Send(errResp)
		log.Fatal(curErr)
	}
	defer cur.Close(context.Background())

	// put all results into result
	result := make([]helpers.Book, 0)
	cur.All(context.Background(), &result)

	if len(result) == 0 {
		jsonResp, _ := json.Marshal(helpers.Result{
			Status: "404",
			Data:   result,
		})
		c.Set("Content-Type", "application/json")
		return c.SendString(string(jsonResp))
	}

	jsonResp, _ := json.Marshal(helpers.Result{
		Status: "200",
		Data:   result,
	})

	// send back the data
	// TODO: change this to message instead
	c.Set("Content-Type", "application/json")
	return c.SendString(string(jsonResp))
}
