package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(name)
}
