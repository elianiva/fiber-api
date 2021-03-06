package handlers

import (
	"context"
	"encoding/json"

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
		helpers.ThrowErr(c, err, "500")
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
		helpers.ThrowErr(c, curErr, "500")
	}
	defer cur.Close(context.Background())

	// put all results into result
	result := make([]helpers.Book, 0)
	cur.All(context.Background(), &result)

	if len(result) == 0 {
		c.Set("Content-Type", "application/json")
		jsonRes, _ := json.Marshal(helpers.Result{
			Status: "200",
			Data:   result,
		})
		return c.Send(jsonRes)
	}

	c.Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(helpers.Result{
		Status:  "200",
		Message: "Data has been successfully retrieved.",
		Data:    result,
	})
	return c.Send(jsonResp)
}
