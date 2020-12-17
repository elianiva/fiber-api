package helpers

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoConnection() (*mongo.Client, error) {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongodbURI := os.Getenv("MONGODB_URI")

	// connect to mongodb
	client, err := mongo.Connect(
		context.Background(), options.Client().ApplyURI(mongodbURI),
	)
	if err != nil {
		log.Fatal(err)
	}

	// check if it can connect to mongodb or not
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// returns mongodb client
	return client, nil
}

func GetMongoCollection(DBName, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoConnection()
	if err != nil {
		return nil, err
	}

	collection := client.Database(DBName).Collection(CollectionName)

	return collection, nil
}
