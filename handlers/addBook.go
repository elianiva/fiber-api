package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/elianiva/fiber-api/helpers"
	"go.mongodb.org/mongo-driver/bson"

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

	count, err := collection.CountDocuments(context.Background(), bson.M{"name": c.FormValue("name")})
	if err != nil {
		helpers.ThrowErr(c, fileErr, "500")
	}

	zeroRes := make([]helpers.Book, 0)
	if count >= 1 {
		jsonRes, _ := json.Marshal(helpers.Result{
			Status:  "200",
			Message: "Data already exists",
			Data:    zeroRes,
		})
		c.Send(jsonRes)
	}

	// make new data instance
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	book := helpers.Book{
		Name:   c.FormValue("name"),
		Author: c.FormValue("author"),
		Pages:  pages,
		ImgUrl: fmt.Sprintf("/public/images/%d-%s", timestamp, file.Filename),
	}

	// save file
	saveErr := c.SaveFile(file, fmt.Sprintf("./public/images/%s", file.Filename))
	if saveErr != nil {
		helpers.ThrowErr(c, saveErr, "500")
	}

	// insert data to mongodb
	_, insertErr := collection.InsertOne(context.Background(), book)
	if insertErr != nil {
		helpers.ThrowErr(c, insertErr, "500")
	}

	c.Set("Content-Type", "application/json")
	result := make([]helpers.Book, 0)
	jsonResp, _ := json.Marshal(helpers.Result{
		Status:  "201",
		Message: "Data has been successfully created.",
		Data:    result,
	})
	return c.Send(jsonResp)
}
