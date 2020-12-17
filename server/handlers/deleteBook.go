package handlers

import (
	"context"
	"encoding/json"

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
		helpers.ThrowErr(c, err, "500")
	}

	// get params
	var filter bson.M
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	filter = bson.M{"_id": objID}

	// get cursor to iterate through available data
	_, delErr := collection.DeleteOne(context.Background(), filter)
	if delErr != nil {
		helpers.ThrowErr(c, delErr, "500")
	}

	// TODO: change this to success message instead
	c.Set("Content-Type", "application/json")
	result := make([]helpers.Book, 0)
	jsonResp, _ := json.Marshal(helpers.Result{
		Status:  "200",
		Message: "Data has been successfully deleted.",
		Data:    result,
	})
	return c.Send(jsonResp)
}
