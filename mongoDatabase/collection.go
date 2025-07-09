package mongoDatabase

import (
	"crud-api/database"
	"crud-api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserCollection() *mongo.Collection {
	return database.MongoDatabase.Collection("users")
}

// Create
func CreateUserInDB(user *models.User) error {
	_, err := UserCollection().InsertOne(context.TODO(), user)
	return err
}

// Read
func GetUserFromDB(id string) (*models.User, error) {
	var user models.User
	err := UserCollection().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}

// Update
func UpdateUserInDB(id string, update bson.M) error {
	_, err := UserCollection().UpdateByID(context.TODO(), id, update)
	return err
}

func GetAllUsersFromDB(page int64, limit int64) ([]models.User, error) {
	skip := (page - 1) * limit

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit)

	cursor, err := UserCollection().Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	return users, nil
}
// Delete
func DeleteUserFromDB(id string) error {
	_, err := UserCollection().DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
