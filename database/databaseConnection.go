package database

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

func DBinstance() mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Err loading .env file")
	}

	mongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect()
}
