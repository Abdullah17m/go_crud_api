package models

import (
	"crud-api/database"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func UserCollection() *mongo.Collection {
	return database.MongoDatabase.Collection("users")
}