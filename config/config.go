package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI string
	MongoDB  string
	AppPort  string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	MongoURI = os.Getenv("MONGO_URI")
	MongoDB = os.Getenv("MONGO_DB")

	if MongoURI == "" || MongoDB == "" {
		log.Fatal("Missing MONGO_URI or MONGO_DB in environment")
	}

	if AppPort == "" {
		AppPort = "8080" // fallback default
	}
}
