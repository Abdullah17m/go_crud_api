package database

import (
	"context"
	"log"
	"time"

	"crud-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(config.MongoDB)
}
